package main

import "fmt"

func main() {
	eb := englishBot{}
	printGreeting(eb)

}

type bot interface {
	getGreeting() string
}
type englishBot struct{}

type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "Hello!!"
}
func (spanishBot) getGreeting() string {
	return "Holaa!!"
}
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
