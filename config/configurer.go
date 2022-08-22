package config

import (
	"encoding/csv"
	"fmt"
	"os"
)

// struct for holding quotes and character from csv file
type quotes struct {
	quote     string
	character string
}

// ConfigureRepository fills in data i.e. quotes in the specified map using either
// hardcoded data, data read from csv file or database.
func ConfigureRepository(source string, destMap *map[string][]string) {

	switch source {
	case SourceLocal:
		fillHardcodedData(destMap)
	case SourceCsv:
		fillCSVData(destMap)
	case SourceDatabase:
		fillDataFromDatabase()
	default:
		fmt.Println("Not implemented")
		// TODO: handle error

	}

}

func fillHardcodedData(destMap *map[string][]string) {

	*destMap = make(map[string][]string, 10)

	obiWanQuotes := []string{
		"Hello there.",
		"Remember, the force will always be with you"}

	lukeSkywalkerQuotes := []string{
		"I won't fail you. I am not afraid.",
		"I am here to rescue you"}

	hanSoloQuotes := []string{
		"Luke, May the force be with you",
		"Never tell me the odds",
		"Laugh it up, fuzzball!"}

	leiaOrganaQuotes := []string{
		"Aren't you too short for a stormtropper?",
		"Help me Obi wan kenobi. you are my only hope."}

	(*destMap)["Obi Wan"] = obiWanQuotes
	(*destMap)["Luke Skywalker"] = lukeSkywalkerQuotes
	(*destMap)["Han Solo"] = hanSoloQuotes
	(*destMap)["Leia Organa"] = leiaOrganaQuotes

}

func fillCSVData(destMap *map[string][]string) {

	file, error := os.Open("./quotes.csv")

	if error != nil {

		// TODO: Panic.
		fmt.Println("Error Occured:", error)

	}

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error opening file:", err)
	} else {
		// TODO: extract data and populate the datastore with the quotes.
		fmt.Println(records)
	}

	// TODO: use quotes from CSV file.
	fillHardcodedData(destMap)
}

// not implemented yet.
func fillDataFromDatabase() {

	// TODO:

}
