package helper

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"unicode"

	"github.com/krifik/bridging-hl7/model"
	"github.com/krifik/bridging-hl7/utils"

	"github.com/joho/godotenv"
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

func WriteLineByLine(data []string, fileName string) error {
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
	return nil
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
	for _, w := range words {
		aliasNameSlice = append(aliasNameSlice, string(w[0]))
	}
	aliasName := strings.Join(aliasNameSlice, "")
	return aliasName
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
