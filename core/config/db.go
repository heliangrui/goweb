package config

import (
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var GDb *gorm.DB

func InitDB() {
	l := zap.L()
	getPostgres := GetPostgres()
	l.Info("开始加载数据库")
	dsn := "host=" + getPostgres.hostname + " user=" + getPostgres.username + " password=" + getPostgres.password + " dbname=" + getPostgres.database + " port=" + getPostgres.port + " sslmode=disable TimeZone=Asia/Shanghai"
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	GDb = Db
	if err != nil {
		l.Error("加载数据库失败！！！！" + err.Error())
	}
	sqlDb, _ := Db.DB()
	//最大连接数量
	sqlDb.SetMaxIdleConns(100)
	//最大线程数
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetConnMaxLifetime(100 * time.Second)
	l.Info("数据库配置加载完毕")
	return
}
