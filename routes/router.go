package routes

import (
	"ginblog/api"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	router := r.Group("api/")

	r1 := router.Group("v1/")
	{
		//用户模块
		r1.POST("user/add", api.AddUser)       //添加用户
		r1.GET("user/list", api.QueryUserList) //查询用户列表
		r1.PUT("user/:id", api.UpdateUser)     //更新用户信息
		r1.DELETE("user/:id", api.DelUser)     //删除用户

		//分类模块

		//文章模块

	}

	err := r.Run(utils.HttpPort)
	if err != nil {
		return
	}
}
