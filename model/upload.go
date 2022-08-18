package model

import (
	"ginblog/utils"
	"mime/multipart"
	//"github.com/qiniu/api/v7/"
)

var (
	AccessKey = utils.AccessKey
	ScretKey  = utils.SecretKey
	Bucket    = utils.Bucket
	ImgUrl    = utils.QiniuSever
)

func UploadFile(file multipart.File, fileSize int64) (string, int) {
	return "", 0
}
