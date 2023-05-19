package utils

import "math/rand"

func RandString(n int) (ret string) {
	allString := "qwertyuiopasdfghjklzxcvbnm0123456789"
	ret = ""
	for i := 0; i < n; i++ {
		r := rand.Intn(len(allString))
		ret = ret + allString[r:r+1]
	}
	return
}
