package service

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/krifik/bridging-hl7/exception"
	"github.com/krifik/bridging-hl7/model"
	"github.com/krifik/bridging-hl7/repository"
	"github.com/krifik/bridging-hl7/sftp"
	"github.com/krifik/bridging-hl7/utils"

	helper "github.com/krifik/bridging-hl7/helper"

	"github.com/joho/godotenv"
)

// import from consumer.go
var globalConsumer chan struct{}

type FileServiceImpl struct {
	FileRepository repository.FileRepository
}

func NewFileServiceImpl(fileRepository repository.FileRepository) FileService {
	return &FileServiceImpl{FileRepository: fileRepository}
}

// GetContentFile retrieves the content of a file given its URL in JSON format.
//
// url: the URL of the file to retrieve.
// returns: a model.Json object containing the content of the file.

func (f *FileServiceImpl) GetContentFile(url string) model.Json {
	err := godotenv.Load()
	if err != nil {
		utils.SendMessage("LINE 36\n" + "Log Type: Error\n" + "Error: \n" + err.Error() + "\n")
	}
	// pp.Print(os.Getenv("ORDERDIR"))
	if os.Getenv("APP_MODE") == "DEBUG" {
		fileName := "LAB-20230504-00001.txt"
		dir := os.Getenv("ORDERDIR")
		url = dir + "/" + fileName
	}
	// file, fileError := os.OpenFile(dir+"/"+fileName, os.O_RDONLY, 0644)
	fileContent := helper.GetContent(url)
	json := utils.TransformToRightJson(fileContent)

	return json
}
func (f *FileServiceImpl) GetFiles() []string {
	err := godotenv.Load()
	if err != nil {
		utils.SendMessage("LINE 50\n" + "Log Type: Error\n" + "Error: \n" + err.Error() + "\n")
	}
	dir := os.Getenv("ORDERDIR")
	if err != nil {
		utils.SendMessage("LINE 54	\n" + "Log Type: Error\n" + "Error: \n" + err.Error() + "\n")
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		utils.SendMessage("LINE 58\n" + "Log Type: Error\n" + "Error: \n" + err.Error() + "\n")
	}
	var results []string
	for _, entri := range entries {
		results = append(results, entri.Name())
	}
	return results
}
func (f *FileServiceImpl) SearchFile() string {
	return ""
}

// CreateFileResult creates a new file result.
//
// JSONRequest is the request model.
// Returns a string.

func (f *FileServiceImpl) CreateFileResult(request model.JSONRequest) (string, error) {
	var file model.FileResult
	rand.Seed(time.Now().UnixNano())
	// Generate a random integer between 0 and 999999
	num := rand.Intn(999999)
	file.Msh.MessageDT = "message_dt=" + time.Now().Format("200601021504")
	str := strconv.Itoa(num)
	file.Msh.Type = "[MSH]"
	file.Msh.MessageID = "message_id=" + "TDR-3000" + str
	file.Msh.Version = "version=" + "2.3"
	file.Obr.Type = "[OBR]"
	file.Obr.PID = "pid=" + request.Data.Response.Demographics.Patient.MRN
	file.Obr.Apid = "apid=" + ""
	file.Obr.Pname = "pname=" + request.Data.Response.Demographics.Patient.FullName
	file.Obr.Pidentityno = "pidentityno=" + request.Data.Response.Demographics.Patient.IDNumber
	file.Obr.Pmobileno = "pmobileno=" + request.Data.Response.Demographics.Patient.PhoneNumber
	file.Obr.Street = "street=" + request.Data.Response.Demographics.Patient.Address
	file.Obr.Title = "title=" + ""
	if request.Data.Response.Demographics.SourceType == "OUTPATIENT" {
		file.Obr.Ptype = "ptype=" + "OP"
	} else {
		file.Obr.Ptype = "ptype=" + "IP"
	}
	bd, err := time.Parse("2006-01-02", request.Data.Response.Demographics.Patient.DateOfBirth)
	exception.SendLogIfErorr(err, "103")
	file.Obr.BirthDt = "birth_dt=" + bd.Format("200601021504")
	if request.Data.Response.Demographics.Patient.Gender == "MALE" {
		file.Obr.Sex = "sex=" + "1"
	} else {
		file.Obr.Sex = "sex=" + "2"
	}
	file.Obr.Ono = "ono=" + request.Data.Response.NoOrder
	file.Obr.Lno = "lno=" + request.Data.Response.Demographics.RegNumber
	rd, err := time.Parse("2006-01-02T15:04:05.000Z", request.Data.Response.Demographics.RegistrationDate)
	if err != nil {
		utils.SendMessage("LINE 109\n" + " Log Type: Error\n" + "Error: \n" + err.Error() + "\n")
	}
	file.Obr.RequestDt = "request_dt=" + rd.Format("200601021504")
	sd, err := time.Parse("2006-01-02T15:04:05.000Z", request.Data.Response.Demographics.CollectDate)
	exception.SendLogIfErorr(err, "117")
	file.Obr.SpecimenDt = "speciment_dt=" + sd.Format("200601021504")
	file.Obr.Source = "source=" + request.Data.Response.Demographics.SourceName + "|" + request.Data.Response.Demographics.SourceID
	file.Obr.Clinician = "clinician=" + request.Data.Response.Demographics.DoctorName + "|" + request.Data.Response.Demographics.DoctorID
	if request.Data.Response.Demographics.IsCyto {
		file.Obr.Priority = "priority=CITO"
	} else {
		file.Obr.Priority = "priority=NON CITO"
	}
	file.Obr.Pstatus = "pstatus=" + request.Data.Response.Demographics.PartnerName + "|" + request.Data.Response.Demographics.PartnerID
	file.Obr.Visitno = "visitno=" + request.Data.Response.Demographics.VisitNumber
	file.Obr.OrderTestID = "order_test_id=" + request.Data.Response.Demographics.OrderTestID
	file.Obr.Comment = "comment=" + request.Data.Response.Demographics.Diagnose
	onoFileName := request.Data.Response.NoOrder
	var obxs []model.OBX

	for _, value := range request.Data.Response.Examinations {
		for _, panel := range value.Children {
			if len(panel.Children) > 0 {
				for _, test := range panel.Children {
					var testStr string
					if panel.PanelID == 0 {
						testStr = panel.TestName + "|" + panel.TestName
					} else {
						panelRes := helper.SearchExaminationsByPanelID(request.Data.Response.Examinations, int(panel.PanelID))
						if panelRes == nil {
							testStr = panel.TestName + "|" + panel.TestName
						} else {
							testStr = panelRes.TestName + "|" + panel.TestName
						}
					}
					aliasName := helper.GetAliasName(test.ValidatedBy)
					testStr += "|" + test.TestName + "|" + test.ExamValue + "|" + test.UnitName + "|" + test.NormalValueText + "|" + test.ExamValueFlag + "|" + aliasName + "^" + test.ValidatedBy + "|" + "ALL^ALL" + "|" + panel.Comment
					obxs = append(obxs, model.OBX{Item: testStr})
				}
			} else {
				var testStr string
				if panel.PanelID == 0 {
					testStr = panel.TestName + "|" + panel.TestName
				} else {
					panelRes := helper.SearchExaminationsByPanelID(request.Data.Response.Examinations, int(panel.PanelID))

					if panelRes == nil {
						testStr = panel.TestName + "|" + panel.TestName
					} else {
						testStr = panelRes.TestName + "|" + panel.TestName
					}
				}
				var aliasName string
				var sep string
				if panel.ValidatedBy == "" {
					aliasName = ""
					sep = ""
				} else {
					aliasName = helper.GetAliasName(panel.ValidatedBy)
					sep = "^"
				}
				testStr += "|" + panel.TestName + "|" + panel.ExamValue + "|" + panel.UnitName + "|" + panel.NormalValueText + "|" + panel.ExamValueFlag + "|" + aliasName + sep + panel.ValidatedBy + "|" + "ALL^ALL" + "|" + panel.Comment
				obxs = append(obxs, model.OBX{Item: testStr})
			}
		}
	}
	var obxResult []string
	for i, item := range obxs {
		obxs[i].Item = "obx" + strconv.Itoa(i+1) + "=" + item.Item
		obxResult = append(obxResult, obxs[i].Item)
	}
	file.Obx.Items = obxResult
	file.Obx.Type = "[OBX]"
	fileValues := helper.GetStructValues(file)
	resultFile, fileName, errFile := helper.WriteLineByLine(fileValues, onoFileName)

	sftp.Upload(resultFile.Name(), fileName, file.Obr.Ono, file.Obr.Lno)
	return "File created", errFile
}
