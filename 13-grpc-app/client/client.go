package main

import (
	"context"
	"fmt"
	"grpc-app/proto"

	"google.golang.org/grpc"
)

func main() {
	client, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer client.Close()

	clientConn := proto.NewAppServiceClient(client)
	req := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	res, _ := clientConn.Add(context.Background(), req)
	fmt.Println(res.GetResult())
}
