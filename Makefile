# Declare targets as .PHONY
.PHONY: migrate_up run_server start

migrate_up:
	@echo "Running database migrations..."
	migrate -path ./migrate -path=internal/migrations -database "postgres://sunilkumar:hNzIMEnyspIwS1z62dea@drape-db-server.cpwmw0ss0w72.ap-south-1.rds.amazonaws.com:5432/postgres?sslmode=require&sslrootcert=internal/keys/ap-south-1-bundle.pem" up

migrate_down:
	@echo "Running database migrations..."
	migrate -path ./migrate -path=internal/migrations -database "postgres://sunilkumar:hNzIMEnyspIwS1z62dea@drape-db-server.cpwmw0ss0w72.ap-south-1.rds.amazonaws.com:5432/postgres?sslmode=require&sslrootcert=internal/keys/ap-south-1-bundle.pem" down

run_server:
	@echo "Starting Go server..."
	go run cmd/server/main.go

start: migrate_up run_server
	@echo "Migrations complete, starting backend server."
