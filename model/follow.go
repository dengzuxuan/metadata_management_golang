package model

import "time"

type UserFollowRecord struct {
	Id         int    `gorm:"id" json:"id"`
	Userid     int    `gorm:"userid" json:"userid"`
	Touserid   int    `gorm:"touserid" json:"touserid"`
	Createtime string `gorm:"createtime" json:"createtime"`
}

func (this *UserFollowRecord) TableName() string {
	return "UserFollowRecord"
}
func AddFollowRecord(userid, touserid int) {
	newFollowRecord := UserFollowRecord{
		Userid:     userid,
		Touserid:   touserid,
		Createtime: time.Now().Format("2006-01-02 15:04:05"),
	}
	_ = db.Debug().Create(&newFollowRecord)
}

func DelFollowRecord(userid, touserid int) {
	info := UserFollowRecord{}
	db.Where("userid=?", userid).Where("touserid=?", touserid).Delete(&info)
}

func CheckFollow(userid, touserid int) bool {
	userfollow := UserFollowRecord{}
	_ = db.Where("userid=?", userid).Where("touserid=?", touserid).Find(&userfollow)
	if userfollow.Id == 0 {
		return false
	}
	return true
}

type UserInfoBrief struct {
	Userid   int    `json:"userid"`
	Username string ` json:"username"`
	RoleInfo string `json:"role_info"`
	Avatar   string `json:"avatar"`
}

func GetAllFollowInfos(userid int) ([]UserInfoBrief, []UserInfoBrief) {
	userfollow := []UserFollowRecord{}
	userfollowInfo := []UserInfoBrief{}
	_ = db.Where("userid=?", userid).Find(&userfollow)
	for _, record := range userfollow {
		userInfo := GetUserInfos(record.Touserid)
		userfollowInfo = append(userfollowInfo, UserInfoBrief{
			Userid:   userInfo.Id,
			Username: userInfo.Username,
			RoleInfo: userInfo.RoleInfo,
			Avatar:   userInfo.Avatar,
		})
	}

	userfollowed := []UserFollowRecord{}
	userfollowedInfo := []UserInfoBrief{}
	_ = db.Where("touserid=?", userid).Find(&userfollowed)
	for _, record := range userfollowed {
		userInfo := GetUserInfos(record.Userid)
		userfollowedInfo = append(userfollowedInfo, UserInfoBrief{
			Userid:   userInfo.Id,
			Username: userInfo.Username,
			RoleInfo: userInfo.RoleInfo,
			Avatar:   userInfo.Avatar,
		})
	}
	return userfollowInfo, userfollowedInfo
}

func GetAllFollowMy(userid int) []UserInfoBrief {
	userfollowed := []UserFollowRecord{}
	userfollowedInfo := []UserInfoBrief{}
	_ = db.Where("touserid=?", userid).Find(&userfollowed)
	for _, record := range userfollowed {
		userInfo := GetUserInfos(record.Id)
		userfollowedInfo = append(userfollowedInfo, UserInfoBrief{
			Username: userInfo.Username,
			RoleInfo: userInfo.RoleInfo,
			Avatar:   userInfo.Avatar,
		})
	}
	return userfollowedInfo
}
