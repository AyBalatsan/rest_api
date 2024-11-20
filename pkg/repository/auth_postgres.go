package repository

import (
	"fmt"

	todo "github.com/AyBalatsan/Rest_API"
	"github.com/go-sqlx/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	// тут мы записываем переменную id
	if err := row.Scan(&id); err != nil {
		return 0, nil
	}
	return id, nil //добавить вывод токена
}

func (r *AuthPostgres) GetUser(username, password string) (todo.User, error) {
	var user todo.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	// Получаем из базы
	err := r.db.Get(&user, query, username, password)

	return user, err // Почему мы где то возвращаем err ? err.ERROR() ? logrus.Error(err)
}