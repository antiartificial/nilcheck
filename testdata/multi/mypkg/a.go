package mypkg

type User struct {
	Name string
}

func GetUsers() []*User {
	return []*User{nil, &User{Name: "Alice"}}
}
