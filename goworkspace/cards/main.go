package main

func main() {
	//cards := newDeck()
	//fmt.Println(cards.toString())
	//cards.saveToFile("bytecards")
	cards := newDeckFromFile("bytecards")
	cards.shuffle()
	cards.print()

}
