package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerValue := ParseCard(card1) + ParseCard(card2)
	dealerValue := ParseCard(dealerCard)

	// Check for pair of aces
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}

	// Check for blackjack (21 with ace and 10-value card)
	if playerValue == 21 {
		if dealerValue == 11 || dealerValue == 10 { // dealer has ace or 10-value
			return "S" // Stand
		}
		return "W" // Win
	}

	// Hard totals logic
	if playerValue >= 17 {
		return "S" // Stand
	}

	if playerValue >= 12 && playerValue <= 16 {
		if dealerValue >= 7 { // dealer has 7 or higher
			return "H" // Hit
		}
		return "S" // Stand
	}

	if playerValue <= 11 {
		return "H" // Hit
	}

	// Default case (shouldn't reach here with valid input)
	return "S"
}
