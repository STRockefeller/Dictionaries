package dictionaries

import (
	"github.com/STRockefeller/dictionaries/instances/english/freedictionary"
	"github.com/STRockefeller/dictionaries/instances/japanese/jishoorg"
)

// Interface definition for English Dictionaries which inherits from the freedictionary interface.
type EnglishDictionaries interface {
	freedictionary.Dictionary
}

// Interface definition for Japanese Dictionaries which inherits from the jishoorg interface.
type JapaneseDictionaries interface {
	jishoorg.Dictionary
}

// Creates a new instance of an English Dictionary.
func NewEnglishDictionary[T EnglishDictionaries]() T {
	return *new(T)
}

// Creates a new instance of a Japanese Dictionary.
func NewJapaneseDictionary[T JapaneseDictionaries]() T {
	return *new(T)
}
