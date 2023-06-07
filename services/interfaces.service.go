package services

import (
	"github.com/Poomipat-Ch/StockManagement/dto"
	"github.com/Poomipat-Ch/StockManagement/models"
)

type UserService interface {
	GetUsers() ([]*models.User, error)
	GetUserByID(id int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *dto.CreateUserRequest) (*dto.CreateUserResponse, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
}
