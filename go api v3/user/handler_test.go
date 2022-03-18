package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/mburaksoran/fiberv2restapi/db"
	"github.com/stretchr/testify/assert"
)

func TestGetHandler(t *testing.T) {
	database, err := db.Connect()
	assert.Nil(t, err)
	repo := NewRepository(database)
	service := NewService(repo)
	handler := NewHandler(service)
	app := fiber.New()
	app.Get("/users/:id", handler.Get)
	req := httptest.NewRequest("GET", fmt.Sprintf("/users/%d", 1), nil)
	resp, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestCreateHandler(t *testing.T) {
	database, err := db.Connect()
	assert.Nil(t, err)
	repo := NewRepository(database)
	service := NewService(repo)
	handler := NewHandler(service)
	app := fiber.New()
	app.Post("/users/", handler.Create)
	f := &Model{Name: "test123", Email: "Test1234@email.com"}
	data, _ := json.Marshal(f)
	reader := bytes.NewReader(data)
	req := httptest.NewRequest(http.MethodPost, "/users/", reader)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 201, resp.StatusCode)
}

func TestUpdateHandler(t *testing.T) {
	database, err := db.Connect()
	assert.Nil(t, err)
	repo := NewRepository(database)
	service := NewService(repo)
	handler := NewHandler(service)
	app := fiber.New()
	app.Put("/users/:id", handler.Update)
	f := &Model{Name: "testing_Update", Email: "Test1234_update@email.com"}
	data, _ := json.Marshal(f)
	reader := bytes.NewReader(data)
	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/users/%d", 1), reader)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestDeleteHandler(t *testing.T) {
	database, err := db.Connect()
	assert.Nil(t, err)
	repo := NewRepository(database)
	service := NewService(repo)
	handler := NewHandler(service)
	app := fiber.New()
	app.Delete("/users/:id", handler.Delete)
	req := httptest.NewRequest("DELETE", fmt.Sprintf("/users/%d", 1), nil)
	resp, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

// func TestGetHandler_ReturnErrorFromService(t *testing.T) {
// 	database, err := db.Connect()
// 	assert.Nil(t, err)
// 	repo := NewRepository(database)
// 	service := &MockService{}
// 	service.On("Get", mock.AnythingOfType("uint")).Return(nil, errors.New("Get return hata"))
// 	handler := NewHandler(service)
// 	app := fiber.New()
// 	app.Get("/users/:id", handler.Get)
// 	id, err := repo.Create(Model{Name: "test", Email: "test@eamail.com"})
// 	assert.Nil(t, err)
// 	req := httptest.NewRequest("GET", fmt.Sprintf("/users/%d", id), nil)
// 	resp, err := app.Test(req)
// 	assert.Nil(t, err)
// 	body, err := ioutil.ReadAll(resp.Body)
// 	assert.Nil(t, err)
// 	fmt.Println(string(body))
// 	assert.Equal(t, 400, resp.StatusCode)
// }
