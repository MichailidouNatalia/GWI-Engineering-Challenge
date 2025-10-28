package user

import "time"

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string // hashed
	CreatedAt time.Time
	UpdateAt  time.Time
}
