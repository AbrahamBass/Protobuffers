package main

import (
	"log"
	"net"

	"github.com/go/protogrpc/database"
	"github.com/go/protogrpc/server"
	"github.com/go/protogrpc/studentpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	list, err := net.Listen("tcp", ":5060")

	if err != nil {
		log.Fatal(err)
	}

	repo, err := database.NewPostgresRepository("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")

	serve := server.NewStudentServer(repo)

	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	studentpb.RegisterStudentServiceServer(s, serve)

	reflection.Register(s)

	if err := s.Serve(list); err != nil {
		log.Fatal(err)
	}
}
