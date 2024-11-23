package utils

import "github.com/gin-gonic/gin"

// SendErrorResponse mengirimkan respons error
func SendErrorResponse(c *gin.Context, statusCode int, message string, err error) {
	var errorDetails interface{}
	if err != nil {
		errorDetails = err.Error()
	} else {
		errorDetails = nil
	}

	c.JSON(statusCode, map[string]interface{}{
		"message": message,
		"error":   errorDetails,
	})
}

// SendSuccessResponse mengirimkan respons sukses
func SendSuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, map[string]interface{}{
		"message": message,
		"data":    data,
	})
}
