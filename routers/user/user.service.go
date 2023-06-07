package user

import "github.com/jmoiron/sqlx"

type UserService interface {
	GetUsers() ([]*User, error)
	GetUserByID(id int) (*User, error)
	GetUserByEmail(email string) (*User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(id int) error
}

type userService struct {
	db *sqlx.DB
}

func NewUserService(db *sqlx.DB) UserService {
	return &userService{db: db}
}

// GetUsers implements UserService.
func (u *userService) GetUsers() ([]*User, error) {
	users := []*User{}

	err := u.db.Select(&users, "SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	return users, nil
}

// CreateUser implements UserService.
func (u *userService) CreateUser(user *User) error {
	tx := u.db.MustBegin()
	tx.NamedExec("INSERT INTO users (first_name, last_name, email, password_hashed) VALUES ($1, $2, $3, $4)", user)
	tx.Commit()
	return nil
}

// DeleteUser implements UserService.
func (u *userService) DeleteUser(id int) error {
	panic("unimplemented")
}

// GetUserByEmail implements UserService.
func (u *userService) GetUserByEmail(email string) (*User, error) {
	panic("unimplemented")
}

// GetUserByID implements UserService.
func (u *userService) GetUserByID(id int) (*User, error) {
	panic("unimplemented")
}

// UpdateUser implements UserService.
func (u *userService) UpdateUser(user *User) error {
	panic("unimplemented")
}
