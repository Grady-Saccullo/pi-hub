package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuccessSender(c *gin.Context, data *[]ApiData) {
	a := ApiReturn{
		//data,
		data:   data,
		ok:     true,
		errors: nil,
	}

	c.JSON(http.StatusOK, ConvertApiReturnToJSON(a))
}
