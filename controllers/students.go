package controllers

import (
	"gin-rest-api/database"
	"gin-rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func findStudentById(id string) (models.Student, error) {
	var student models.Student
	result := database.DB.First(&student, id)

	return student, result.Error
}

func findStudentByCPF(cpf string) (models.Student, error) {
	var student models.Student
	result := database.DB.Where("cpf = ?", cpf).First(&student)

	return student, result.Error
}

func respondWithError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

func respondWithSuccess(c *gin.Context, data interface{}, status ...int) {
	responseStatus := http.StatusOK
	if len(status) > 0 {
		responseStatus = status[0]
	}

	c.JSON(responseStatus, data)
}

func GetAllStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)

	respondWithSuccess(c, students)
}

func GetStudentById(c *gin.Context) {
	id := c.Params.ByName("id")
	student, err := findStudentById(id)
	if err != nil {
		respondWithError(c, http.StatusNotFound, "No student found!")
		return
	}

	respondWithSuccess(c, student)
}

func GetStudentByCPF(c *gin.Context) {
	cpf := c.Params.ByName("cpf")
	student, err := findStudentByCPF(cpf)
	if err != nil {
		respondWithError(c, http.StatusNotFound, "No student found!")
		return
	}

	respondWithSuccess(c, student)
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		respondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := student.Validate(); err != nil {
		respondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	database.DB.Create(&student)
	respondWithSuccess(c, student, http.StatusCreated)
}

func UpdateStudent(c *gin.Context) {
	id := c.Params.ByName("id")
	student, err := findStudentById(id)
	if err != nil {
		respondWithError(c, http.StatusNotFound, "No student found!")
		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		respondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := student.Validate(); err != nil {
		respondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	database.DB.Save(&student)
	respondWithSuccess(c, student)
}

func DeleteStudentById(c *gin.Context) {
	id := c.Params.ByName("id")
	student, err := findStudentById(id)
	if err != nil {
		respondWithError(c, http.StatusNotFound, "No student found!")
		return
	}

	database.DB.Delete(&student)
	respondWithSuccess(c, gin.H{
		"message": "Student deleted successfully!",
		"student": student,
	})
}

func Index(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"students": students,
	})
}

func NotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
