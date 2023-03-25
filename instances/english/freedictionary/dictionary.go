package freedictionary

import (
	"fmt"

	"github.com/STRockefeller/dictionaries/instances/internal/apicall"
	"github.com/STRockefeller/dictionaries/instances/internal/utils"
)

type Dictionary struct{}

func NewDictionary() *Dictionary { return &Dictionary{} }

func (d Dictionary) Search(word string) (Result, error) {
	return apicall.Call(apicall.Get(urlAssemble(word)), utils.ParseResponseBody[Result])
}

func urlAssemble(word string) string {
	return fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%s", word)
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
