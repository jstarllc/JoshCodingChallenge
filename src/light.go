package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Full state of a light
type Light struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Room string `json:"room"`
	On   bool   `json:"on"`
	// Light brightness 0-255
	Brightness uint8 `json:"brightness"`
}

// Summary of a light
type LightConcise struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Room string `json:"room"`
}

// Fields included in a light update request
type LightUpdate struct {
	// Include to set light name
	Name *string `json:"name,omitempty"`
	// Include to set light room
	Room *string `json:"room,omitempty"`
	// Include to set light on/off
	On *bool `json:"on,omitempty"`
	// Include to set light brightness 0-255
	Brightness *uint8 `json:"brightness,omitempty"`
}

type ErrorResp struct {
	Message string `json:"error"`
}

// toConcise converts a Light to a LightConcise
func (l *Light) toConcise() LightConcise {
	return LightConcise{l.ID, l.Name, l.Room}
}

// fromUpdate Populates a light from an update request
func (l *Light) fromUpdate(update LightUpdate) {
	if update.Name != nil {
		l.Name = *update.Name
	}
	if update.Room != nil {
		l.Room = *update.Room
	}
	if update.On != nil {
		l.On = *update.On
	}
	if update.Brightness != nil {
		l.Brightness = *update.Brightness
	}
}

// isValid checks if the light update is valid
func (l *LightUpdate) isValid() bool {
	var empty LightUpdate
	return *l != empty
}

var lights map[string]Light

// InitLights initializes lights
func InitLights(l string) {
	data, err := os.ReadFile(l)
	if err != nil {
		panic("Invalid lights data file")
	}
	json.Unmarshal([]byte(data), &lights)
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
	var ls []LightConcise
	for _, l := range lights {
		ls = append(ls, l.toConcise())
	}
	c.IndentedJSON(http.StatusOK, ls)
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
	l, ok := lights[id]
	if !ok {
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
	var l Light

	// Deserialize the light data from body
	if err := c.BindJSON(&l); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResp{"invalid light data in body"})
		return
	}

	if l.ID == "" {
		l.ID = uuid.New().String()
	}

	// Check whether light exists
	if _, ok := lights[l.ID]; ok {
		c.IndentedJSON(http.StatusBadRequest, ErrorResp{"light already exists with that ID"})
		return
	}

	// Add the new light to our collection
	lights[l.ID] = l
	c.IndentedJSON(http.StatusCreated, l)
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

	// Check whether light exists
	if _, ok := lights[id]; !ok {
		c.IndentedJSON(http.StatusNotFound, ErrorResp{"light not found"})
		return
	}

	// Remove the light
	delete(lights, id)
	c.Status(http.StatusNoContent)
}

// UpdateLightByID godoc
// @Summary Update a light.
// @Description Update the state of a light in the system by ID.
// @Tags lights
// @Param				lightID path string true "ID of light"
// @Param				state body LightUpdate true "State fields to update"
// @ID update-light-by-id
// @Produce json
// @Success 200 {object} Light
// @Failure 400 {object} ErrorResp "invalid fields in body"
// @Failure 404 {object} ErrorResp "light not found"
// @Router /lights/{lightID} [put]
func UpdateLightByID(c *gin.Context) {
	id := c.Param("id")

	// Check whether light exists
	var l Light
	var ok bool
	if l, ok = lights[id]; !ok {
		c.IndentedJSON(http.StatusNotFound, ErrorResp{"light not found"})
		return
	}

	// Deserialize the update data from body
	var updateReq LightUpdate
	if err := c.BindJSON(&updateReq); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResp{"invalid fields in body"})
		return
	}

	// Validate the update fields
	if !updateReq.isValid() {
		c.IndentedJSON(http.StatusBadRequest, ErrorResp{"no valid fields in body"})
		return
	}

	// Update the light
	l.fromUpdate(updateReq)
	lights[id] = l
	c.IndentedJSON(http.StatusOK, l)
}
