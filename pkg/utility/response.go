package utility

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context, status int, message string, responseMap map[string]string) {
	c.JSON(status, gin.H{
		"Success": true,
		"Message": message,
		"Data":    responseMap,
	})
}

func FailureResponse(c *gin.Context, status int, message string, err error) {
	c.JSON(status, gin.H{
		"Success": false,
		"Message": message,
		"Error":   err.Error(),
	})
}

func JsonValidationFailure(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"Success": false,
		"Error":   err.Error(),
		"Message": "Could not process the given data",
	})
}
