package main

import (
	"context"
	"log"

	"github.com/armon/go-socks5"
)

// CustomRuleSet để log thông tin về các kết nối
type CustomRuleSet struct{}

// Hàm Allow sẽ log chi tiết về các yêu cầu tới proxy
func (r CustomRuleSet) Allow(ctx context.Context, req *socks5.Request) (context.Context, bool) {
	// Log thông tin về yêu cầu SOCKS5
	log.Printf("New connection from %s to %s, command: %d", req.RemoteAddr.String(), req.DestAddr.String(), req.Command)

	return ctx, true
}

func main() {
	// Tạo cấu hình SOCKS5 server với custom RuleSet
	conf := &socks5.Config{
		Rules: CustomRuleSet{},
	}

	// Khởi tạo server SOCKS5
	server, err := socks5.New(conf)
	if err != nil {
		log.Fatal("Lỗi khi tạo SOCKS5 server:", err)
	}

	// Lắng nghe và phục vụ trên cổng 1080
	log.Println("Máy chủ proxy SOCKS5 đang chạy trên cổng 1080...")
	err = server.ListenAndServe("tcp", "0.0.0.0:1080")
	if err != nil {
		log.Fatal("Lỗi khi chạy SOCKS5 server:", err)
	}
}
