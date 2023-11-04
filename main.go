package main

import "fmt"

const (
	Hearts = iota
	Diamonds
	Clubs
	Spades
)

const (
	Ace = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	Suit int
	Rank int
}

// func cardToMultiHot(cardList []Card) []int {
// 	// Initialize the multi-hot encoding with zeros
// 	cardMultiHot := make([]int, 52)

// 	// Iterate through the card list and set corresponding indices to 1
// 	for _, card := range cardList {
// 		index := (card.Suit-1)*13 + (card.Rank - 1)
// 		cardMultiHot[index] = 1
// 	}

//		return cardMultiHot
//	}
func cardToMultiHot(cardList []Card) []int {
	// Initialize the multi-hot encoding with zeros
	cardMultiHot := make([]int, 52)

	// Iterate through the card list and set corresponding indices to 1
	for _, card := range cardList {
		index := (card.Suit)*13 + (card.Rank - 1)
		cardMultiHot[index] = 1
	}

	return cardMultiHot
}

// UnshuffledDeck generates an unshuffled deck of 52 cards with the typical order.
func UnshuffledDeck() []Card {
	var deck []Card

	// Create the 52-card deck by iterating over suits and ranks.
	for suit := Hearts; suit <= Spades; suit++ {
		for rank := Ace; rank <= King; rank++ {
			card := Card{Suit: suit, Rank: rank}
			deck = append(deck, card)
		}
	}

	return deck
}

func main() {
	// Generate an unshuffled deck.
	deck := UnshuffledDeck()

	// Print the unshuffled deck.
	for i, card := range deck {
		fmt.Printf("Card %d: Suit %d, Rank %d\n", i+1, card.Suit, card.Rank)
	}
}
