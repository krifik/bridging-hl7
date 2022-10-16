package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"mangojek-backend/config"
	"mangojek-backend/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func truncateUser(db *gorm.DB) {
	db.Raw("TRUNCATE users")
}
func TestUserControllerInsertSuccess(t *testing.T) {
	configuration := config.NewConfiguration("../.env.test")
	db := config.NewPostgresDatabase(configuration)
	truncateUser(db)
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	createUserRequest := model.CreateUserRequest{
		Name: "Test",
		// LastName:  "Test",
		Email:    "test@email.com",
		Password: "test",
	}

	requestBody, _ := json.Marshal(createUserRequest)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/users", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	res, _ := app.Test(request)
	webResp := model.WebResponse{}
	respBody, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(respBody, &webResp)
	assert.Equal(t, 200, webResp.Code)
	assert.Equal(t, "OK", webResp.Status)
	jsonData, _ := json.Marshal(webResp.Data)
	createUserResponse := model.CreateUserResponse{}
	json.Unmarshal(jsonData, &createUserResponse)
	assert.NotNil(t, createUserResponse.Id)
	assert.Equal(t, createUserRequest.Email, createUserResponse.Email)
<<<<<<< HEAD
	// assert.Equal(t, createUserRequest.FirstName, createUserRequest.FirstName)
=======
>>>>>>> 16b29b5c3d7e09108329e85929ba6122e42b441c
	assert.Equal(t, createUserRequest.Name, createUserResponse.Name)
	assert.Equal(t, createUserRequest.Password, createUserResponse.Password)

}
