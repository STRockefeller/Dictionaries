# Dictionaries

Dictionaries is a Go package that provides access to dictionaries for different languages.

## Installation

``` bash
go get github.com/STRockefeller/dictionaries
```

## Usage

This package exposes two interfaces: EnglishDictionaries and JapaneseDictionaries. They provide access to the following dictionaries:

* English : FreeDictionary
* Japanese: JishoOrg

## Example

```go
package main

import (
	"fmt"

	"github.com/STRockefeller/dictionaries"
	"github.com/STRockefeller/dictionaries/instances/english/freedictionary"
	"github.com/STRockefeller/dictionaries/instances/japanese/jishoorg"
)

func main() {
	// English dictionary
	{
		dictionary := dictionaries.NewEnglishDictionary[freedictionary.Dictionary]()
		definition, err := dictionary.Search("word")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(definition)
		}
	}

	// Japanese dictionary
	{
		dictionary := dictionaries.NewJapaneseDictionary[jishoorg.Dictionary]()
		definition, err := dictionary.Search("単語")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(definition)
		}
	}
}

```

In the above example, we create instances of both English and Japanese dictionaries and get their respective definitions.
