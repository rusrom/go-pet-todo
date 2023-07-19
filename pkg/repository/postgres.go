package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable      = "users"
	itemsTable      = "items"
	listsTable      = "lists"
	listsItemsTable = "lists_items"
	usersListsTable = "users_lists"
)

type DbConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SSLMode  string
}

func NewDb(c DbConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open(
		"postgres", fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			c.Host, c.Port, c.Username, c.Database, c.Password, c.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
