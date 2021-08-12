package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"net"
	"time"

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

func (s *server) Prime(req *proto.PrimeRequest, stream proto.AppService_PrimeServer) error {
	pRange := req.GetPRange()
	start := pRange.GetStart()
	end := pRange.GetEnd()
	for i := start; i <= end; i++ {
		time.Sleep(500 * time.Millisecond)
		if isPrime(i) {
			res := &proto.PrimeResponse{
				No: i,
			}
			if err := stream.Send(res); err != nil {
				return err
			}
		}
	}
	return nil
}

func isPrime(no int64) bool {
	var i int64
	for i = 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func (s *server) Greet(stream proto.AppService_GreetServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		greeting := req.GetGreeting()
		first_name := greeting.GetFirstName()
		last_name := greeting.GetLastName()
		greetMsg := fmt.Sprintf("Hi %s %s, Have a nice day!", first_name, last_name)
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Sending : ", greetMsg)
		resp := &proto.GreetResponse{GreetMessage: greetMsg}
		stream.Send(resp)
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
