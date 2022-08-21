package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	AccessKey  string
	SecretKey  string
	Bucket     string
	QiniuSever string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误", err)
	}
	LoadServer(file)
	LoadData(file)
	QiNiuSever(file)

}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":9090")
	JwtKey = file.Section("server").Key("JwtKey").MustString("frontendWang")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("127.0.0.1")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassword = file.Section("database").Key("DbPassword").MustString("wmq12138")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}

func QiNiuSever(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").MustString("uTqVVr28kQngc_29kqcmaDuFUuRY7uhKs2-LaiI6")
	SecretKey = file.Section("qiniu").Key("SecretKey").MustString("7yvNZ8ATw7rztceSb0G5YS9pL_Trz57LUbf7c88c")
	Bucket = file.Section("qiniu").Key("Bucket").MustString("blog-frontendwang")
	QiniuSever = file.Section("qiniu").Key("QiniuSever").MustString("http://rgtoella9.hd-bkt.clouddn.com/")

}
