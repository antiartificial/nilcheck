package method

type User struct {
	Name string
}

func (u *User) GetName() string {
	return u.Name // want "potential nil dereference"
}

func (u *User) SafeGetName() string {
	if u != nil {
		return u.Name
	}
	return ""
}

func getUser() *User {
	return nil
}

func main() {
	u := getUser()
	u.GetName()     // want "potential nil dereference in method call u.GetName without prior nil check"
	u.SafeGetName() // Safe
}
