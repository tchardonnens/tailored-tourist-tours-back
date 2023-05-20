package main

import (
	"net/http"
	"t3/m/v2/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	r.Use(cors.New(config))

	models.ConnectDatabase()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"message": "pong",
			})
	})

	r.POST("/api/v1/parameters", func(c *gin.Context) {
		var parameters models.Parameters
		err := c.BindJSON(&parameters)
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"message": "bad request",
				})
			return
		}
		c.JSON(
			http.StatusOK,
			gin.H{
				"message":  "parameters created",
				"Location": parameters.Location,
				"Days":     parameters.Days,
				"Types":    parameters.Types,
			})
	})

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
