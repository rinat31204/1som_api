package user

type Repository interface {
	Create(user *User) error
	FindById(id string) (*User, error)
	ExistsUserByIdOrEmail(id, email string) (bool, error)
}
