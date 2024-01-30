package database

type DbConfig struct {
	dbUser          string
	dbPassword      string
	dbName          string
	dbHost          string
	dbPort          string
	dbMigrationPath *string // Optional attribute if not set will be nil
}

// Create a new db config type
func NewDbConfig(user string, password, string, name string, host string, port string) *DbConfig {
	return &DbConfig{
		dbUser:     user,
		dbPassword: password,
		dbName:     name,
		dbHost:     host,
		dbPort:     port,
	}
}

// If dbMigration is needed in the api we use this function to set the path
func (config *DbConfig) WithMigration(path string) *DbConfig {
	config.dbMigrationPath = &path
	return config
}