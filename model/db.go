package model

import (
	"fmt"
	"ginblog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var (
	db  *gorm.DB
	err error
)

// InitDb 数据库初始化
func InitDb() {
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPort,
		utils.DbName)
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "gormv2_", //表前缀
			SingularTable: true,
		},
		// gorm日志模式：silent
		Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
	})
	if err != nil {
		fmt.Println("连接数据库失败", err)
	}
	d, _ := db.DB()
	//数据库迁移
	err := db.AutoMigrate(
		&User{},
		&Article{},
		&Category{},
	)
	if err != nil {
		fmt.Println("数据库迁移失败", err)
		return
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	d.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	d.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	d.SetConnMaxLifetime(time.Second * 10)

	//defer func(d *sql.DB) {
	//	err := d.Close()
	//	if err != nil {
	//		fmt.Println("数据库关闭失败", err)
	//	}
	//}(d)
}
