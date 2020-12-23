package main

import "fmt"

func main() {
    //var card string = "Ace of Spades"
    // card type is inferred by the right-hand side
    // : for variable initialization
    card := "Ace of Spades"
    card = "Five of Diamonds "
    fmt.Println(card)
}
