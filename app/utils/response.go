package utils

import (
	"github.com/gin-gonic/gin"
)

//JSONresponse is a generic function to response
func JSONresponse(ctx *gin.Context, status int, message interface{}) {
	ctx.JSON(status, gin.H{
		"message": message,
	})
}
