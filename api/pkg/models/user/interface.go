package user

type Interface interface {
	Create(user *User) error
	Delete(userID uint) error
	Update(userID uint, user *User) error
	Get() ([]*User, error)
	Validate(name, pass string) (*User, error)
}
