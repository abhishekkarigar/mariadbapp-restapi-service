package domain

type User struct {
	Id    string `gorm:"primary_key"`
	Name  string
	Email string `gorm:"unique;not null"`
}
type UserStore interface {
	GetAll() ([]User, error)
	CreateUser(User) error
	DeleteUser(string) error
	GetUser(User, string) (User, error)
	// Update(string, User) error

	// GetById(string) (User, error)

}
