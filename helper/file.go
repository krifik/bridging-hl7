package helper

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/k0kubun/pp"
	"github.com/krifik/bridging-hl7/config"
	"github.com/krifik/bridging-hl7/exception"
	"github.com/krifik/bridging-hl7/model"
	"github.com/krifik/bridging-hl7/utils"
	"github.com/pkg/sftp"
	amqp "github.com/rabbitmq/amqp091-go"
)

// GetContent returns a map of key-value pairs parsed from a file at the given URL.
//
// url: string representing the file path.
//
// Returns a map[string]interface{} with keys and values parsed from the file.

func GetContent(url string) map[string]interface{} {
	file, fileError := os.OpenFile(url, os.O_RDONLY, 0644)
	var text = make([]byte, 1024)
	if fileError != nil {
		utils.SendMessage("LINE 26\n" + "Log Type: Error\n" + "Error: \n" + fileError.Error() + "\n")
	}
	defer file.Close()
	file.Read(text)
	str := strings.NewReader(string(text))
	r := bufio.NewReader(str)

	var results = make(map[string]interface{})

	for {
		token, _, err := r.ReadLine()
		if len(token) > 0 {
			splittedStr := strings.Split(string(token), "=")

			if len(splittedStr) == 2 {
				if splittedStr[1] != "" && splittedStr[0] != "" {
					key := splittedStr[0]
					val := splittedStr[1]
					results[key] = val
				}
			}
		}

		if err != nil {
			break
		}
	}
	return results
}
func GetLatestFileName(dir string) string {
	godotenv.Load()

	latestFile, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		log.Println(err)
	}

	if len(latestFile) == 0 {
		log.Println("No files found in ORDERDIR")
	}

	latestFileName := filepath.Base(latestFile[len(latestFile)-1])
	return latestFileName
}

func RemoveAlphabet(str string) string {

	filteredStr := strings.Builder{}
	for _, r := range str {
		if !unicode.IsLetter(r) {
			filteredStr.WriteRune(r)
		}
	}
	return filteredStr.String()
}

func WriteLineByLine(data []string, fileName string) (*os.File, string, error) {
	err := godotenv.Load()
	if err != nil {
		utils.SendMessage("LINE 86 \nLog Type: Error\n" + "Error: \n" + err.Error() + "\n")
	}
	dir := os.Getenv("RESULTDIR")
	file, err := os.Create(dir + "/" + fileName + ".txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a new buffered writer to write to the file
	writer := bufio.NewWriter(file)

	// Write each line to the file

	for _, line := range data {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}

	// Flush the buffer to ensure all data has been written to the file
	err = writer.Flush()
	if err != nil {
		panic(err)
	}
	return file, fileName, nil
}

func GetStructValues(s interface{}) []string {
	// Check if the argument is a struct before calling reflect.ValueOf on it
	if reflect.TypeOf(s).Kind() != reflect.Struct {
		return []string{}
	}

	// Get the value of the struct using reflection
	value := reflect.ValueOf(s)

	// Create a slice to hold the values of the struct fields as strings
	var values []string

	// Iterate over the struct fields and append their values as strings to the slice
	for i := 0; i < value.NumField(); i++ {
		fieldValue := value.Field(i)

		// If the field is a slice of structs, iterate over the slice and call GetStructValues on each struct
		if fieldValue.Kind() == reflect.Slice {
			for j := 0; j < fieldValue.Len(); j++ {
				values = append(values, fieldValue.Index(j).Interface().(string))
			}
		} else if fieldValue.Kind() == reflect.Struct {
			// If the field is a nested struct, recursively call GetStructValues on it
			nestedValues := GetStructValues(fieldValue.Interface())
			values = append(values, nestedValues...)
		} else {
			// Convert the field value to a string and append it to the slice
			values = append(values, fmt.Sprint(fieldValue.Interface()))
		}
	}

	return values
}

func GetAliasName(name string) string {
	words := strings.Split(name, " ")
	var aliasNameSlice []string
	if name != "" {
		for _, w := range words {
			aliasNameSlice = append(aliasNameSlice, string(w[0]))
		}
		aliasName := strings.Join(aliasNameSlice, "")
		return aliasName
	}
	return ""
}

func SearchExaminationsByPanelID(exams []model.Examinations, targetPanelID int) *model.Children {
	for _, exam := range exams {
		if len(exam.Children) > 0 {
			for _, child := range exam.Children {
				if child.TestID == targetPanelID {
					return &child
				} else if len(child.Children) > 0 {
					if result := SearchExaminationsByPanelID([]model.Examinations{{Children: child.Children}}, targetPanelID); result != nil {
						return result
					}
				}
			}
		}
	}
	return nil
}
func GetContentSftpFile(fileName string, client *sftp.Client) model.Json {
	file, fileError := client.Open(os.Getenv("SFTP_ORDER_DIR") + "/" + fileName)
	var text = make([]byte, math.MaxInt32)
	if fileError != nil {
		utils.SendMessage("LINE 26\n" + "Log Type: Error\n" + "Error: \n" + fileError.Error() + "\n")
	}
	defer file.Close()
	file.Read(text)
	str := strings.NewReader(string(text))
	r := bufio.NewReader(str)

	var results = make(map[string]interface{})

	for {
		token, _, err := r.ReadLine()
		if len(token) > 0 {
			splittedStr := strings.Split(string(token), "=")
			if len(splittedStr) == 2 {
				if splittedStr[1] != "" && splittedStr[0] != "" {
					key := splittedStr[0]
					val := splittedStr[1]
					results[key] = val
				}
			}
		}

		if err != nil {
			break
		}
	}
	json := utils.TransformToRightJson(results)
	return json
}
func SendToAPI(fileContent model.Json) {
	var client fiber.Client
	a := client.Post(os.Getenv("API_EXTERNAL"))
	if fileContent.OrderJson.NoOrder != "" {
		a.ContentType("application/json")
		pp.Println("FILE CONTENT" + fileContent.OrderJson.NoOrder)
		a.JSON(fileContent)
		pp.Println("SENDING DATA WITH ORDER NUMBER : " + fileContent.OrderJson.NoOrder)
		var data interface{}
		code, _, errs := a.Struct(&data) // ...
		var slices []string
		for _, v := range errs {
			slices = append(slices, v.Error())
		}
		pp.Println(code)
		_, err := json.MarshalIndent(fileContent, "", "    ")
		exception.SendLogIfErorr(err, "75")
		// utils.SendMessage("SENT JSON" + string(b))

		errsStr := strings.Join(slices[:], "\n")
		// utils.SendMessage(time.Now().Format("02-01-2006 15:04:05") + " \n\nLog Type: Info \n\n" + "API Request Message : \n\n" + "Code " + strconv.Itoa(code) + " 👍")
		if len(errs) > 0 {
			utils.SendMessage("LINE 73 \nAPI KE MEDQLAB SERVICE ERROR\n" + "Log Type: Error\n\n" + "Error: \n" + errsStr + "\n")
		}
	}
}
func GetPDF(orderNumber string, labNumber string) (*os.File, error) {
	err := godotenv.Load()
	exception.SendLogIfErorr(err, "247")
	var xCons string = os.Getenv("X-CONS")
	var xSign string = os.Getenv("X-SIGN")
	labNumber = strings.Split(labNumber, "=")[1]
	apiUrl := os.Getenv("API_EXTERNAL") + "/api/v1/getResult/pdf"
	// Send an HTTP GET request to the API endpoint for the PDF file.
	req, err := http.NewRequest("GET", apiUrl, nil)
	req.URL.RawQuery = "no_laboratorium=" + labNumber + "&language=id" + "&paper=a4" + "plain=false"

	if err != nil {
		return nil, err
	}
	req.Header.Set("x-cons", xCons)
	req.Header.Set("x-sign", xSign)

	// Send the HTTP request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	orderNumber = strings.Split(orderNumber, "=")[1]
	file, err := os.Create(os.Getenv("RESULTDIR") + orderNumber + ".pdf")
	if err != nil {
		return nil, err
	}
	log.Printf("HTTP status: %d", resp.StatusCode)
	log.Printf("Request URL: %s", req.URL)
	for key, values := range req.Header {
		for _, value := range values {
			log.Printf("%s: %s", key, value)
		}
	}
	// Write the response body to the file.
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		utils.SendMessage("LINE 288 \nAPI KE MEDQLAB SERVICE ERROR\n" + "Log Type: Error\n\n" + "Error: \n" + "Status Code " + strconv.Itoa(resp.StatusCode))
		return nil, err
	}

	return file, nil
}

func SendJsonToRabbitMQ(request model.Json) error {

	ch, conn := config.InitializedRabbitMQ()
	defer ch.Close()
	defer conn.Close()

	jsonData, errJson := json.Marshal(request)
	// pp.Print(string(jsonData))
	exception.SendLogIfErorr(errJson, "22")
	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonData,
	}
	if request.OrderJson.NoOrder != "" {
		scheme := pp.ColorScheme{
			// Integer: pp.Green | pp.Bold,
			Float:  pp.Black | pp.BackgroundWhite | pp.Bold,
			String: pp.Green,
		}

		// Register it for usage
		pp.SetColorScheme(scheme)
		pp.Println("SEND DATA TO RABBITMQ WITH NO ORDER" + request.OrderJson.NoOrder)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		defer ch.Close()
		// Attempt to publish a message to the queue.
		if err := ch.PublishWithContext(
			ctx,
			"",               // exchange
			"bridging_order", // queue name
			false,            // mandatory
			false,            // immediate
			message,          // message to publish
		); err != nil {
			return err
		}
	}

	return nil
}
