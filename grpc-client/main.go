package main

import (
	"context"
	pb "github.com/go-tour/grpc-client/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, _ := grpc.Dial(":9000", grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewCheeseClient(conn)
	_ = MyTest(client)
}

func MyTest(client pb.CheeseClient) error {
	// 测试登录接口
	func() {
		resp, _ := client.Login(context.Background(), &pb.LoginRequest{LoginName: "eddycjy", PassWord: "123"})
		log.Printf("client.SayHello resp: %v", resp)
	}()

	// 测试注册接口
	func() {
		resp, _ := client.Register(context.Background(), &pb.RegisterRequest{LoginName: "eddycjy", PassWord: "123"})
		log.Printf("client.SayHello resp: %v", resp)
	}()

	// 测试下棋接口
	func() {
		resp, _ := client.Play(context.Background(), &pb.PlayRequest{User: 1, Row: 1, Column: 1})
		log.Printf("client.SayHello resp: %v", resp)
	}()

	// 测试投降接口
	func() {
		resp, _ := client.Surrender(context.Background(), &pb.SurrenderRequest{User: 1})
		log.Printf("client.SayHello resp: %v", resp)
	}()

	// 测试获取棋盘接口
	func() {
		resp, _ := client.GetChessBoard(context.Background(), &pb.ChessBoardRequest{})
		log.Printf("client.SayHello resp: %v", resp)
	}()

	return nil
}
