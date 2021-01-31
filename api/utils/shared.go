package utils

import "github.com/gin-gonic/gin"

type (
	ApiReturn struct {
		//data []interface{} `json:"data"`
		ok     bool          `json:"ok"`
		errors []interface{} `json:"errors"`
		data   *[]ApiData    `json:"data"`
	}

	ApiData map[string]interface{}
)

func ConvertApiReturnToJSON(i ApiReturn) map[string]interface{} {
	return gin.H{
		"data":   &i.data,
		"ok":     i.ok,
		"errors": i.errors,
	}
}
