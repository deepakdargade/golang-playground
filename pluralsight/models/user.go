package models

type User struct {
	ID        int
	FirstName string
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AdUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}
