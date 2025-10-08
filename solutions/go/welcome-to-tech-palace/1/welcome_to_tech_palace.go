package techpalace
import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	doom := "Welcome to the Tech Palace, " + strings.ToUpper(customer)
	return doom
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	str := strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine)
	return str
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	oldMsg = strings.Replace(oldMsg, "*", "", -1)
	oldMsg = strings.TrimSpace(oldMsg)
	return oldMsg
}
