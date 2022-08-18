package model

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title        string `gorm:"type:varchar(100);not null" json:"title"`
	Cid          int    `gorm:"type:int;not null" json:"cid"`
	Desc         string `gorm:"type:varchar(200)" json:"desc"`
	Content      string `gorm:"type:longtext" json:"content"`
	Img          string `gorm:"type:varchar(100)" json:"img"`
	CommentCount int    `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int    `gorm:"type:int;not null;default:0" json:"read_count"`
}

// CreateArt 添加文章
func CreateArt(data *Article) int {
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetArticleListInCate 查询文章下的所有文章
func GetArticleListInCate(cid int, pageSize int, pageNum int) ([]Article, int, int64) {
	var art []Article
	var total int64
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid=?", cid).Find(&art).Error
	db.Model(&art).Where("cid=?", cid).Count(&total)
	if err != nil {
		return art, errmsg.ERROR_CATENAME_NOT_FOUND, 0
	}
	return art, errmsg.SUCCESS, total
}

// QueryArticleInfo 查询单个文章 todo
func QueryArticleInfo(id int) (Article, int) {
	var art Article
	err = db.Preload("Category").Where("id=?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ARTICLE_NOT_FOUND
	}
	return art, errmsg.SUCCESS
}

// GetArt 查询文章列表
func GetArt(pageSize int, pageNum int) ([]Article, int, int64) {
	var art []Article
	var total int64
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&art).Error
	db.Model(&art).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return art, errmsg.SUCCESS, total
}

// UpdateArt  编辑文章
func UpdateArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err = db.Model(&art).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteArt  删除
func DeleteArt(id int) int {
	var art Article
	err = db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
