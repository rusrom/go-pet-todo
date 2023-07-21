db:
	docker run --rm --name todo-db -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres

wait-db-up:
	sleep 5

migrate-up:
	migrate -path ./migration -database "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable" up

migrate-down:
	migrate -path ./migration -database "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable" down

run: db wait-db-up migrate-up
	go run cmd/main.go
