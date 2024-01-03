package main

import (
	"fmt"
	"google.golang.org/grpc"
	"im/services/user/api/v1"
	"im/services/user/config"
	"im/services/user/interfaces"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := config.Init(); err != nil {
		panic(err)
	}

	// 启动 gRPC 服务器
	lis, err := net.Listen("tcp", fmt.Sprintf("%s", config.C.GRPC.Addr))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	handler := interfaces.NewGrpcHandler(&config.C)
	api.RegisterUserServiceServer(grpcServer, handler)

	fmt.Printf("gRPC server is running on addr: %s\n", config.C.GRPC.Addr)

	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			grpcServer.Stop()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
