package slice

type User struct {
	Name string
}

func getUsers() []*User {
	return []*User{nil, &User{Name: "Alice"}}
}

func getSafeUsers() []*User {
	return []*User{&User{Name: "Alice"}}
}

func main() {
	users := getUsers()
	println(users[0].Name) // want "potential nil dereference of users[0].Name without prior nil check on element"

	safeUsers := getSafeUsers()
	println(safeUsers[0].Name) // Safe
}
