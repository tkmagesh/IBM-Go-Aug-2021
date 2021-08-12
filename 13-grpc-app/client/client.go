package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	client, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer client.Close()

	clientConn := proto.NewAppServiceClient(client)
	//doRequestResponse(clientConn)
	//doClientStreaming(clientConn)
	//doServerStreaming(clientConn)
	//doBiDirectionalStreaming(clientConn)
	doRequestResponseWithTimeout(clientConn)

}

func doRequestResponseWithTimeout(client proto.AppServiceClient) {
	req := &proto.AddRequest{X: 100, Y: 200}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	resp, err := client.Add(ctx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Fatalln("Timeout was hit! Deadline Exceeded")
			} else {
				log.Fatalln(err)
			}
		}
	} else {
		log.Println("Add Result:", resp.GetResult())
	}
}

func doBiDirectionalStreaming(client proto.AppServiceClient) {
	stream, err := client.Greet(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	requestData := []*proto.GreetRequest{
		&proto.GreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "Magesh",
				LastName:  "Kuppan",
			},
		},
		&proto.GreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "Suresh",
				LastName:  "Kannan",
			},
		},
		&proto.GreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "Ramesh",
				LastName:  "Jayaraman",
			},
		},
		&proto.GreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "Rajesh",
				LastName:  "Pandit",
			},
		},
		&proto.GreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "Naresh",
				LastName:  "Kumar",
			},
		},
	}

	go func() {
		for _, req := range requestData {
			stream.Send(req)
		}
		stream.CloseSend()
	}()
	/* wg := &sync.WaitGroup{}
	wg.Add(1) */
	done := make(chan bool)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("Thats all folks!")
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("Greet Result:", res.GetGreetMessage())
		}
		//done <- true
		close(done)
	}()
	<-done
}

func doServerStreaming(clientConn proto.AppServiceClient) {
	pRange := &proto.PrimeRange{
		Start: 3,
		End:   100,
	}
	primeReq := &proto.PrimeRequest{
		PRange: pRange,
	}
	stream, err := clientConn.Prime(context.Background(), primeReq)
	if err != nil {
		log.Fatalf("%v.Prime(_) = _, %v", clientConn, err)
	}
	for {
		pr, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.Recv() = _, %v", stream, err)
		}
		fmt.Println("Prime No : ", pr.GetNo())
	}
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
