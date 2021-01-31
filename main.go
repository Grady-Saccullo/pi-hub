package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pi-hub/api/sensors"
)

const PORT = 8080

type Sensor struct {
	Name string `json:name`
	Type string `json:type`
}

func StartServer() {
	router := gin.Default()
	api := router.Group("/api")
	v1 := api.Group("/v1")
	sensors.RoutesHandler(v1.Group("/sensors"))
	//doors.POST("/", func(c *gin.Context) {
	//	var json security.Door
	//	if err := c.ShouldBindJSON(&json); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	//	}
	//
	//	fmt.Println("attempting to create doc")
	//
	//	if err := security.AddDoorSensor(json); err != nil {
	//		fmt.Println(fmt.Sprintf("err: %d", err))
	//		return
	//	}
	//})

	router.Run(fmt.Sprintf(":%d", PORT))
}

func main() {

	StartServer()
}
