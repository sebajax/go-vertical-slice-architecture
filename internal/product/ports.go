package user

// User port interface definition for depedency injection
type IUserRepository interface {
	Save(u *User) (int64, error)
	GetByEmail(email string) (User, bool, error)
}
