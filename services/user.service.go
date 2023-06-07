package services

import (
	"github.com/Poomipat-Ch/StockManagement/dto"
	"github.com/Poomipat-Ch/StockManagement/models"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

type userService struct {
	db *sqlx.DB
	v  *validator.Validate
}

// CreateUser implements UserService.
func (u *userService) CreateUser(payload *dto.CreateUserRequest) (*dto.CreateUserResponse, error) {
	newUser, err := models.CreateNewUser(payload.FirstName, payload.LastName, payload.Email, payload.Password)

	if err != nil {
		return nil, err
	}

	tx := u.db.MustBegin()
	tx.NamedExec("INSERT INTO users (first_name, last_name, email, password_hashed) VALUES (:first_name, :last_name, :email, :password_hashed)", newUser)
	tx.Commit()

	return &dto.CreateUserResponse{ID: 1}, nil
}

// DeleteUser implements UserService.
func (*userService) DeleteUser(id int) error {
	panic("unimplemented")
}

// GetUserByEmail implements UserService.
func (*userService) GetUserByEmail(email string) (*models.User, error) {
	panic("unimplemented")
}

// GetUserByID implements UserService.
func (*userService) GetUserByID(id int) (*models.User, error) {
	panic("unimplemented")
}

// GetUsers implements UserService.
func (*userService) GetUsers() ([]*models.User, error) {
	panic("unimplemented")
}

// UpdateUser implements UserService.
func (*userService) UpdateUser(user *models.User) error {
	panic("unimplemented")
}

func NewUserService(db *sqlx.DB, v *validator.Validate) UserService {
	return &userService{db: db, v: v}
}
