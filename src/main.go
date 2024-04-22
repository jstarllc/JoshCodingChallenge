package main

import (
	"flag"
	"fmt"
	"josh-go/josh-coding-challenge/light"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @title Josh.ai Coding Challenge API
// @version	1.0
// @description API for interacting with lighting hub simulator for use in the Josh.ai Backend Engineer Coding Challenge.
// @description Note: Try it Out will only work if the server is running on localhost:8080.
// @host 	localhost:8080
// @schemes http
// @BasePath /
func main() {
	// CLI params
	var port int
	var data string
	flag.IntVar(&port, "port", 8080, "Server port number")
	flag.StringVar(&data, "data", "", "Initial lights JSON data file")
	flag.Parse()

	// Initialize light data
	light.InitLights(data)

	// Initialize server
	router := gin.Default()
	router.Use(cors.Default())

	router.StaticFS("/static", http.Dir("./public_html"))
	router.LoadHTMLFiles("./public_html/index.html")

	router.GET("/lights", light.GetLights)
	router.GET("/lights/:id", light.GetLightByID)
	router.POST("/lights", light.AddLight)
	router.DELETE("/lights/:id", light.DeleteLightByID)
	router.PUT("/lights/:id", light.UpdateLightByID)

	// Host simple GUI for control
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Run the server
	addr := fmt.Sprintf(":%d", port)
	router.Run(addr)
}
