package domain

type User struct {
	Id    string `gorm:"primary_key"`
	Name  string
	Email string
}
type UserStore interface {
	GetAll() ([]User, error)
	CreateUser(User) error
	// Update(string, User) error
	// Delete(string) error
	// GetById(string) (User, error)

}
