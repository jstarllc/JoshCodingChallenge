package light

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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

type ErrorResp struct {
	Message string `json:"error"`
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

// InitLights initializes lights
func InitLights(l string) {
	if len(l) == 0 {
		json.Unmarshal([]byte(defaultLights), &lights)
		return
	}
	data, err := os.ReadFile(l)
	if err != nil {
		panic("Invalid lights data file")
	}
	json.Unmarshal([]byte(data), &lights)
}

// findLight returns the light with matching ID
func findLight(id string) (Light, error) {
	for _, l := range lights {
		if l.ID == id {
			return l, nil
		}
	}
	return Light{}, fmt.Errorf("light not found")
}

// GetLights godoc
// @Summary Get summary of all lights.
// @Description Get list of all lights in the system.
// @Tags lights
// @ID get-lights
// @Produce json
// @Success 200 {array} LightConcise
// @Router /lights [get]
func GetLights(c *gin.Context) {
	var lightsConcise []LightConcise
	for _, l := range lights {
		lightsConcise = append(lightsConcise, toConcise(&l))
	}
	c.IndentedJSON(http.StatusOK, lightsConcise)
}

// GetLightByID godoc
// @Summary Get details about a light.
// @Description Get detailed state of a light in the system.
// @Tags lights
// @Param				lightID path string true "ID of light"
// @ID get-light-by-id
// @Produce json
// @Success 200 {object} Light
// @Failure 404 {object} ErrorResp "light not found"
// @Router /lights/{lightID} [get]
func GetLightByID(c *gin.Context) {
	id := c.Param("id")
	l, err := findLight(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, ErrorResp{"light not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, l)
}

// AddLight godoc
// @Summary Add a light.
// @Description Add a new light to the system.
// @Tags lights
// @Param				light body Light true "Full state of light to add"
// @ID add-light
// @Produce json
// @Success 201 {object} Light
// @Failure 400 {object} ErrorResp "invalid light data in body"
// @Router /lights [post]
func AddLight(c *gin.Context) {
	var newLight Light

	// Deserialize the light data from body
	if err := c.BindJSON(&newLight); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResp{"invalid light data in body"})
		return
	}

	if _, err := findLight(newLight.ID); err == nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResp{"light already exists with that ID"})
		return
	}

	// Add the new light to our collection
	lights = append(lights, newLight)
	c.IndentedJSON(http.StatusCreated, newLight)
}

// DeleteLightByID godoc
// @Summary Delete a light.
// @Description Remove a light from the system by ID.
// @Tags lights
// @Param				lightID path string true "ID of light"
// @ID delete-light-by-id
// @Produce json
// @Success 204
// @Failure 404 {object} ErrorResp "light not found"
// @Router /lights/{lightID} [delete]
func DeleteLightByID(c *gin.Context) {
	id := c.Param("id")

	// Keep all but the requested light
	var newLights Lights
	for _, l := range lights {
		if l.ID != id {
			newLights = append(newLights, l)
		}
	}
	if len(lights) == len(newLights) {
		c.IndentedJSON(http.StatusNotFound, ErrorResp{"light not found"})
		return
	}
	lights = newLights
	c.Status(http.StatusNoContent)
}

// UpdateLightByID godoc
// @Summary Update a light.
// @Description Update the state of a light in the system by ID.
// @Tags lights
// @Param				lightID path string true "ID of light"
// @Param				state body LightUpdateReq true "State fields to update"
// @ID update-light-by-id
// @Produce json
// @Success 200 {object} Light
// @Failure 400 {object} ErrorResp "invalid fields in body"
// @Failure 404 {object} ErrorResp "light not found"
// @Router /lights/{lightID} [put]
func UpdateLightByID(c *gin.Context) {
	id := c.Param("id")

	var updateReq LightUpdateReq
	if err := c.BindJSON(&updateReq); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResp{"invalid fields in body"})
		return
	}

	var emptyReq LightUpdateReq
	if updateReq == emptyReq {
		c.IndentedJSON(http.StatusBadRequest, ErrorResp{"no valid fields in body"})
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

	c.IndentedJSON(http.StatusNotFound, ErrorResp{"light not found"})
}
