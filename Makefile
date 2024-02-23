.PHONY: build-server start-server stop-server live-reload test test-coverage clean-deps format clean-imports lint vet check-shadow lint-format migrate-create migrate-up

# Define variable for migration directory and PostgreSQL URL
MIGRATION_DIR = /migration
POSTGRESQL_URL = your_postgresql_connection_string # Replace with your actual connection string

# Docker tasks
build-server:
	docker-compose -p go-vertical-slice-architecture build

start-server:
	docker-compose up -d

stop-server:
	docker-compose down

# Standalone usage for live reloading
live-reload:
	air

# Testing
test:
	go test ./...

test-coverage:
	go test -cover ./...

# Cleaning, Formatting, Linting, and Vetting
clean-deps:
	go mod tidy

format:
	go fmt ./...

clean-imports:
	goimports -l -w .

lint:
	golangci-lint run ./...

vet:
	go vet ./...

check-shadow:
	shadow ./...

lint-format:
	go fmt ./...
	go vet ./...
	golangci-lint run ./...

# Database Migration
migrate-create:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)

migrate-up:
	migrate -database $(POSTGRESQL_URL) -path $(MIGRATIONS_DIR) up

# Usage instructions:
# - To build the server: make build-server
# - To start the server: make start-server
# - To stop the server: make stop-server
# - For live reloading during development: make live-reload
# - To run tests: make test
# - To check test coverage: make test-coverage
# - To clean dependencies: make clean-deps
# - To format code: make format
# - To clean unused imports: make clean-imports
# - To lint code: make lint
# - To vet code: make vet
# - To check for shadowed variables: make check-shadow
# - To lint, format and vet your once: make lint-format
# - To create a migration script (replace your_script_name with the actual name): make migrate-create name=your_script_name
# - To run migration scripts: make migrate-up