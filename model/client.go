package model

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"others-part/utils"
)

var db *gorm.DB
var err error

func InitMysqlDb() {
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	fmt.Println(args)
	db, err = gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		logrus.Error("failed open mysql", err)
		return
	}
	logrus.Info("open mysql success")
	err = db.AutoMigrate(&User{})
	if err != nil {
		logrus.Error("failed open mysql", err)
		return
	}
}
