package utils

import "github.com/gin-gonic/gin"

// SendErrorResponse mengirimkan respons error
func SendErrorResponse(c *gin.Context, statusCode int, message string, err error) {
	c.JSON(statusCode, map[string]interface{}{
		"message": message,
		"error":   err.Error(),
	})
}

// SendSuccessResponse mengirimkan respons sukses
func SendSuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, map[string]interface{}{
		"message": message,
		"data":    data,
	})
}
