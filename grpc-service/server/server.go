package server

import (
	"context"
	pb "github.com/go-tour/grpc-service/proto"
)

type RpcServer struct {
}

func (s *RpcServer) Login(ctx context.Context, r *pb.LoginRequest) (*pb.LoginReply, error) {
	return &pb.LoginReply{Code: 1, Message: "login return"}, nil
}

func (s *RpcServer) Register(ctx context.Context, r *pb.RegisterRequest) (*pb.RegisterReply, error) {
	return &pb.RegisterReply{Code: 1, Message: "register return"}, nil
}

func (s *RpcServer) Play(ctx context.Context, r *pb.PlayRequest) (*pb.PlayReply, error) {
	return &pb.PlayReply{ErrorCode: 0, GameAns: 1}, nil
}

func (s *RpcServer) Surrender(ctx context.Context, r *pb.SurrenderRequest) (*pb.SurrenderReply, error) {
	return &pb.SurrenderReply{ErrorCode: 0, GameAns: 2}, nil
}

func (s *RpcServer) GetChessBoard(context.Context, *pb.ChessBoardRequest) (*pb.ChessBoardReplay, error) {
	var points = []*pb.Point{
		{Row: 1, Column: 1},
		{Row: 2, Column: 2},
		{Row: 3, Column: 3},
	}

	return &pb.ChessBoardReplay{Points: points}, nil
}
