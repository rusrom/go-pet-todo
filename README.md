## Awesome Go  

https://github.com/avelino/awesome-go


## PostgreSQL Docker container  
The PostgreSQL object-relational database system provides reliability and data integrity.  
https://hub.docker.com/_/postgres  
```shell
docker run --rm --name todo-db -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres
```

## migrate  
Database migrations. CLI and Golang library  
https://github.com/golang-migrate/migrate  
https://github.com/golang-migrate/migrate#cli-usage  
https://github.com/golang-migrate/migrate/blob/master/GETTING_STARTED.md  

migrate CLI    
https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
```text
migrate create -ext sql -dir ./migration -seq fill_with_data

migrate -path ./migration -database "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable" up
migrate -path ./migration -database "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable" up 2

migrate -path ./migration -database "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable" down
migrate -path ./migration -database "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable" down 1
```

## sqlx  

Library which provides a set of extensions on go's standard database/sql library  

https://github.com/jmoiron/sqlx  

```shell
go get github.com/jmoiron/sqlx
```


## pq  

Pure Go postgres driver for Go's database/sql package  
https://github.com/lib/pq


## GoDotEnv  

Go (golang) port of the Ruby dotenv project (which loads env vars from a `.env` file)  

https://github.com/joho/godotenv
```shell
go get github.com/joho/godotenv
```


## Logrus

Structured logger for Go (golang), completely API compatible with the standard library logger  

https://github.com/Sirupsen/logrus

```shell
go get github.com/sirupsen/logrus
```


## jwt-go

Go (or 'golang' for search engine friendliness) implementation of JSON Web Tokens  
prev: https://github.com/dgrijalva/jwt-go

current: https://github.com/golang-jwt/jwt
```shell
go get -u github.com/golang-jwt/jwt/v5
```

How to generate jwt  
https://pkg.go.dev/github.com/golang-jwt/jwt/v5  
https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-NewWithClaims-RegisteredClaims  


## swag  

Swag converts Go annotations to Swagger Documentation 2.0  

https://github.com/swaggo/swag

**_Step 1: Install swag_**:  
_Starting in **Go 1.17**, installing executables with go get is deprecated. go install may be used instead:_
```shell
go install github.com/swaggo/swag/cmd/swag@latest
```

**_Step 2: init in the project_**:

Run `swag` init in the project's root folder which contains the main.go file.
This will parse your comments and generate the required files (docs folder and docs/docs.go).

```shell
swag init
```

If `main.go` with General API annotations do not live in project dir, you can let `swag` know with `-g` flag.

```shell
swag init -g cmd/main.go
```

As result will be generated:
```text
docs/
    docs.go
    swagger.json
    swagger.yaml
```

**_!!! PROBLEM RUNNING SWAG COMMAND !!!_**
```text
swag : The term 'swag' is not recognized as the name of a cmdlet, function, script file, or operable program
```

https://stackoverflow.com/questions/73387155/swag-the-term-swag-is-not-recognized-as-the-name-of-a-cmdlet-function-scri

You need to add `$GOPATH/bin` to `$PATH`.  
In linux it would be by adding the following to your bash profile:

```shell
export PATH=$PATH:$(go env GOPATH)/bin
```
Or just add row above to `.envrc` file in project directory


**_Step 3: Install gin-swagger_**  
https://github.com/swaggo/gin-swagger

```shell
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

Import following in your code:
```go
import (
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    _ "github.com/rusrom/yt-todo/docs"
)
```

**_Step 4: Check SWAGGER documentation on URL_**  
http://localhost:8080/swagger/index.html




## SQL QUERIES
```sql
-- List To Do

SELECT * FROM users WHERE username='rusrom';

SELECT * FROM lists tl INNER JOIN users_lists ul ON tl.id = ul.list_id where user_id = 1;

SELECT * FROM lists l INNER JOIN users_lists ul ON l.id = ul.list_id WHERE l.id = 5 AND user_id = 1;

DELETE FROM lists l USING users_lists ul WHERE l.id = ul.list_id AND l.id = 9 AND ul.user_id = 1;

UPDATE lists l SET title='Kharkov', description='bla bla bla' FROM users_lists ul WHERE l.id=10 AND ul.user_id=1 AND ul.list_id=10;



-- List Items

SELECT i.id, i.title, i.description, i.done FROM items i INNER JOIN lists_items li ON i.id = li.item_id WHERE li.list_id = 4 ORDER BY i.id;

SELECT i.id, title, description, done FROM items i
INNER JOIN lists_items li ON i.id = li.item_id
INNER JOIN users_lists ul ON li.list_id = ul.list_id
WHERE li.list_id = 4 and ul.user_id = 1
ORDER BY i.id;


SELECT i.id, i.title, i.description, i.done FROM items i
INNER JOIN lists_items li ON i.id = li.item_id
INNER JOIN users_lists ul ON li.list_id = ul.list_id
WHERE i.id = 6 AND ul.user_id = 1;


DELETE FROM items i
USING lists_items li, users_lists ul
WHERE i.id = li.item_id AND li.list_id = ul.list_id AND li.item_id = 19 AND ul.user_id = 1;



-- PostgreSQL CONCAT Function
-- https://www.postgresqltutorial.com/postgresql-string-functions/postgresql-concat-function/
-- update items set title = concat(id, title) where id in (22, 23, 24, 25, 26, 27, 28, 29);
update items set title = concat_ws('-', id, 'some item', concat('user id: ', 2)) where id in (22, 23, 24, 25, 26, 27, 28, 29);



UPDATE items i SET title = 'Kharkov', description = 'Brave city', done = true
FROM lists_items li, users_lists ul
WHERE i.id = li.item_id and li.list_id = ul.list_id and li.item_id = 22 and ul.user_id = 2;

```
