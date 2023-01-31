package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/go/protogrpc/testpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cc, err := grpc.Dial("localhost:5070", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	c := testpb.NewTestServiceClient(cc)

	// DoUnary(c)
	// DoClientStreaming(c)
	// DoServerStreaming(c)
	DoBidirectionalStreaming(c)
}

func DoUnary(c testpb.TestServiceClient) {
	req := &testpb.GetTestRequest{
		Id: "u1",
	}

	res, err := c.GetTest(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("response from server: ", res)
}

func DoClientStreaming(c testpb.TestServiceClient) {
	questions := []*testpb.Question{
		{
			Id:       "p1",
			Answer:   "hola",
			Question: "kjja",
			TestId:   "u1",
		},
		{
			Id:       "p2",
			Answer:   "hola abraham",
			Question: "que edad tengo",
			TestId:   "u1",
		},
	}

	stream, err := c.SetQuestion(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, question := range questions {
		log.Println("mesage send: ", question.Id)
		stream.Send(question)
		time.Sleep(2 * time.Second)
	}

	msg, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("server response: ", msg)
}

func DoServerStreaming(c testpb.TestServiceClient) {
	req := &testpb.GetStudentsPerTestRequest{
		TestId: "u1",
	}

	stream, err := c.GetStudentsPerTest(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Response server: ", msg)
	}

}

func DoBidirectionalStreaming(c testpb.TestServiceClient) {
	answer := testpb.TakeTestRequest{
		Answer: "42",
	}

	numberOfQuestions := 4

	waitChannel := make(chan struct{})

	stream, err := c.TakeTest(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for i:=0; i <numberOfQuestions; i++ {
			stream.Send(&answer)
			time.Sleep(1 * time.Second)
		}	
	}()


	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
				break
			}
			log.Println("response server: ", res)
		}
		close(waitChannel)
	}()

	<- waitChannel
}
