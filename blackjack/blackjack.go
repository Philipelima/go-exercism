package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    var value int
	switch  {
        case card == "ace":
         	value = 11
        case card == "two": 
        	value = 2
        case card == "three":
        	value = 3
        case card == "four": 
        	value = 4
        case card == "five": 
        	value = 5
        case card == "six":
        	value = 6
        case card == "seven":
			value = 7
        case card == "eight":
			value = 8
        case card == "nine": 
			value = 9
        case card == "ten" || card == "jack" || card == "queen" || card == "king":
        	value = 10
        default:
        	value = 0
    }
    return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	total := ParseCard(card1) + ParseCard(card2)
    dealerTotal := ParseCard(dealerCard)

    var result string 
    
    switch {
        case card1 == card2 && card1 == "ace": 
        	result = "P"
        case total == 21 && (dealerCard != "ace" && dealerTotal != 10): 
        	result = "W"
        case total == 21:
			result = "S"
        case total == 21 && dealerCard == "ace": 
        	result = "S"
        case total >= 17 && total <= 20: 
			result = "S"
        case total >= 12 && total <= 16 && dealerTotal < 7:
        	result = "S"
        case total >= 12 && total <= 16 && dealerTotal >= 7:
        	result = "H"
        case total <= 11: 
        	result = "H"
    }
    return result
}
