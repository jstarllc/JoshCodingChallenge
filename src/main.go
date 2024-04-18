package main

import (
	"josh-go/josh-coding-challenge/light"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Josh.ai Coding Challenge API
// @version	1.0
// @description API for interacting with lighting hub simulator for use in the Josh.ai Backend Engineer Coding Challenge.
// @host 	localhost:8080
// @BasePath /
func main() {
	light.InitLights(nil)

	router := gin.Default()
	router.StaticFS("/static", http.Dir("./public_html"))
	router.LoadHTMLFiles("./public_html/index.html")

	router.GET("/lights", light.GetLights)
	router.GET("/lights/:id", light.GetLightByID)
	router.POST("/lights", light.AddLight)
	router.DELETE("/lights/:id", light.DeleteLightByID)
	router.PUT("/lights/:id", light.UpdateLightByID)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run("localhost:8080")
}
