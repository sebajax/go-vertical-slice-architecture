package infrastructure

import (
	"database/sql"

	"github.com/sebajax/go-architecture-angrycoders/internal/user"
	"github.com/sebajax/go-architecture-angrycoders/pkg/database"
)

// User repository for querying the database
type userRepository struct {
	db *database.DbConn
}

// Create a user instance repository
func NewUserRepository(dbcon *database.DbConn) *userRepository {
	return &userRepository{db: dbcon}
}

// Stores a new user in the database
func (repo *userRepository) Save(u *user.User) (int64, error) {
	query := `INSERT INTO users (identity_number, first_name, last_name, email, date_of_birth) 
				VALUES ($1, $2, $3)`
	data, err := repo.db.DbPool.Exec(query, u.IdentityNumber, u.FirstName, u.LastName, u.Email, u.DateOfBirth)
	if err != nil {
		return 0, err
	}

	// Get the the id inserted in the database
	userId, _ := data.LastInsertId()

	// No errors return the user id inserted
	return userId, nil
}

// Gets the user by the email
func (repo *userRepository) GetByEmail(email string) (*user.User, bool, error) {
	u := user.User{}
	query := `SELECT id, identity_number, first_name, last_name, email, role, date_of_birth, created_at, password 
				FROM users 
				WHERE email = $1`
	err := repo.db.DbPool.QueryRow(query, email).Scan(&u.Id, &u.IdentityNumber, &u.FirstName, &u.LastName, &u.Email, &u.DateOfBirth, &u.CreatedAt)
	if err != nil {
		// Not found, but not an error
		if err == sql.ErrNoRows {
			return nil, false, nil
		}
		// An actual error occurred
		return nil, false, err
	}

	// Found the item
	return &u, true, nil
}
