package api

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

var code int

// CheckUserExist CheckUser 查询用户是否存在
func CheckUserExist(c *gin.Context) {

}

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// QueryUser 查询单个用户
func QueryUser(c *gin.Context) {

}

// QueryUserList 查询用户列表
func QueryUserList(c *gin.Context) {

}

// UpdateUser 编辑用户
func UpdateUser(c *gin.Context) {

}

// DelUser 删除用户
func DelUser(c *gin.Context) {

}
