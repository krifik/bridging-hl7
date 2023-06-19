package sftp

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/k0kubun/pp"
	"github.com/krifik/bridging-hl7/config"
	"github.com/krifik/bridging-hl7/entity"
	"github.com/krifik/bridging-hl7/exception"
	"github.com/krifik/bridging-hl7/helper"
	"gorm.io/gorm"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// type SFTPClient struct {
// 	client *sftp.Client
// }

func Upload(file, fileName, orderNumber, labNumber string) {
	err := godotenv.Load()
	exception.SendLogIfErorr(err, "13")

	config := &ssh.ClientConfig{
		User: os.Getenv("SFTP_USER"),
		Auth: []ssh.AuthMethod{
			ssh.Password(os.Getenv("SFTP_PASSWORD")),
			// interactiveSSHAuthMethod(os.Getenv("SFTP_PASSWORD")),
		},
		// HostKeyCallback: hostKeyCallback,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         30 * time.Second,
	}
	conn, err := ssh.Dial("tcp", os.Getenv("SFTP_URL"), config)
	if err != nil {
		pp.Println("Error Koneksi : ", err.Error())

	}
	defer conn.Close()
	pdf, err := helper.GetPDF(orderNumber, labNumber)
	exception.SendLogIfErorr(err, "42")
	defer pdf.Close()
	sftp, err := sftp.NewClient(conn)
	if err != nil {
		exception.SendLogIfErorr(err, "44")
	}
	localFile, err := os.Open(file)
	if err != nil {
		exception.SendLogIfErorr(err, "47")
	}
	defer localFile.Close()
	remoteFile, err := sftp.OpenFile(os.Getenv("SFTP_RESULT_DIR")+"/"+fileName+".txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC)

	if err != nil {
		exception.SendLogIfErorr(err, "53")
	}
	remoteFilePdf, err := sftp.Create(os.Getenv("SFTP_RESULT_DIR") + "/" + fileName + ".pdf")
	exception.SendLogIfErorr(err, "53")

	defer remoteFile.Close()
	defer remoteFilePdf.Close()
	_, err = remoteFile.ReadFrom(localFile)
	if err != nil {
		exception.SendLogIfErorr(err, "59")
	}

	pdfData, err := ioutil.ReadFile(pdf.Name())
	exception.SendLogIfErorr(err, "70")
	_, err = remoteFilePdf.Write(pdfData)
	if err != nil {
		exception.SendLogIfErorr(err, "59")
	}
	if err == nil {
		err = os.Remove(file)
		exception.SendLogIfErorr(err, "68")
		err = os.Remove(pdf.Name())
		exception.SendLogIfErorr(err, "70")
		pp.SetColorScheme(pp.ColorScheme{
			String: pp.Red,
		})
		pp.Println("Local file removed successfully! " + fileName)
		pp.ResetColorScheme()
		exception.SendLogIfErorr(err, "63")
		pp.SetColorScheme(pp.ColorScheme{
			String: pp.Cyan,
		})

		pp.Println("File uploaded successfully to sftp server! " + fileName)
		pp.ResetColorScheme()
	}
	// ini yang bikin gabisa upload file
	// defer sftp.Close()
}

func Watcher() {
	err := godotenv.Load()
	exception.SendLogIfErorr(err, "13")
	db := config.InitializedSqlite()

	// Directory to monitor for new files
	remoteDir := os.Getenv("SFTP_ORDER_DIR")
	sftpUrl := os.Getenv("SFTP_URL")
	user := os.Getenv("SFTP_USER")
	password := os.Getenv("SFTP_PASSWORD")
	// Track the last modified time of the latest file
	var lastModified time.Time
	var sftpClient *sftp.Client
	for {
		select {
		case <-time.After(5 * time.Second):
			if sftpClient == nil {
				pp.Println("Connecting to SFTP server...")
				config := &ssh.ClientConfig{
					User: user,
					Auth: []ssh.AuthMethod{
						ssh.Password(password),
						// interactiveSSHAuthMethod(os.Getenv("SFTP_PASSWORD")),
					},
					// HostKeyCallback: hostKeyCallback,
					HostKeyCallback: ssh.InsecureIgnoreHostKey(),
					Timeout:         30 * time.Second,
				}
				conn, errConn := ssh.Dial("tcp", sftpUrl, config)
				if errConn != nil {
					log.Println("Failed to dial: " + errConn.Error())
					// conn.Close()
				}

				sftpClient, err = sftp.NewClient(conn)
				if err != nil {
					log.Println("Failed to create SFTP client: " + err.Error())
				}
			}
		}

		if sftpClient != nil {
			pp.Println("Connected to SFTP Server!")
			// List files in the remote directory
			files, err := sftpClient.ReadDir(remoteDir)
			if err != nil {
				fmt.Printf("Failed to read remote directory: %v", err)

			}

			// Iterate over the files

			for _, file := range files {

				// Check if the file is newer than the last known modification time

				if file.ModTime().After(lastModified) {

					scheme := pp.ColorScheme{
						// Integer: pp.Green | pp.Bold,
						Float:  pp.Black | pp.BackgroundWhite | pp.Bold,
						String: pp.Blue,
					}
					pp.SetColorScheme(scheme)
					pp.Println("New file detected:", file.Name())
					fileContent := helper.GetContentSftpFile(file.Name(), sftpClient)
					// Update the last known modification time
					// helper.SendToAPI(fileContent)
					// pp.Println(fileContent)
					err = helper.SendJsonToRabbitMQ(fileContent)
					exception.SendLogIfErorr(err, "122")

					lastModified = file.ModTime()
					if err == nil {
						fileExist := db.Where("file_name = ?", file.Name()).First(&entity.File{})

						if errors.Is(fileExist.Error, gorm.ErrRecordNotFound) {
							db.Create(&entity.File{FileName: file.Name(), ReadState: true})
							pp.Println("Store to DB successfully!")
						}
						sftpClient.Remove(os.Getenv("SFTP_ORDER_DIR") + "/" + file.Name() + ".txt")
						pp.SetColorScheme(pp.ColorScheme{
							String: pp.Red,
						})
						pp.Println("Delete file " + file.Name() + "after send to rabbitmq successfully!")
					}
				}
				ReSendFileExist := entity.File{}
				isExist := db.Where("file_name = ?", file.Name()).Or(db.Where("read_state = ?", true)).First(&ReSendFileExist)
				// pp.Println(reFindFile)
				if errors.Is(isExist.Error, gorm.ErrRecordNotFound) {
					db.Create(&entity.File{FileName: file.Name(), ReadState: true})
					fileContentReFindCreate := helper.GetContentSftpFile(file.Name(), sftpClient)
					err = helper.SendJsonToRabbitMQ(fileContentReFindCreate)
					exception.SendLogIfErorr(err, "186")
					pp.Println("Store to DB successfully!")
				}
				var filesReadState []entity.File
				db.Where("read_state = ?", false).Find(&filesReadState)
				fileUpdate := entity.File{
					ReadState: true,
				}

				if len(filesReadState) > 0 {
					db.Model(filesReadState).Updates(fileUpdate)
					for _, item := range filesReadState {
						pp.Println("UPDATED READ STATE TO 1 : " + item.FileName)
						fileContentReFind := helper.GetContentSftpFile(item.FileName, sftpClient)
						err = helper.SendJsonToRabbitMQ(fileContentReFind)
						exception.SendLogIfErorr(err, "122")
					}
				}

			}
		} else {
			continue
		}
	}

}
