package utils

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

var (
	Db         string
	DbPort     string
	DbHost     string
	DbPassWord string
	DbName     string
	DbUser     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		logrus.Error("配置文件错误", err)
		return
	}
	loadMysql(file)
}
func loadMysql(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbHost = file.Section("database").Key("DbHost").MustString("hadoop102")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("123456")
	DbName = file.Section("database").Key("DbName").MustString("root")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
}
