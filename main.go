package main

import (
	"context"
	"log"

	"github.com/armon/go-socks5"
)

type CustomRuleSet struct{}

func (r CustomRuleSet) Allow(ctx context.Context, req *socks5.Request) (context.Context, bool) {
	log.Printf("New connection from %s to %s, command: %d", req.RemoteAddr.String(), req.DestAddr.String(), req.Command)
	return ctx, true
}

func main() {
	conf := &socks5.Config{
		Rules: CustomRuleSet{},
	}
	server, err := socks5.New(conf)
	if err != nil {
		log.Fatal("Lỗi khi tạo SOCKS5 server:", err)
	}

	log.Println("Máy chủ proxy SOCKS5 đang chạy trên cổng 1111... giao thực tcp")
	err = server.ListenAndServe("tcp", "0.0.0.0:1111")
	if err != nil {
		log.Fatal("Lỗi khi chạy SOCKS5 server:", err)
	}
}
