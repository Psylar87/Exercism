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
	firstCard := ParseCard(card1)
	secondCard := ParseCard(card2)
	dealer := ParseCard(dealerCard)

	switch {
	case firstCard == 11 && secondCard == 11:
		return "P"
	case firstCard+secondCard == 21:
		if dealer == 11 || dealer == 10 {
			return "S"
		} else {
			return "W"
		}
	case firstCard+secondCard >= 17 && firstCard+secondCard <= 20:
		return "S"
	case firstCard+secondCard >= 12 && firstCard+secondCard <= 16:
		if dealer >= 7 {
			return "H"
		} else {
			return "S"
		}
	case firstCard+secondCard <= 11:
		return "H"
	default:
		return ""
	}
}
