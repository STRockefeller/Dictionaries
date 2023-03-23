package dictionaries

import (
	"github.com/STRockefeller/dictionaries/instances/english/freedictionary"
	"github.com/STRockefeller/dictionaries/instances/japanese/jishoorg"
)

type EnglishDictionaries interface {
	freedictionary.Dictionary
}

type JapaneseDictionaries interface {
	jishoorg.Dictionary
}

func NewEnglishDictionary[T EnglishDictionaries]() T {
	return *new(T)
}

func NewJapaneseDictionary[T JapaneseDictionaries]() T {
	return *new(T)
}
