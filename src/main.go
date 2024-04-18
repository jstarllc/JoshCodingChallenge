package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Light struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Room       string `json:"room"`
	On         bool   `json:"on"`
	Brightness uint8  `json:"brightness"`
}

type LightConcise struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Room string `json:"room"`
}

type LightUpdateReq struct {
	Name       *string `json:"name",omitempty`
	Room       *string `json:"room",omitempty`
	On         *bool   `json:"on",omitempty`
	Brightness *uint8  `json:"brightness",omitempty`
}

// toConcise converts a Light to a LightConcise
func toConcise(l *Light) LightConcise {
	return LightConcise{l.ID, l.Name, l.Room}
}

type Lights = []Light

const (
	defaultLights = `[
        {
            "id": "F06B0ED2-CC50-4EDD-A0F8-5C98A3C9D151",
            "name": "Living Room Lamp Left",
            "room": "Living Room",
            "on": true,
            "brightness": 242
        },
        {
            "id": "AD7E8BB4-FF0A-4B9F-B676-E94BAF878E8F",
            "name": "Living Room Lamp Right",
            "room": "Living Room",
            "on": true,
            "brightness": 242
        },
        {
            "id": "2ED8EB8E-E38B-4F86-8B9D-CBAF37BE7275",
            "name": "Kitchen Overheads",
            "room": "Kitchen",
            "on": false,
            "brightness": 0
        },
        {
            "id": "AF80E5C2-B235-471D-8DF9-490703699EDA",
            "name": "Kitchen Chandelier",
            "room": "Kitchen",
            "on": false,
            "brightness": 0
        },
        {
            "id": "2C85BB59-C136-49F1-A429-F02F52B6C765",
            "name": "Pantry Light",
            "room": "Kitchen",
            "on": false,
            "brightness": 0
        },
        {
            "id": "56A00EC5-E3D5-4A6B-A2CF-6D88C8F6464C",
            "name": "Office Sconce 1",
            "room": "Office",
            "on": true,
            "brightness": 17
        },
        {
            "id": "EDC1B691-A5AF-4524-9B57-80341D90BFA2",
            "name": "Office Sconce 2",
            "room": "Office",
            "on": true,
            "brightness": 35
        },
        {
            "id": "9DF68FEC-06AC-46BF-AB61-57E9D4E963E8",
            "name": "Office Downlights",
            "room": "Office",
            "on": true,
            "brightness": 114
        },
        {
            "id": "1B1920D5-22F1-43AA-9D35-371B2075D33D",
            "name": "Porch Light",
            "room": "Porch",
            "on": false,
            "brightness": 100
        },
        {
            "id": "6FF183B0-710D-40DE-94A8-81584D8AE3E8",
            "name": "String Lights",
            "room": "Back Yard",
            "on": false,
            "brightness": 120
        }
    ]`
)

var (
	lights Lights
)

// getLights godoc
// @Summary Get summary of all lights.
// @Description Get list of all lights in the system.
// @ID get-lights
// @Produce json
// @Success 200 {array} LightConcise
// @Router /lights [get]
func getLights(c *gin.Context) {
	var lightsConcise []LightConcise
	for _, l := range lights {
		lightsConcise = append(lightsConcise, toConcise(&l))
	}
	c.IndentedJSON(http.StatusOK, lightsConcise)
}

// getLightByID godoc
// @Summary Get details about a light.
// @Description Get detailed state of a light in the system.
// @Param				lightID path string true "ID of light"
// @ID get-light-by-id
// @Produce json
// @Success 200 {object} Light
// @Router /lights/{lightID} [get]
func getLightByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of lights, looking for
	// an light whose ID value matches the parameter.
	for _, l := range lights {
		if l.ID == id {
			c.IndentedJSON(http.StatusOK, l)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "light not found"})
}

// addLight godoc
// @Summary Add a light.
// @Description Add a new light to the system.
// @Param				light body Light true "Full state of light to add"
// @ID add-light
// @Produce json
// @Success 200 {object} Light
// @Router /lights [post]
func addLight(c *gin.Context) {
	var newLight Light

	// Call BindJSON to bind the received JSON to
	// newLight.
	if err := c.BindJSON(&newLight); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	// Add the new light to the slice.
	lights = append(lights, newLight)
	c.IndentedJSON(http.StatusCreated, newLight)
}

// deleteLightByID godoc
// @Summary Delete a light.
// @Description Remove a light from the system by ID.
// @Param				lightID path string true "ID of light"
// @ID delete-light-by-id
// @Produce json
// @Success 204
// @Router /lights/{lightID} [delete]
func deleteLightByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of lights, looking for
	// an light whose ID value matches the parameter.
	var newLights Lights
	for _, l := range lights {
		if l.ID != id {
			newLights = append(newLights, l)
		}
	}
	if len(lights) == len(newLights) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "light not found"})
	}
	lights = newLights
	c.Status(http.StatusNoContent)
}

// updateLightByID godoc
// @Summary Update a light.
// @Description Update the state of a light in the system by ID.
// @Param				lightID path string true "ID of light"
// @Param				state body LightUpdateReq true "State fields to update"
// @ID update-light-by-id
// @Produce json
// @Success 200 {object} Light
// @Router /lights/{lightID} [put]
func updateLightByID(c *gin.Context) {
	id := c.Param("id")

	var updateReq LightUpdateReq
	if err := c.BindJSON(&updateReq); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	for i, l := range lights {
		if l.ID == id {
			if updateReq.Name != nil {
				l.Name = *updateReq.Name
			}
			if updateReq.Room != nil {
				l.Room = *updateReq.Room
			}
			if updateReq.On != nil {
				l.On = *updateReq.On
			}
			if updateReq.Brightness != nil {
				l.Brightness = *updateReq.Brightness
			}
			lights[i] = l
			c.IndentedJSON(http.StatusOK, l)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "light not found"})
}

// @title Josh.ai Coding Challenge API
// @version	1.0
// @description API for interacting with lighting hub simulator for use in the Josh.ai Backend Engineer Coding Challenge.
// @host 	localhost:8080
// @BasePath /
func main() {
	json.Unmarshal([]byte(defaultLights), &lights)

	router := gin.Default()
	router.StaticFS("/static", http.Dir("./public_html"))
	router.LoadHTMLFiles("./public_html/index.html")

	router.GET("/lights", getLights)
	router.GET("/lights/:id", getLightByID)
	router.POST("/lights", addLight)
	router.DELETE("/lights/:id", deleteLightByID)
	router.PUT("/lights/:id", updateLightByID)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run("localhost:8080")
}
