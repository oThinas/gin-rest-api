package routes

import (
	"gin-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest(r *gin.Engine) {
	r.GET("/api/students", controllers.GetAllStudents)
	r.GET("/api/students/:id", controllers.GetStudentById)
	r.GET("/api/students/cpf/:cpf", controllers.GetStudentByCPF)
	r.POST("/api/students", controllers.CreateStudent)
	r.PUT("/api/students/:id", controllers.UpdateStudent)
	r.DELETE("/api/students/:id", controllers.DeleteStudentById)

	r.GET("/index", controllers.Index)
	r.NoRoute(controllers.NotFound)
}
