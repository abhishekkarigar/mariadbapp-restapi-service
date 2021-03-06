package domain

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Id    int
	Name  string
	Email string
}
type UserStore interface {
	GetAll() ([]User, error)
	CreateUser(User) error
	DeleteUser(string) error
	GetUser(User, string) (User, error)
	// Update(string, User) error

	// GetById(string) (User, error)

}
