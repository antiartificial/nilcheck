package chained

type User struct {
	Address *Address
}

type Address struct {
	City string
}

func (u *User) GetAddress() *Address {
	return u.Address
}

func getUser() *User {
	return &User{Address: nil}
}

func main() {
	u := getUser()
	println(u.GetAddress().City) // want "potential nil dereference in method call u.GetAddress without prior nil check"
}
