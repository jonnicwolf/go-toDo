package main

import "fmt"

func main() {
	todos := Todos{}
	todos.add("but milk")
	todos.add("but bread")
	fmt.Printf("%+v\n\n", todos)
	todos.delete(0)
	fmt.Printf("%+v", todos)
}
