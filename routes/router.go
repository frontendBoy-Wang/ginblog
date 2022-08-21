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
	r.Use(middleware.Cros())
	Auth := r.Group("api/v1")
	Auth.Use(middleware.JwtToken())
	{
		//用户模块
		Auth.POST("user/add", api.AddUser)   //添加用户
		Auth.PUT("user/:id", api.UpdateUser) //更新用户信息
		Auth.DELETE("user/:id", api.DelUser) //删除用户

		//分类模块
		Auth.POST("category/add", api.AddCate)   //添加分类
		Auth.PUT("category/:id", api.UpdateCate) //更新分类信息
		Auth.DELETE("category/:id", api.DelCate) //删除分类

		//文章模块
		Auth.POST("article/add", api.AddArticle)   //添加文章
		Auth.PUT("article/:id", api.UpdateArticle) //更新文章
		Auth.DELETE("article/:id", api.DelArticle) //删除文章

		//文件上传
		//Auth.POST("upload", api.Upload) //七牛云文件上传

	}

	router := r.Group("api/v1")
	{
		router.GET("user/list", api.QueryUserList) //查询用户列表

		router.GET("category/list", api.QueryCateList)         //查询分类列表
		router.GET("category", api.QueryCateInfo)              //查询单个分类
		router.GET("category/article", api.QueryArticleInCate) //查询分类下的文章

		router.GET("article/list", api.QueryArticleList) //查询文章列表
		router.GET("article", api.QueryArticleInfo)      //查询单个文章
		router.POST("login", api.Login)                  //登陆

		router.POST("upload", api.Upload) //七牛云文件上传
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
