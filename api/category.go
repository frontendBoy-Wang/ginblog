package api

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddCate 添加分类
func AddCate(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCate(data.Name)
	if code == errmsg.SUCCESS {
		model.CreateCate(&data)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		code = errmsg.ERROR_CATENAME_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//QueryCateInfo 查询单个分类
func QueryCateInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	data, code := model.GetCateInfo(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

// QueryCateList 查询分类列表
func QueryCateList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "0"))
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "0"))

	if pageNum == 0 {
		pageSize = -1
	}
	if pageSize == 0 {
		pageSize = -1
	}
	data := model.GetCate(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// UpdateCate 编辑分类
func UpdateCate(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&data)
	if err != nil {
		return
	}
	code = model.CheckCate(data.Name)
	if code == errmsg.SUCCESS {
		model.UpdateCate(id, &data)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DelCate 删除分类
func DelCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteCate(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
