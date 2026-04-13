package user

type Repository interface {
	Create(user *User) error
	FindById(id string) (*User, error)
	ExistsUser(id string) (bool, error)
}
