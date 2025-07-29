package repository

import (
	_ "embed"

	"github.com/choipopik/todo-app"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

var (
	//go:embed queries/create_user_query.sql
	createUserQuery string
	//go:embed queries/get_user_query.sql
	getUserQuery string
)

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int

	row := r.db.QueryRow(createUserQuery, user.Name, user.Username, user.Password)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r AuthPostgres) GetUser(username, password string) (todo.User, error) {
	var user todo.User

	err := r.db.Get(&user, getUserQuery, username, password)

	return user, err
}
