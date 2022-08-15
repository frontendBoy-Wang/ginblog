package routes

import (
	"context"
	"ginblog/api"
	"ginblog/middleware"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	r.Use(middleware.Cros(), middleware.CalcTimeMiddleWare())
	router := r.Group("api/")

	r1 := router.Group("v1/")
	{
		//用户模块
		r1.POST("user/add", api.AddUser)       //添加用户
		r1.GET("user/list", api.QueryUserList) //查询用户列表
		r1.PUT("user/:id", api.UpdateUser)     //更新用户信息
		r1.DELETE("user/:id", api.DelUser)     //删除用户

		//分类模块
		r1.POST("category/add", api.AddCate)       //添加分类
		r1.GET("category/list", api.QueryCateList) //查询分类列表
		r1.PUT("category/:id", api.UpdateCate)     //更新分类信息
		r1.DELETE("category/:id", api.DelCate)     //删除分类
		//文章模块

	}

	srv := &http.Server{
		Addr:    utils.HttpPort,
		Handler: r,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

	//err := r.Run(utils.HttpPort)
	//if err != nil {
	//	return
	//}
}
