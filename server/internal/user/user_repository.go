package user

import (
	"context"
	"database/sql"
)

type DBTx interface {
	ExecContext(ctx context.Context, query string, arg ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type respository struct {
	db DBTx
}

func NewRepository(db DBTx) Respository {
	return &respository{db: db}
}

func (r *respository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var lastInsertId int
	query := "INSERT INTO chatusers(username,password,email) VALUES ($1,$2,$3) returning id"
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Email).Scan(&lastInsertId)

	if err != nil {
		return &User{}, err
	}

	user.ID = int64(lastInsertId)
	return user, nil
}

func (r *respository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	u := User{}
	query := "SELECT id,email,username,password FROM chatusers WHERE email=$1"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Username, &u.Email, &u.Password)

	if err != nil {
		return &User{}, nil
	}

	return &u, nil
}
