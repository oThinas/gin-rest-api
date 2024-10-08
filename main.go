package main

import (
	"gin-rest-api/database"
	"gin-rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/assets", "./assets")
	r.SetTrustedProxies(nil)
	routes.HandleRequest(r)
	r.Run(":8080")
}
