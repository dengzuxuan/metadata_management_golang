package model

type Message struct {
	UserId   int    `json:"user_id"`
	Content  string `json:"content"`   //具体内容
	TypeInfo string `json:"type_info"` //类型 1.收藏 2.点赞 3.关注 4.回复 5.评论 6.系统通知
	UserName string `json:"user_name"` //用户名称
	Time     int64  `json:"time"`      //时间戳
	Date     string `json:"date"`      //日期 2023-05-25
	Date2    string `json:"date_2"`    //时间 12:42
	Type     string `json:"type"`
}

type Messages []Message

func (s Messages) Len() int {
	return len(s)
}
func (s Messages) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Messages) Less(i, j int) bool {
	return s[i].Time < s[j].Time
}
