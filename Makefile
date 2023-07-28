db:
	docker run --rm --name todo-db -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres;
	sleep 4

migrate-up:
	migrate -path ./migration -database "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable" up

migrate-down:
	migrate -path ./migration -database "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable" down

start-api:
	go run cmd/main.go

run: db migrate-up start-api

user-1:
	curl --request POST \
      --url http://localhost:8080/auth/sign-up \
      --header 'Content-Type: application/json' \
      --data '{"name": "ruslan", "username": "rusrom", "password": "1qaz2wsx"}'

user-2:
	curl --request POST \
      --url http://localhost:8080/auth/sign-up \
      --header 'Content-Type: application/json' \
      --data '{"name": "oksana", "username": "oksankag", "password": "123456789"}'

new-users: user-1 user-2
