package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func JSONMethod(content interface{}) map[string]interface{} {
	var name map[string]interface{}
	if marshalContent, err := json.Marshal(content); err != nil {
		fmt.Println(err)
	} else {
		d := json.NewDecoder(bytes.NewReader(marshalContent))
		d.UseNumber() // 设置将float64转为一个number
		if err := d.Decode(&name); err != nil {
			fmt.Println(err)
		} else {
			for k, v := range name {
				name[k] = v
			}
		}
	}
	return name
}

func TimeStringToUnix(str string) int64 {
	// go语言固定日期模版
	timeLayout := "2006-01-02 15:04:05"
	location, _ := time.LoadLocation("Asia/Shanghai") // 指定东八区时区
	times, _ := time.ParseInLocation(timeLayout, str, location)
	return times.Unix()
}

// 时间戳转时间（字符串）
func UnixToTime(unix int64) string {
	timeLayout := "2006-01-02 15:04:05"
	location, _ := time.LoadLocation("Asia/Shanghai") // 指定东八区时区
	return time.Unix(unix, 0).In(location).Format(timeLayout)
}

func UnixToTime2(value interface{}) string {
	t, err := time.Parse(time.RFC3339, fmt.Sprintf("%v", value))
	if err != nil {
		return ""
	}

	someTime := t.Format("2006-01-02 15:04:05")
	someTime = strings.ReplaceAll(someTime, ":", " ")
	someTime = someTime[:len(someTime)-3]

	return someTime
}

func ChangeShowType(timeOri string) string {
	timeLayout := "2006-01-02 15:04:05"
	location, _ := time.LoadLocation("Asia/Shanghai")
	times, _ := time.ParseInLocation(timeLayout, timeOri, location)
	layout2 := "2006/1/2 15:04"
	return time.Unix(times.Unix(), 0).In(location).Format(layout2)
}
