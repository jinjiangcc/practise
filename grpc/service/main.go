package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "practise/grpc/proto"
)

const PORT = "9001"

type HelloServer struct{

}

func (*HelloServer)Hello(ctx context.Context,arg *pb.HelloRequest)  (*pb.HelloResponse, error) {
	response := &pb.HelloResponse{
	}
	if arg.Request == "hello"{
		response.Response = "world"
	}
	return response, nil
}

func main(){
	server := grpc.NewServer()
	pb.RegisterHelloWorldServiceServer(server, &HelloServer{})
	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	err = server.Serve(lis)
	if err != nil{
		log.Fatalf("start server: %v", err)
	}
}

