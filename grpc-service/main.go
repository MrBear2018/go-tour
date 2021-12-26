package main

import (
	pb "github.com/go-tour/grpc-service/proto"
	rpcServer "github.com/go-tour/grpc-service/server"
	"google.golang.org/grpc"
	"net"
)

var port string

func init() {}

func main() {
	server := grpc.NewServer()
	pb.RegisterCheeseServer(server, &rpcServer.RpcServer{})
	lis, _ := net.Listen("tcp", ":9000")
	server.Serve(lis)
}
