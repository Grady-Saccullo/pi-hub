package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	InputError struct {
		Origin  string `json:"source,omitempty"`
		Message string `json:"message""`
	}
)

func InputErrorSender(c *gin.Context, inputErrors []InputError) {
	var errs []interface{}
	for _, e := range inputErrors {
		errs = append(errs, e)
	}

	a := ApiReturn{
		data:   nil,
		ok:     false,
		errors: errs,
	}
	c.JSON(http.StatusUnprocessableEntity, ConvertApiReturnToJSON(a))
}
