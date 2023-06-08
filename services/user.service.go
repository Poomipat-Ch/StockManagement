package services

import (
	"fmt"

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
func (u *userService) CreateUser(payload *dto.CreateUserRequest) (*dto.CreateUserResponse, []*models.ValidationError, error) {

	if err := u.v.Struct(payload); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil, nil, err
		}
		validateErrs := []*models.ValidationError{}
		for _, err := range err.(validator.ValidationErrors) {
			e := &models.ValidationError{
				Namespace:       err.Namespace(),
				Field:           err.Field(),
				StructNamespace: err.StructNamespace(),
				StructField:     err.StructField(),
				Tag:             err.Tag(),
				ActualTag:       err.ActualTag(),
				Kind:            fmt.Sprintf("%v", err.Kind()),
				Type:            fmt.Sprintf("%v", err.Type()),
				Value:           fmt.Sprintf("%v", err.Value()),
				Param:           err.Param(),
				Message:         err.Error(),
			}

			validateErrs = append(validateErrs, e)
		}

		return nil, validateErrs, nil
	}

	newUser, err := models.CreateNewUser(payload.FirstName, payload.LastName, payload.Email, payload.Password)

	if err != nil {
		return nil, nil, err
	}

	tx := u.db.MustBegin()
	tx.NamedExec("INSERT INTO users (first_name, last_name, email, password_hashed) VALUES (:first_name, :last_name, :email, :password_hashed)", newUser)
	tx.Commit()

	return &dto.CreateUserResponse{ID: 1}, nil, nil
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
func (u *userService) GetUsers() ([]*models.User, error) {
	users := []*models.User{}

	if err := u.db.Select(&users, "SELECT * FROM users"); err != nil {
		return nil, err
	}

	return users, nil
}

// UpdateUser implements UserService.
func (*userService) UpdateUser(user *models.User) error {
	panic("unimplemented")
}

func NewUserService(db *sqlx.DB, v *validator.Validate) UserService {
	return &userService{db: db, v: v}
}
