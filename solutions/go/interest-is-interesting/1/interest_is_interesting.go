package interest



func InterestRate(balance float64) float64 {
	var rate float64
	if balance < 0 {
		rate = 3.213
	}
	if balance >= 0 && balance < 1000 {
		rate = 0.5
	}
	if balance >= 1000 && balance < 5000 {
		rate = 1.621
	}
	if balance >= 5000 {
		rate = 2.475
	}
	return rate
}

func Interest(balance float64) float64 {
	return balance / 100 * InterestRate(balance)
}

func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
    const epsilon = 1e-5  // увеличиваем до 0.00001
    years := 0
    for balance+epsilon < targetBalance {
        balance = AnnualBalanceUpdate(balance)
        years++
    }
    return years
}


