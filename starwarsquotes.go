// Package starwarsquotesgo provides random quote(s) from the star wars universe.
package starwarsquotesgo

// Learning Note: use of "factored" import statement.
import (
	"fmt"
	"math/rand"
	"time"

	"github.com/prathamesh13/starwarsquotesgo/config"
)

var defaultquotes []string
var characterQuoteMap map[string][]string

// Initializes the objects that hold the quotes which are returned when requested by
// the calling module.
func Init(source string) {

	config.ConfigureRepository(source, &characterQuoteMap)

}

// setupValidated checks if the quotes array and character quotes map have been initialized.
// the objects have to be initialized by calling Init from the package.
// this method is not an exported name.
func setupValidated() bool {

	if characterQuoteMap == nil {

		return false
	}

	return true
}

// GetRandomQuotes is an exported name. Any function that is intended to be called
// by code in other modules should have starting letter Capitalized.
func getQuotes() [10]string {

	var quotes [10]string

	quotes[0] = "May the force be with you!"
	quotes[1] = "Remember the force will always be with you"
	quotes[2] = "Never tell me the odds"
	quotes[3] = "Aren't you too short for a stormtropper?"
	quotes[4] = "Try, you must."
	quotes[5] = "Help me Obi wan kenobi. you are my only hope."
	quotes[6] = "Fear leads to anger, anger leads to hate, hate leads to suffering."
	quotes[7] = "Let go of your hate."
	quotes[8] = "Laugh it up, fuzzball!"
	quotes[9] = "This is the way."

	return quotes
}

// GetRandomQuote returns a random quote from star wars universe.
func GetRandomQuote() string {

	rand.Seed(time.Now().UnixNano())

	randomIndex := rand.Intn(10)

	return getQuotes()[randomIndex]
}

func GetRandomQuoteByCharacter(character string) (string, error) {

	if !setupValidated() {
		return Answer, nil
	}

	// TODO:
	// each character can have different number of quotes including empty array.
	// Need to account for that.
	return characterQuoteMap[character][getRandomNumber(2)], nil

}

// GetFormattedQuote returns quote surrounded by double quotes ("") and name of the
// character that said it.
// eg. "Luke, May the force be with you" - Han Solo
func GetFormattedQuote(quote string, character string) string {

	// Go doesn't support default values for parameters, in case we wanted to
	// have the function accept a default value in case the author of the quote
	// is not known.

	var formattedQuote string

	// TODO: use better way/syntax to handle this.
	// Use regex to replace the empty author?
	// variadic functions?
	if character == "" {
		formattedQuote = fmt.Sprintf("\"%v\"", quote)
	} else {
		formattedQuote = fmt.Sprintf("\"%v\" - %v", quote, character)
	}

	return formattedQuote
}

// TODO:
func getRandomNumber(boundry int) int {

	rand.Seed(time.Now().UnixNano())

	return rand.Intn(boundry)

}

// Overloaded functions are not supported in go. Refer : https://go.dev/doc/faq#overloading
// GetRandomQuote returns a random quote by a character from star wars universe specified
// by the parameter charater.
/* func GetRandomQuote(String character) String {

	return "not implemented"

} */
