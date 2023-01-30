package main

import (
	"log"
	"net"

	"github.com/go/protogrpc/database"
	"github.com/go/protogrpc/server"
	"github.com/go/protogrpc/testpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	list, err := net.Listen("tcp", ":5070")

	if err != nil {
		log.Fatal(err)
	}

	repo, err := database.NewPostgresRepository("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")

	serve := server.NewTestServer(repo)

	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	testpb.RegisterTestServiceServer(s, serve)

	reflection.Register(s)

	if err := s.Serve(list); err != nil {
		log.Fatal(err)
	}
}
