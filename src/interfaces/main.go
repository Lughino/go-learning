package main

import "fmt"

type bot interface {
	getGreetings() string
}
type englishBot struct{}
type spanishBot struct{}

// in case I don't need to use the receiver I can omit the variable like:
// func (englishBot) getGreetings() string ...
func (eb englishBot) getGreetings() string {
	return "Hello There!"
}
func (sb spanishBot) getGreetings() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreetings())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
