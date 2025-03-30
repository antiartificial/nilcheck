package cache

type User struct {
	Name string
}

func getUsers() []*User {
	return []*User{nil, &User{Name: "Alice"}}
}

func main() {
	users := getUsers()
	println(users[0].Name) // want "nil slice element dereference"
}
