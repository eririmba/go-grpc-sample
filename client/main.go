package main

import (
	"fmt"
	"log"

	pb "../proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:13009", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Dial:", err)
	}
	defer conn.Close()

	c := pb.NewUserCRUDClient(conn)
	user := &pb.User{
		Id: 15,
	}

	me, err := c.GetUser(context.Background(), user)
	if err != nil {
		log.Fatalln("GetUser:", err)
	}

	fmt.Println(me.Name + "さん、こんにちは")
	fmt.Println("現在のいいね数は " + fmt.Sprint(me.LikeCount) + " です")
	fmt.Println("登録日: " + me.CreatedAt)
}
