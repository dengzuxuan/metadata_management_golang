package main

import (
	"others-part/model"
	"others-part/router"
)

func main() {
	model.InitMysqlDb()
	r := router.InitRouter()
	r.Run(":8890")
}
