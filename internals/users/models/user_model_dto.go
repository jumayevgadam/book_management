package models

type User struct {
	ID       int    `db:"id"`
	UserName string `db:"user_name"`
	Email    string `db:"email"`
	Password string `db:"-"`
}
