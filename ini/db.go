package ini

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"ranbao/global"
	"ranbao/service/model"
)

func InitDB() {
	type Config struct {
		DSN string `yaml:"dsn"`
	}
	c := Config{
		DSN: fmt.Sprintf("root:1520226681wW@tcp(127.0.0.1:%s)/ranbao?charset=utf8mb4&parseTime=True&loc=Local", Conf.MysqlConfig.Port),
	}

	db, err := gorm.Open(mysql.Open(c.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	global.Db = db
	err = db.AutoMigrate(model.User{})
	if err != nil {
		panic(err)
	}
}
