package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	client, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer client.Close()

	clientConn := proto.NewAppServiceClient(client)
	//doRequestResponse(clientConn)
	doClientStreaming(clientConn)

}

//client streaming
func doClientStreaming(clientConn proto.AppServiceClient) {
	clientStream, err := clientConn.Average(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	nos := []int64{3, 1, 4, 2, 5}
	for _, no := range nos {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Sending ", no)
		req := &proto.AverageRequest{
			No: no,
		}
		clientStream.Send(req)
	}
	res, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Average = ", res.GetResult())
}

//request & response
func doRequestResponse(clientConn proto.AppServiceClient) {
	req := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	res, _ := clientConn.Add(context.Background(), req)
	fmt.Println(res.GetResult())
}
