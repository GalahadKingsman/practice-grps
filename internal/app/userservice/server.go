package userservice

import (
	pb "api"
	"context"
	"messenger_user/internal/database"
	"messenger_user/internal/models"
)

type Service struct{}

func New() *Service { return &Service{} }

func (s *Service) Create(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
	}
	_, err := database.CreateUser(user)
	if err != nil {
		return &pb.RegisterResponse{Success: false, Message: "Ошибка при регистрации"}, err
	}
	return &pb.RegisterResponse{Success: true, Message: "Успешная регистрация"}, nil
}
