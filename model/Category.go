package model

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// CheckCate 查询分类是否存在
func CheckCate(name string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS
}

// CreateCate 添加分类
func CreateCate(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//GetCateInfo 查询单个分类
func GetCateInfo(id int) (Category, int) {
	var cate Category
	err = db.Where("id = ?", id).First(&cate).Error
	if err != nil {
		return cate, errmsg.ERROR
	}
	return cate, errmsg.SUCCESS
}

// GetCate 查询分类列表
func GetCate(pageSize int, pageNum int) []Category {
	var cate []Category
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cate
}

// UpdateCate  编辑分类
func UpdateCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err := db.Model(&cate).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteCate  删除
func DeleteCate(id int) int {
	var cate Category
	err := db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
