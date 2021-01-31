package security

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"pi-hub/api/utils"
	"pi-hub/services/database/sensors/security"
)

func DoorsRoutesHandler(g *gin.RouterGroup) {
	g.POST("/", func(c *gin.Context) {
		DoorCreateHandler(c)
	})
}

func DoorCreateHandler(c *gin.Context) {
	var inputErrors []utils.InputError
	var json security.Door

	if err := c.ShouldBindJSON(&json); err != nil {
		inputErrors = append(inputErrors, utils.InputError{
			Message: "Error binding json input",
		})
		utils.InputErrorSender(c, inputErrors)
		return
	}

	validate := validator.New()
	err := validate.Struct(json)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Namespace() {
			case "Door.Type":
				if err.Tag() == "required" {
					inputErrors = append(inputErrors, utils.InputError{
						Origin:  "type",
						Message: "Sensor type is required",
					})
				} else if err.Tag() == "oneof" {
					inputErrors = append(inputErrors, utils.InputError{
						Origin:  "type",
						Message: "Invalid type, must be a type of door.lock or door.open",
					})
				}
				break
			case "Door.Name":
				if err.Tag() == "required" {
					inputErrors = append(inputErrors, utils.InputError{
						Origin:  "name",
						Message: "Sensor name is required",
					})
				}
				break
			case "Door.MAC":
				if err.Tag() == "mac" {
					inputErrors = append(inputErrors, utils.InputError{
						Origin:  "mac",
						Message: "Must be valid MAC address",
					})
				}
				break
			default:
				inputErrors = append(inputErrors, utils.InputError{
					Message: "Unknown error",
				})
			}
		}

		utils.InputErrorSender(c, inputErrors)
		return
	}

	id, err := security.AddDoorSensorDoc(json)
	if err != nil {
		inputErrors = append(inputErrors, utils.InputError{
			Message: err.Error(),
		})
		utils.InputErrorSender(c, inputErrors)
		return
	}

	var d []utils.ApiData

	d = append(d, utils.ApiData{
		"id": id,
	})

	utils.SuccessSender(c, &d)
}
