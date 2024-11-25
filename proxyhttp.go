package main

import (
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

func main() {
	// Tạo một proxy HTTP mới
	proxy := goproxy.NewProxyHttpServer()

	// Cấu hình proxy để ghi lại nhật ký chi tiết của các yêu cầu
	proxy.Verbose = true

	// Khởi động máy chủ proxy trên cổng 1111
	log.Println("Máy chủ proxy HTTP đang chạy trên cổng 1111...")
	err := http.ListenAndServe("0.0.0.0:1111", proxy)
	if err != nil {
		log.Fatal("Lỗi khi chạy máy chủ proxy HTTP:", err)
	}
}
