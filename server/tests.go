package server

import (
	"context"
	"io"

	"github.com/go/protogrpc/models"
	"github.com/go/protogrpc/repository"
	"github.com/go/protogrpc/testpb"
)

type TestServer struct {
	repo repository.Repository
	testpb.UnimplementedTestServiceServer
}

func NewTestServer(repo repository.Repository) *TestServer {
	return &TestServer{repo: repo}
}

func (this *TestServer) GetTest(ctx context.Context, req *testpb.GetTestRequest) (*testpb.Test, error) {
	test, err := this.repo.GetTest(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &testpb.Test{
		Id:   test.Id,
		Name: test.Name,
	}, nil
}

func (this *TestServer) SetTest(ctx context.Context, req *testpb.Test) (*testpb.SetTestResponse, error) {
	test := &models.Test{
		Id:   req.Id,
		Name: req.Name,
	}
	err := this.repo.SetTest(ctx, test)

	if err != nil {
		return nil, err
	}

	return &testpb.SetTestResponse{
		Id:   test.Id,
		Name: test.Name,
	}, nil
}

func (this *TestServer) SetQuestion(stream testpb.TestService_SetQuestionServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: true,
			})
		}

		if err != nil {
			return err
		}

		question := &models.Question{
			Id: msg.Id,
			Answer: msg.Answer,
			Question: msg.Question,
			TestId: msg.TestId,
		}

		err = this.repo.SetQuestion(context.Background(), question)
		if err != nil {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: false,
			})
		}
	}
}
