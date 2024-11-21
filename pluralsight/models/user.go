package models

type User struct {
	ID        int
	FirstName string
}

var (
	users  []*User
	nextID = 1
)
