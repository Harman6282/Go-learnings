migrate-up:
	migrate -path db/migrations -database "postgresql://admin:adminpassword@localhost:5432/chat-db?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migrations -database "postgresql://admin:adminpassword@localhost:5432/chat-db?sslmode=disable" -verbose down

.PHONY: migrate-up migrate-down