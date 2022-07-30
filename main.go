package main

import (
	"ginblog/model"
	"ginblog/routes"
)

func main() {
	//初始化MySql数据库
	model.InitDb()
	routes.InitRouter()
}
