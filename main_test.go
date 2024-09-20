package main

import (
	"bytes"
	"encoding/json"
	"gin-rest-api/controllers"
	"gin-rest-api/database"
	"gin-rest-api/models"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var studentId uint

func newStudentMock() models.Student {
	return models.Student{Name: "Test", CPF: "12345678901", RG: "123456789"}
}

func createStudentMock(student *models.Student) {
	database.DB.Create(student)
	studentId = student.ID
}

func deleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student, studentId)
}

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/api/students", controllers.GetAllStudents)
	r.GET("/api/students/cpf/:cpf", controllers.GetStudentByCPF)
	r.GET("/api/students/:id", controllers.GetStudentById)
	r.POST("/api/students", controllers.CreateStudent)
	r.PUT("/api/students/:id", controllers.UpdateStudent)
	r.DELETE("/api/students/:id", controllers.DeleteStudentById)

	return r
}

func TestGetAllStudents(t *testing.T) {
	database.Connect()
	r := SetupRouter()
	student := newStudentMock()
	createStudentMock(&student)
	defer deleteStudentMock()

	request, _ := http.NewRequest("GET", "/api/students", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)

	var students []models.Student
	err := json.Unmarshal(response.Body.Bytes(), &students)
	assert.NoError(t, err)
	assert.Len(t, students, 1)

	assert.Equal(t, student.Name, students[0].Name)
	assert.Equal(t, student.CPF, students[0].CPF)
	assert.Equal(t, student.RG, students[0].RG)
}

func TestGetStudentByCPF(t *testing.T) {
	database.Connect()
	r := SetupRouter()
	student := newStudentMock()
	createStudentMock(&student)
	defer deleteStudentMock()

	request, _ := http.NewRequest("GET", "/api/students/cpf/12345678901", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)

	var studentResponse models.Student
	err := json.Unmarshal(response.Body.Bytes(), &studentResponse)
	assert.NoError(t, err)

	assert.Equal(t, student.Name, studentResponse.Name)
	assert.Equal(t, student.CPF, studentResponse.CPF)
	assert.Equal(t, student.RG, studentResponse.RG)
}

func TestGetStudentById(t *testing.T) {
	database.Connect()
	r := SetupRouter()
	student := newStudentMock()
	createStudentMock(&student)
	defer deleteStudentMock()

	request, _ := http.NewRequest("GET", "/api/students/"+strconv.FormatUint(uint64(studentId), 10), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)

	var studentResponse models.Student
	err := json.Unmarshal(response.Body.Bytes(), &studentResponse)
	assert.NoError(t, err)

	assert.Equal(t, student.Name, studentResponse.Name)
	assert.Equal(t, student.CPF, studentResponse.CPF)
	assert.Equal(t, student.RG, studentResponse.RG)
}

func TestDeleteStudentById(t *testing.T) {
	database.Connect()
	r := SetupRouter()
	student := newStudentMock()
	createStudentMock(&student)

	request, _ := http.NewRequest("DELETE", "/api/students/"+strconv.FormatUint(uint64(studentId), 10), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)

	var responseBody map[string]interface{}
	err := json.Unmarshal(response.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	assert.Equal(t, "Student deleted successfully!", responseBody["message"])

	deletedStudent := responseBody["student"].(map[string]interface{})
	assert.Equal(t, student.Name, deletedStudent["name"])
	assert.Equal(t, student.CPF, deletedStudent["cpf"])
	assert.Equal(t, student.RG, deletedStudent["rg"])
}

func TestUpdateStudent(t *testing.T) {
	database.Connect()
	r := SetupRouter()
	student := newStudentMock()
	createStudentMock(&student)
	defer deleteStudentMock()

	studentUpdated := student
	studentUpdated.Name = "Updated"
	studentUpdatedJson, _ := json.Marshal(studentUpdated)
	request, _ := http.NewRequest(
		"PUT",
		"/api/students/"+strconv.FormatUint(uint64(studentId), 10),
		bytes.NewBuffer(studentUpdatedJson),
	)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)

	var studentResponse models.Student
	err := json.Unmarshal(response.Body.Bytes(), &studentResponse)
	assert.NoError(t, err)

	assert.Equal(t, studentUpdated.Name, studentResponse.Name)
	assert.Equal(t, studentUpdated.CPF, studentResponse.CPF)
	assert.Equal(t, studentUpdated.RG, studentResponse.RG)
}

func TestCreateStudent(t *testing.T) {
	database.Connect()
	r := SetupRouter()
	student := newStudentMock()

	studentJson, _ := json.Marshal(student)
	request, _ := http.NewRequest(
		"POST",
		"/api/students",
		bytes.NewBuffer(studentJson),
	)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusCreated, response.Code)

	var studentResponse models.Student
	err := json.Unmarshal(response.Body.Bytes(), &studentResponse)
	assert.NoError(t, err)

	assert.Equal(t, student.Name, studentResponse.Name)
	assert.Equal(t, student.CPF, studentResponse.CPF)
	assert.Equal(t, student.RG, studentResponse.RG)
}
