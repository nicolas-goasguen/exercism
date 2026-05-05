package blackjack

const (
	two   = "two"
	three = "three"
	four  = "four"
	five  = "five"
	six   = "six"
	seven = "seven"
	eight = "eight"
	nine  = "nine"
	ten   = "ten"
	jack  = "jack"
	queen = "queen"
	king  = "king"
	ace   = "ace"
)

const (
	decisionStand = "S"
	decisionHit   = "H"
	decisionSplit = "P"
	decisionWin   = "W"
)

const otherValue = 0

var cardValue = map[string]int{
	two:   2,
	three: 3,
	four:  4,
	five:  5,
	six:   6,
	seven: 7,
	eight: 8,
	nine:  9,
	ten:   10,
	jack:  10,
	queen: 10,
	king:  10,
	ace:   11,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	v, exists := cardValue[card]
	if !exists {
		return otherValue
	}
	return v
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	sum := ParseCard(card1) + ParseCard(card2)
	switch {
	case sum == 22:
		return decisionSplit
	case sum == 21:
		if ParseCard(dealerCard) < 10 {
			return decisionWin
		} else {
			return decisionStand
		}
	case sum >= 17:
		return decisionStand
	case sum >= 12:
		if ParseCard(dealerCard) < 7 {
			return decisionStand
		} else {
			return decisionHit
		}
	default:
		return decisionHit
	}
}
