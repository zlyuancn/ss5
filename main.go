package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/zlyuancn/go-socks5"
)

func main() {
	addr := flag.String("addr", ":1080", "监听地址")
	user := flag.String("user", "", "用户名")
	pwd := flag.String("pwd", "", "密码")
	flag.Parse()

	conf := &socks5.Config{}
	if *user != "" && *pwd != "" {
		conf.Credentials = socks5.StaticCredentials{*user: *pwd}
	}

	// 创建socks5服务
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	// 开始监听
	l, err := net.Listen("tcp", *addr)
	if err != nil {
		panic(fmt.Errorf("监听失败: %v", err))
	}

	// 打印启动信息
	var showUser string
	if *user != "" && *pwd != "" {
		showUser = fmt.Sprintf(", 用户: %s, 密码: %s", *user, *pwd)
	}
	fmt.Printf("开始监听: %s%s\n", *addr, showUser)

	// 启动
	if err = server.Serve(l); err != nil {
		panic(fmt.Errorf("启动失败: %v", err))
	}
	fmt.Println("停止")
}
