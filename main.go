package main

// import "fmt"

func main() {
	todos := Todos{}
	todos.add("but milk")
	todos.add("but bread")
	todos.toggle(0)
	todos.print()
}