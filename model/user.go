package model

import (
	"fmt"
	"others-part/utils"
	"strconv"
)

type User struct {
	Id         int    `gorm:"id"`
	Username   string `gorm:"username"`
	Password   string `gorm:"password"`
	Role       int    `gorm:"role"`
	Avatar     string `gorm:"avatar"`
	Status     string `gorm:"status"`
	Department string `gorm:"department"`
}

func (this *User) TableName() string {
	return "User"
}
func AddUser(username string, password string, email string) int {
	newUser := User{
		Username: username,
		Password: password,
	}
	err = db.Create(&newUser).Error
	if err != nil {
		return utils.ERROR_CREAT_WRONG
	}
	return utils.SUCCESS
}

//	func CheckAdd(username string, email string) (code int) {
//		var user User
//		db.Debug().Select("ID").Where("email = ?", email).First(&user)
//		if user.ID != 0 {
//			return utils.ERROR_TEL_USED
//		}
//		db.Debug().Select("ID").Where("username = ?", username).First(&user)
//		if user.ID != 0 {
//			return utils.ERROR_USERNAME_USED
//		}
//		return utils.SUCCESS
//	}
//
//	func CheckLogin(username string, password string) (code int, ID string) {
//		var user User
//		db.Select("ID").Where("username = ?", username).First(&user)
//		if user.ID == 0 {
//			return utils.ERROR_USERNAME_NOT_EXIST, ""
//		}
//		db.Select("password").Where("username = ?", username).First(&user)
//		if user.Password == password {
//			return utils.SUCCESS, strconv.Itoa(int(user.ID))
//		} else {
//			return utils.ERROR_PASSWORD_WRONG, ""
//		}
//	}
//
//	func CheckLoginEmail(email string) int {
//		var user User
//		db.Select("ID").Where("email = ?", email).First(&user)
//		if user.ID == 0 {
//			return utils.ERROR_EMAIL_NOT_EXIST
//		}
//		return utils.SUCCESS
//	}
//
//	func FindLoginId(email string) string {
//		var user User
//		db.Select("ID").Where("email = ?", email).First(&user)
//		return strconv.Itoa(int(user.ID))
//	}
//
//	func AddMusicKind(id string, musicKind string) int {
//		err = db.Model(&User{}).Where("id = ?", id).Update("liketypes", musicKind).Error
//		if err != nil {
//			return utils.ERROR
//		}
//		return utils.SUCCESS
//	}
func FindUser(id string) (int, User) {
	var user User
	err = db.Debug().Where("ID = ?", id).First(&user).Error
	if err != nil {
		return utils.ERROR, user
	}
	return utils.SUCCESS, user
}
func UpdateUser(id string, user User) int {
	err = db.Where("id = ?", id).Updates(user).Error
	if err != nil {
		return utils.ERROR_CHANGE_WRONG
	}
	return utils.SUCCESS
}
func CheckLogin(username string, password string) (code int, ID string) {
	var user User
	db.Select("id").Where("username = ?", username).First(&user)
	if user.Id == 0 {
		return utils.ERROR_USERNAME_NOT_EXIST, ""
	}
	db.Select("password").Where("username = ?", username).First(&user)
	if user.Password == password {
		db.Where("username = ?", username).First(&user)
		fmt.Println(user.Id, user.Username, user.Password, user.Role)
		return utils.SUCCESS, strconv.Itoa(int(user.Id))
	} else {
		return utils.ERROR_PASSWORD_WRONG, ""
	}

}
