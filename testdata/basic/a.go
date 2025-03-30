package basic

type User struct {
	Name string
}

func getUser() *User {
	return nil
}

func main() {
	u := getUser()
	println(u.Name) // want "potential nil dereference of u.Name without prior nil check"

	if u != nil {
		println(u.Name) // Safe
	}
}
