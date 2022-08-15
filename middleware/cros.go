package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Cros() gin.HandlerFunc {
	return func(c *gin.Context) {
		//fmt.Println("[MyLog] 用户ip:", c.ClientIP())
		//fmt.Println("[MyLog] 用户request:", c.Request)
		fmt.Println("跨域中间件")
		c.Next()
	}
}

func CalcTimeMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		// 统计时间
		since := time.Since(start)
		fmt.Println("程序用时：", since)
	}
}
