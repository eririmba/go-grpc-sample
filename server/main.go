package main

import (
	"fmt"
	"log"
	"net"

	pb "../proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Service struct{}

func (*Service) GetUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	fmt.Println("ユーザID: ", user.Id)
	fmt.Println("を送信しました")

	return &pb.User{
		Name:      "えりりんば",
		LikeCount: 240,
		CreatedAt: "2017-12-27",
	}, nil
}

func main() {
	s := grpc.NewServer()
	pb.RegisterUserCRUDServer(s, &Service{})

	l, err := net.Listen("tcp", ":13009")
	if err != nil {
		log.Fatalln(err)
	}

	s.Serve(l)
}
