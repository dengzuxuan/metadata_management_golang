package model

import (
	"fmt"
	"others-part/utils"
	"strconv"
)

type User struct {
	Id         int    `gorm:"id" json:"id"`
	Username   string `gorm:"username" json:"username"`
	Password   string `gorm:"password" json:"password"`
	Role       int    `gorm:"role" json:"role"`
	Avatar     string `gorm:"avatar" json:"avatar"`
	Status     string `gorm:"status" json:"status"`
	Department string `gorm:"department" json:"department"`
	Background string `gorm:"background" json:"background"`
	Telephone  string `gorm:"telephone" json:"telephone"`
	Email      string `gorm:"email" json:"email"`
	Address    string `gorm:"address" json:"address"`
	Place      string `gorm:"place" json:"place"`
	Statement  string `gorm:"statement" json:"statement"`
	Male       string `gorm:"male" json:"male"`
	RoleInfo   string `gorm:"-" json:"roleInfo"`
}
type UserShow struct {
	Id         int    `gorm:"id" json:"id"`
	Username   string `gorm:"username" json:"username"`
	Role       int    `gorm:"role" json:"role"`
	Avatar     string `gorm:"avatar" json:"avatar"`
	Status     string `gorm:"status" json:"status"`
	Department string `gorm:"department" json:"department"`
	Background string `gorm:"background" json:"background"`
	Telephone  string `gorm:"telephone" json:"telephone"`
	Email      string `gorm:"email" json:"email"`
	Address    string `gorm:"address" json:"address"`
	Place      string `gorm:"place" json:"place"`
	Statement  string `gorm:"statement" json:"statement"`
	Male       string `gorm:"male" json:"male"`
	RoleInfo   string `gorm:"-" json:"roleInfo"`
}
type UpdateUserInfo struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Role       int    `json:"role"`
	Avatar     string `json:"avatar"`
	Status     string `json:"status"`
	Department string `json:"department"`
	Background string `json:"background"`
	Telephone  string `json:"telephone"`
	Email      string `json:"email"`
	Address    string `json:"address"`
	Place      string `json:"place"`
	Statement  string `json:"statement"`
	Male       string `json:"male"`
	RoleInfo   string `json:"roleInfo"`
	Newbg      string `json:"newbg"`
	Newavatar  string `json:"newavatar"`
}

var roleMap = map[int]string{
	0: "超级管理员",
	1: "管理员",
	2: "部门管理员",
	3: "普通员工",
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type GuidModel struct {
	Guid string `json:"guid"`
}
type AtlasSearchPre struct {
	Query string `json:"query"`
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
func GetUserAvatar(id int) string {
	var user User
	db.Select("avatar").Where("id = ?", id).First(&user)
	return user.Avatar
}
func GetUserName(id int) string {
	var user User
	db.Where("id = ?", id).First(&user)
	return user.Username
}
func GetUserInfo(id int) (username string, avatar string) {
	var user User
	db.Where("id = ?", id).First(&user)
	return user.Username, user.Avatar
}
func GetUserInfos(id int) User {
	var user User
	db.Where("id = ?", id).First(&user)
	user.RoleInfo = roleMap[user.Role]
	return user
}
func GetUserId(name string) int {
	var user User
	db.Where("username = ?", name).First(&user)
	return user.Id
}

func GetUserRole(userid int) int {
	var user User
	db.Where("userid = ?", userid).First(&user)
	return user.Role
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
func UpdateUser(id int, user User) int {
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
