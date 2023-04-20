package helpers

import "github.com/gin-gonic/gin"

func GetHeader(ctx *gin.Context) string {
	return ctx.Request.Header.Get("Content-Type")
}
