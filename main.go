package main

import (
	"fmt"
	"pkg"
)

func main() {
	card := types.Card {
		ID: 1425
		Balance: 55.5
		PAN: "5505 ****"
		Currency: types.TJS
	}
	const cardId = card.Id
}