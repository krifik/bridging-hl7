package service

import (
	"bridging-hl7/model"
	"bridging-hl7/repository"
	"bridging-hl7/utils"
	"math/rand"
	"os"
	"strconv"
	"time"

	helper "bridging-hl7/helper"

	"github.com/joho/godotenv"
	"github.com/k0kubun/pp"
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
	// fileName := "LAB-20230504-00001.txt"
	// dir := os.Getenv("ORDERDIR")
	// file, fileError := os.OpenFile(dir+"/"+fileName, os.O_RDONLY, 0644)
	pp.Print(url)
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
	return "asd"
}

// CreateFileResult creates a new file result.
//
// JSONRequest is the request model.
// Returns a string.

func (f *FileServiceImpl) CreateFileResult(request model.JSONRequest) (string, error) {
	var file model.FileResult
	rand.Seed(time.Now().UnixNano())
	pp.Print(request)
	// Generate a random integer between 0 and 999999
	num := rand.Intn(999999)
	file.Msh.MessageDT = "message_dt=" + time.Now().Format("200601021504")
	str := strconv.Itoa(num)
	file.Msh.Type = "[MSH]"
	file.Msh.MessageID = "message_id=" + "TDR-3000" + str
	file.Msh.Version = "version=" + "2.3"
	file.Obr.Type = "[OBR]"
	file.Obr.PID = "pid=" + request.Response.Demographics.Patient.MRN
	file.Obr.Apid = "apid=" + ""
	file.Obr.Pname = "pname=" + request.Response.Demographics.Patient.FullName
	file.Obr.Pidentityno = "pidentityno=" + request.Response.Demographics.Patient.IDNumber
	file.Obr.Pmobileno = "pmobileno=" + request.Response.Demographics.Patient.PhoneNumber
	file.Obr.Street = "street=" + request.Response.Demographics.Patient.Address
	file.Obr.Title = "title=" + ""
	if request.Response.Demographics.SourceType == "OUTPATIENT" {
		file.Obr.Ptype = "ptype=" + "OP"
	} else {
		file.Obr.Ptype = "ptype=" + "IP"
	}
	file.Obr.BirthDt = "birth_dt=" + request.Response.Demographics.Patient.DateOfBirth
	if request.Response.Demographics.Patient.Gender == "MALE" {
		file.Obr.Sex = "sex=" + "1"
	} else {
		file.Obr.Sex = "sex=" + "2"
	}
	file.Obr.Ono = "ono=" + request.Response.NoOrder
	file.Obr.Lno = "lno=" + request.Response.Demographics.RegNumber
	rd, err := time.Parse("2006-01-02T15:04:05.000Z", request.Response.Demographics.RegistrationDate)
	if err != nil {
		utils.SendMessage("LINE 109\n" + " Log Type: Error\n" + "Error: \n" + err.Error() + "\n")
	}
	file.Obr.RequestDt = "request_dt=" + rd.Format("200601021504")
	file.Obr.SpecimenDt = "speciment_dt=" + request.Response.Demographics.CollectDate
	file.Obr.Source = "source=" + request.Response.Demographics.SourceName + "|" + request.Response.Demographics.SourceID
	file.Obr.Clinician = "clinician=" + request.Response.Demographics.DoctorName + "|" + request.Response.Demographics.DoctorID
	if request.Response.Demographics.IsCyto {
		file.Obr.Priority = "priority=CITO"
	} else {
		file.Obr.Priority = "priority=NON CITO"
	}
	file.Obr.Pstatus = "pstatus=" + request.Response.Demographics.PartnerName + "|" + request.Response.Demographics.PartnerID
	file.Obr.Visitno = "visitno=" + request.Response.Demographics.VisitNumber
	file.Obr.OrderTestID = "order_test_id=" + request.Response.Demographics.OrderTestID
	file.Obr.Comment = "comment=" + request.Response.Demographics.Diagnose
	onoFileName := request.Response.NoOrder
	var obxs []model.OBX

	for _, value := range request.Response.Examinations {
		for _, panel := range value.Children {
			if len(panel.Children) > 0 {
				for _, test := range panel.Children {
					testStr := panel.TestName
					testStr += "|" + test.TestName + "|" + "NM" + "|" + test.ExamValue + "|" + test.UnitName + "|" + test.NormalValueText + "|" + test.ExamValueFlag + "|" + test.ValidatedBy
					obxs = append(obxs, model.OBX{Item: testStr})
				}
			} else {
				testStr := panel.TestName
				testStr += "|" + panel.TestName + "|" + "NM" + "|" + panel.ExamValue + "|" + panel.UnitName + "|" + panel.NormalValueText + "|" + panel.ExamValueFlag + "|" + panel.ValidatedBy
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
	errFile := helper.WriteLineByLine(fileValues, onoFileName)
	return "File created", errFile
}
