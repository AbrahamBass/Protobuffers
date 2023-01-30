package server

import (
	"context"

	"github.com/go/protogrpc/models"
	"github.com/go/protogrpc/repository"
	"github.com/go/protogrpc/studentpb"
)

type server struct {
	repo repository.Repository
	studentpb.UnimplementedStudentServiceServer
}

func NewStudentServer(repo repository.Repository) *server {
	return &server{repo: repo}
}

func (s *server) GetStudent(ctx context.Context, req *studentpb.GetStudentRequest) (*studentpb.Student, error) {
	student, err := s.repo.GetStudent(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &studentpb.Student{
			Id: student.Id,
			Name: student.Name,
			Age: student.Age,
	}, nil
}

func (s *server) SetStudent(ctx context.Context, req *studentpb.Student) (*studentpb.SetStudentResponse, error) {
	student := &models.Student{
		Id: req.Id,
		Name: req.Name,
		Age: req.Age,
	}

	err := s.repo.SetStudent(ctx, student)
	if err != nil {
		return nil, err
	}
	return &studentpb.SetStudentResponse{
		Id: student.Id,
	}, nil
}