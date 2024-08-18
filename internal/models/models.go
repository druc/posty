package models

import "time"

type Post struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
}

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}
