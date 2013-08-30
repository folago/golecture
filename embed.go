package main

import "fmt"

type Person struct {
    Name string
    Age int
}

type User struct{
    Person
    Id int
}

func main() {
    a := User{}
    a.Name = "Adam"
    a.Age = 42
    a.Id = 1
    fmt.Println(a)
}
