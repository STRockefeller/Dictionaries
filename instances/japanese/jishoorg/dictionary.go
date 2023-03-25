package jishoorg

import (
	"fmt"
	"net/url"

	"github.com/STRockefeller/dictionaries/instances/internal/apicall"
	"github.com/STRockefeller/dictionaries/instances/internal/utils"
)

type Dictionary struct{}

func NewDictionary() *Dictionary { return &Dictionary{} }

func (d Dictionary) Search(word string) (Result, error) {
	return apicall.Call(apicall.Get(urlAssemble(word)), utils.ParseResponseBody[Result])
}

func urlAssemble(word string) string {
	return fmt.Sprintf("https://jisho.org/api/v1/search/words?keyword=%s", url.QueryEscape(word))
}

type Result struct {
	Meta Meta   `json:"meta"`
	Data []Data `json:"data"`
}

type Meta struct {
	Status int `json:"status"`
}

type Data struct {
	Slug        string      `json:"slug"`
	IsCommon    bool        `json:"is_common"`
	Tags        []string    `json:"tags"`
	Jlpt        []string    `json:"jlpt"`
	Japanese    []Japanese  `json:"japanese"`
	Senses      []Sense     `json:"senses"`
	Attribution Attribution `json:"attribution"`
}

type Japanese struct {
	Word    string `json:"word,omitempty"`
	Reading string `json:"reading"`
}

type Sense struct {
	EnglishDefinitions []string   `json:"english_definitions"`
	PartsOfSpeech      []string   `json:"parts_of_speech"`
	Links              []Link     `json:"links"`
	Tags               []string   `json:"tags"`
	Restrictions       []string   `json:"restrictions"`
	SeeAlso            []string   `json:"see_also"`
	Antonyms           []string   `json:"antonyms"`
	Source             []string   `json:"source"`
	Info               []string   `json:"info"`
	Sentences          []Sentence `json:"sentences"`
}

type Link struct {
	Text string `json:"text"`
	URL  string `json:"url"`
}

type Sentence struct {
	English  string `json:"english"`
	Japanese string `json:"japanese"`
}

type Attribution struct {
	Jmdict   bool `json:"jmdict"`
	Jmnedict bool `json:"jmnedict"`
	Dbpedia  any  `json:"dbpedia"`
}
