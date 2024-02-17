package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type DbConn struct {
	DbPool *sql.DB
}

// InitPool initializes the database connection pool and runs migrations.
func InitPool(config *DbConfig) *DbConn {
	var db *sql.DB

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.dbHost, config.dbPort, config.dbUser, config.dbPassword, config.dbName)

	// Initialize the connection pool
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		log.Panicf("error opening database: %v", err)
	}

	// Verify the connection
	if err := db.Ping(); err != nil {
		log.Panicf("error pinging database: %v", err)
	}

	// Run migrations if needed only if migration path is set
	if config.dbMigrationPath != nil {
		if err := runMigration(db, *config.dbMigrationPath); err != nil {
			log.Panicf("error pinging database: %v", err)
		}
	}

	// Return the connection pool pointer singleton connection
	return &DbConn{
		DbPool: db,
	}
}

// ClosePool closes the database connection pool.
func (db *DbConn) ClosePool() error {
	if db.DbPool != nil {
		return db.DbPool.Close()
	}
	return nil
}

// Run database migration files only if migration path is set
func runMigration(db *sql.DB, path string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", path),
		"postgres",
		driver,
	)
	if err != nil {
		return err
	}

	// Migrate up to the latest active version
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
