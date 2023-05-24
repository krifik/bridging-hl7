package sftp

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/k0kubun/pp"
	"github.com/krifik/bridging-hl7/exception"
	"github.com/krifik/bridging-hl7/helper"

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
		// err = os.Remove(pdf.Name())
		// exception.SendLogIfErorr(err, "70")
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
		log.Println("Failed to dial: " + err.Error())
	}
	defer conn.Close()

	sftpClient, err := sftp.NewClient(conn)
	if err != nil {
		log.Println("Failed to create SFTP client: " + err.Error())
	}

	// Directory to monitor for new files
	remoteDir := os.Getenv("SFTP_ORDER_DIR")

	// Track the last modified time of the latest file
	var lastModified time.Time

	for {
		// List files in the remote directory
		files, err := sftpClient.ReadDir(remoteDir)
		if err != nil {
			fmt.Printf("Failed to read remote directory: %v", err)
			return
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
				// if err == nil {
				// 	sftpClient.Remove(os.Getenv("SFTP_ORDER_DIR") + "/" + file.Name() + ".txt")
				// }
			}
		}

		// Sleep for a duration before checking again
		time.Sleep(time.Second * 2) // Adjust the duration as needed
	}
}
