package repository

import (
	"fmt"

	"github.com/choipopik/todo-app"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int

	createUserQuery := fmt.Sprintf("INSERT INTO %s (name, username, passward_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(createUserQuery, user.Name, user.Username, user.Password)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r AuthPostgres) GetUser(username, password string) (todo.User, error) {
	var user todo.User

	getUserQuery := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND passward_hash=$2", usersTable)
	err := r.db.Get(&user, getUserQuery, username, password)

	return user, err
}
