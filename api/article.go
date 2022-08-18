package api

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddArticle 添加文章
func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	code = model.CreateArt(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// QueryArticleInCate 查询分类下的所有文章
func QueryArticleInCate(c *gin.Context) {
	cid, _ := strconv.Atoi(c.Query("cid"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "0"))
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pagenum", "0"))

	if pageNum == 0 {
		pageSize = -1
	}
	if pageSize == 0 {
		pageSize = -1
	}

	data, code, total := model.GetArticleListInCate(cid, pageNum, pageSize)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})

}

//QueryArticleInfo 查询单个文章
func QueryArticleInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	data, code := model.QueryArticleInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// QueryArticleList 查询文章列表
func QueryArticleList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "0"))
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "0"))

	if pageNum == 0 {
		pageSize = -1
	}
	if pageSize == 0 {
		pageSize = -1
	}
	data, code, total := model.GetArt(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// UpdateArticle 编辑文章
func UpdateArticle(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&data)
	if err != nil {
		return
	}
	code = model.UpdateArt(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DelArticle 删除文章
func DelArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteArt(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
