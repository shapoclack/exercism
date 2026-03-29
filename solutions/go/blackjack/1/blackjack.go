package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var value int
    switch card{
        case "ace":
        	value = 11
        case "two":
        	value = 2
        case "three":
        	value = 3
        case "four":
        	value = 4
        case "five":
        	value = 5
        case "six":
        	value = 6
        case "seven":
        	value = 7
        case "eight":
        	value = 8
        case "nine":
        	value = 9
        case "ten":
        	value = 10
        case "jack":
        	value = 10
        case "queen":
        	value = 10
        case "king":
        	value = 10
        default:
        	value = 0
    }
    return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	value1 := ParseCard(card1)
	value2 := ParseCard(card2)
	dealerValue := ParseCard(dealerCard)
	
	// Общая сумма карт игрока
	handTotal := value1 + value2
	
	// Правило 1: Если у вас пара тузов - всегда сплит
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}
	
	// Правило 2: Если у вас блэкджек (21)
	if handTotal == 21 {
		// Автоматическая победа, если у дилера нет туза, фигуры или десятки
		if dealerValue < 10 {
			return "W"
		}
		// Иначе - стоять
		return "S"
	}
	
	// Правило 3: Если сумма в диапазоне [17, 20] - всегда стоять
	if handTotal >= 17 && handTotal <= 20 {
		return "S"
	}
	
	// Правило 4: Если сумма в диапазоне [12, 16]
	if handTotal >= 12 && handTotal <= 16 {
		// Стоять, если у дилера меньше 7, иначе брать карту
		if dealerValue < 7 {
			return "S"
		}
		return "H"
	}
	
	// Правило 5: Если сумма 11 или меньше - всегда брать карту
	if handTotal <= 11 {
		return "H"
	}
	
	return "S" // На всякий случай
}
