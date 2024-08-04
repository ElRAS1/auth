package main

import (
	"log"
	"net"

	userApi "github.com/ELRAS1/auth/pkg/userApi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	userApi.UnimplementedUserApiServer
}

func main() {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()

	reflection.Register(s)
	userApi.RegisterUserApiServer(s, server{})
	log.Printf("server started in port %d", 50051)
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
