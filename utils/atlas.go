package utils

import (
	"fmt"
	"io"
	"net/http"
)

func AtlasApi() {
	// 构建请求
	resp, err := http.Get("Error: connect ECONNREFUSED 127.0.0.1:10809")
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println(string(body))
}
