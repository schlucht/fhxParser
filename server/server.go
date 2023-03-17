package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Port string
}

func Start() {
	router := gin.Default()

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", ReadUnitName)
	router.GET("/up", ReadUnitDetail)

	// router.GET("/up", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "up.html", gin.H{
	// 		"header": "Unit Procedures",
	// 	})
	// })

	fmt.Println("Server run on localhost:8080")
	router.Run("localhost:8080")
}
