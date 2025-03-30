package main

import "testdata/multi/mypkg"

func main() {
	users := mypkg.GetUsers()
	println(users[0].Name) // want "potential nil dereference of users[0].Name without prior nil check on element"
}
