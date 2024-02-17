package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/sebajax/go-vertical-slice-architecture/internal/user"
	"github.com/sebajax/go-vertical-slice-architecture/pkg/database"
)

// User repository for querying the database
type userRepository struct {
	db *database.DbConn
}

// Create a user instance repository
func NewUserRepository(dbcon *database.DbConn) user.UserRepository {
	return &userRepository{db: dbcon}
}

// Stores a new user in the database
func (repo *userRepository) Save(u *user.User) (int64, error) {
	// Get the id inserted in the database
	var id int64

	query := `INSERT INTO client (identity_number, first_name, last_name, email, date_of_birth) 
				VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := repo.db.DbPool.QueryRow(query, u.IdentityNumber, u.FirstName, u.LastName, u.Email, u.DateOfBirth).Scan(&id)
	if err != nil {
		return 0, err
	}

	fmt.Println("id: ", id)

	// No errors return the user id inserted
	return id, nil
}

// Gets the user by the email
func (repo *userRepository) GetByEmail(email string) (*user.User, bool, error) {
	u := user.User{}
	query := `SELECT id, identity_number, first_name, last_name, email, date_of_birth, created_at
				FROM client 
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
