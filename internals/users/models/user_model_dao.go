package models

type UserDAO struct {
	ID       int    `db:"id"`
	UserName string `db:"user_name"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
