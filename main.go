package main

import (
	"ginblog/model"
	"ginblog/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.ForceConsoleColor()
	//初始化MySql数据库
	model.InitDb()
	routes.InitRouter()
}
