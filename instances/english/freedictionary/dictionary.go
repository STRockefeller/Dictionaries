package freedictionary

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/STRockefeller/dictionaries/instances/internal/apicall"
)

type Dictionary struct{}

func NewDictionary() *Dictionary { return &Dictionary{} }

func (d Dictionary) Search(word string) (Result, error) {
	return apicall.Call(apicall.Get(fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%s", word)), parseResponse)

}

func parseResponse(resp *http.Response) (Result, error) {
	var result Result
	err := json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}

type Result []ResultElement

type ResultElement struct {
	Word      string     `json:"word"`
	Phonetic  string     `json:"phonetic"`
	Phonetics []Phonetic `json:"phonetics"`
	Origin    string     `json:"origin"`
	Meanings  []Meaning  `json:"meanings"`
}

type Meaning struct {
	PartOfSpeech string       `json:"partOfSpeech"`
	Definitions  []Definition `json:"definitions"`
}

type Definition struct {
	Definition string `json:"definition"`
	Example    string `json:"example"`
	Synonyms   []any  `json:"synonyms"`
	Antonyms   []any  `json:"antonyms"`
}

type Phonetic struct {
	Text  string  `json:"text"`
	Audio *string `json:"audio,omitempty"`
}
