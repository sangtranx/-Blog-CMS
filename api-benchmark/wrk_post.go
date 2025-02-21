package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Tạo nội dung file post.lua
	luaScript := `wrk.method = "POST"
wrk.headers["Content-Type"] = "application/json"
wrk.body = '{"email": "123@gmail.com", "password": "123456"}'
`

	// Ghi nội dung vào file post.lua
	fileName := "post.lua"
	err := os.WriteFile(fileName, []byte(luaScript), 0644)
	if err != nil {
		fmt.Println("Lỗi khi tạo file:", err)
		return
	}

	fmt.Println("Đã tạo file post.lua thành công.")

	// Chạy lệnh wrk
	cmd := exec.Command("wrk", "-t4", "-c1000", "-d30s", "-s", fileName, "http://localhost:8080/blog/login")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("Lỗi khi chạy wrk:", err)
		return
	}
}
