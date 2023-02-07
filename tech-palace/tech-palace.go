package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	customerName := strings.ToUpper(customer)
	return "Welcome to the Tech Palace, "+customerName
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	starsInLine := strings.Repeat("*", numStarsPerLine)
    return starsInLine +"\n"+welcomeMsg+"\n"+starsInLine
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	removedBorder := strings.ReplaceAll(oldMsg, "*", "")
    newMsg := strings.Trim(removedBorder, "\n ") 

    return newMsg
}
