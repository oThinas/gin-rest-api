package main

import (
	"encoding/json"
	"gin-rest-api/controllers"
	"gin-rest-api/database"
	"gin-rest-api/models"
	"net/http"
	"net/http/httptest"
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
