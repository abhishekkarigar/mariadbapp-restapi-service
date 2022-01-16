package domain

type User struct {
	Id    string `gorm:"primary_key"`
	Name  string
<<<<<<< HEAD
	Email string
=======
	Email string `gorm:"unique;not null"`
>>>>>>> 11ef93d (first commit)
}
type UserStore interface {
	GetAll() ([]User, error)
	CreateUser(User) error
<<<<<<< HEAD
	// Update(string, User) error
	// Delete(string) error
=======
	DeleteUser(string) error
	GetUser(User, string) (User, error)
	// Update(string, User) error

>>>>>>> 11ef93d (first commit)
	// GetById(string) (User, error)

}
