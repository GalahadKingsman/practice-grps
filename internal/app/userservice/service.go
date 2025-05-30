package userservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/GalahadKingsman/messenger_users/internal/database"
	"github.com/GalahadKingsman/messenger_users/internal/models"
	"strings"

	pb "github.com/GalahadKingsman/messenger_users/pkg/messenger_users_api" // Импорт сгенерированного пакета
)

type Service struct {
	pb.UnimplementedUserServiceServer
}

func (s *Service) CreateUser(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	// Валидация входных данных
	if req.GetFirstName() == "" || req.GetEmail() == "" {
		return nil, errors.New("имя и email обязательны для заполнения")
	}

	// Создание пользователя
	user := &models.User{
		Login:     req.GetLogin(),
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Email:     req.GetEmail(),
		Phone:     req.GetPhone(),
	}

	// Сохранение пользователя в базе данных
	id, err := database.CreateUser(*user)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании пользователя: %v", err)
	}

	// Формирование ответа
	return &pb.CreateResponse{
		Success: fmt.Sprintf("Пользователь успешно создан с ID: %s", id),
	}, nil
}
func (s *Service) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var (
		query = "SELECT id, first_name, last_name, email, phone FROM users"
		args  []interface{}
		where []string
	)

	if req.Id != nil {
		where = append(where, "id = ?")
		args = append(args, *req.Id)
	}
	if req.Login != nil {
		where = append(where, "login = ?")
		args = append(args, *req.Login)
	}
	if req.FirstName != nil {
		where = append(where, "first_name = ?")
		args = append(args, *req.FirstName)
	}
	if req.LastName != nil {
		where = append(where, "last_name = ?")
		args = append(args, *req.LastName)
	}
	if req.Email != nil {
		where = append(where, "email = ?")
		args = append(args, *req.Email)
	}
	if req.Phone != nil {
		where = append(where, "phone = ?")
		args = append(args, *req.Phone)
	}

	// Добавляем WHERE, если есть условия
	if len(where) > 0 {
		query += " WHERE " + strings.Join(where, " AND ")
	}

	row := database.DB.QueryRowContext(ctx, query, args...)

	var user models.User
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Phone)
	if err != nil {
		return nil, fmt.Errorf("пользователь не найден: %v", err)
	}
	return &pb.GetUserResponse{Id: user.ID,
		Login:     user.Login,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone}, nil
}

func New() *Service {
	return &Service{}
}
