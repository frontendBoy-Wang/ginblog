package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Success 成功的返回
func Success(c *gin.Context, data gin.H, msg string) {
	Response(c, http.StatusOK, 200, data, msg)

}

// Fail 失败的返回
func Fail(c *gin.Context, data gin.H, msg string) {
	Response(c, http.StatusOK, 500, data, msg)
}

// Response 通用返回
func Response(c *gin.Context, httpStatus, code int, data gin.H, msg string) {
	c.JSON(httpStatus, gin.H{"status": code, "msg": msg, "data": data})
}
