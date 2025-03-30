package edge

type User struct {
	Name string
}

func getUser() *User {
	return nil
}

func main() {
	// Empty file case (no dereference)
	u := getUser()
	_ = u // No warning expected
}
