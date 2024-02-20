package product

// User port interface definition for depedency injection
type UserRepository interface {
	Save(u *User) (int64, error)
	GetByEmail(email string) (*User, bool, error)
}
