package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedAppServiceServer
}

func (s *server) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	result := x + y
	res := &proto.AddResponse{
		Result: result,
	}
	return res, nil
}

func (s *server) Average(stream proto.AppService_AverageServer) error {
	var count int64
	var sum int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			avg := sum / count
			res := &proto.AverageResponse{
				Result: avg,
			}
			stream.SendAndClose(res)
			return nil
		}
		if err != nil {
			log.Fatal(err)
		}
		no := req.GetNo()
		fmt.Println("Received ", no)
		sum += no
		count++
	}
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, &server{})
	e := grpcServer.Serve(listener)
	if e != nil {
		log.Fatalln(e)
	}
}
