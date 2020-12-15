package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "practise/grpc/proto"
)

const PORT = "9001"

func main() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := pb.NewHelloWorldServiceClient(conn)
	request := &pb.HelloRequest{
		Request: "hello",
	}
	response, err := client.Hello(context.Background(), request)
	if err != nil {
		log.Fatalf("client  err: %v", err)
	}

	log.Printf("request: %v, response: %v \n" , request.GetRequest(), response.GetResponse())
}