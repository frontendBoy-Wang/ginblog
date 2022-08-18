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
	AccessKey = file.Section("qiniu").Key("AccessKey").MustString("mysql")
	SecretKey = file.Section("SecretKey").Key("SecretKey").MustString("mysql")
	Bucket = file.Section("Bucket").Key("Bucket").MustString("mysql")
	QiniuSever = file.Section("QiniuSever").Key("QiniuSever").MustString("mysql")

}
