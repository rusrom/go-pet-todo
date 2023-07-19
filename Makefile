db:
	docker run --rm --name todo-db -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres

migrate-up:
	migrate -path ./migration -database "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable" up

migrate-down:
	migrate -path ./migration -database "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable" down
