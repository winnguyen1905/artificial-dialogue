package response

import "github.com/gin-gonic/gin"

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, statusCode int, data interface{}) {
    c.JSON(statusCode, response{
		Code:    statusCode,
        Message: msg[statusCode],
        Data:    data,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, response{
        Code:    statusCode,
        Message: msg[statusCode],
        Data:    data,
    })
}