package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type User struct {
	Person
	Id int
}

func main() {
	u := User{}
	u.Name = "Adam" // u.Person.Name = "Adam"
	u.Age = 42      //u.Person.Age = 42
	u.Id = 1
	fmt.Println(u)
}
