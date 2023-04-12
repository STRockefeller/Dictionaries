package freedictionary

import (
	"net/http"
	"testing"

	"github.com/STRockefeller/dictionaries/instances/internal/testutils"
	rsuite "github.com/STRockefeller/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

/* -------------------------------------------------------------------------- */
/*                             connect to real api                            */
/* -------------------------------------------------------------------------- */

func TestDictionary(t *testing.T) {
	assert := assert.New(t)
	res, err := NewDictionary().Search("code")
	assert.NoError(err)
	assert.Equal("code", res[0].Word)
}

/* -------------------------------------------------------------------------- */
/*                             connect to mock api                            */
/* -------------------------------------------------------------------------- */

type Suite struct {
	rsuite.SuiteTemplate
}

func NewSuite() *Suite {
	return &Suite{
		SuiteTemplate: *rsuite.NewSuiteTemplate().AddSetupTestFuncs(testutils.PatchHttpRequest(urlAssemble(testWord), http.StatusOK, mockResponseBody)),
	}
}

func TestSuite(t *testing.T) {
	suite.Run(t, NewSuite())
}

func (s *Suite) TestDictionary() {
	assert := s.Assertions
	res, err := NewDictionary().Search("code")
	assert.NoError(err)
	assert.Equal("code", res[0].Word)
}

const testWord = "code"
const mockResponseBody = "[{\"word\":\"code\",\"phonetic\":\"/kəʊd/\",\"phonetics\":[{\"text\":\"/kəʊd/\",\"audio\":\"https://api.dictionaryapi.dev/media/pronunciations/en/code-uk.mp3\"},{\"text\":\"/koʊd/\",\"audio\":\"\"}],\"origin\":\"\",\"meanings\":[{\"partOfSpeech\":\"noun\",\"definitions\":[{\"definition\":\"A short symbol, often with little relation to the item it represents.\",\"example\":\"This flavour of soup has been assigned the code WRT-9.\",\"synonyms\":[],\"antonyms\":[]},{\"definition\":\"A body of law, sanctioned by legislation, in which the rules of law to be specifically applied by the courts are set forth in systematic form; a compilation of laws by public authority; a digest.\",\"example\":\"\",\"synonyms\":[],\"antonyms\":[]},{\"definition\":\"Any system of principles, rules or regulations relating to one subject.\",\"example\":\"The medical code is a system of rules for the regulation of the professional conduct of physicians.\",\"synonyms\":[],\"antonyms\":[]},{\"definition\":\"A set of rules for converting information into another form or representation.\",\"example\":\"\",\"synonyms\":[],\"antonyms\":[]},{\"definition\":\"A message represented by rules intended to conceal its meaning.\",\"example\":\"\",\"synonyms\":[],\"antonyms\":[]},{\"definition\":\"A cryptographic system using a codebook that converts words or phrases into codewords.\",\"example\":\"\",\"synonyms\":[],\"antonyms\":[]},{\"definition\":\"Instructions for a computer, written in a programming language; the input of a translator, an interpreter or a browser, namely: source code, machine code, bytecode.\",\"example\":\"I wrote some code to reformat text documents.\",\"synonyms\":[],\"antonyms\":[]},{\"definition\":\"(scientific programming) A program.\",\"example\":\"\",\"synonyms\":[],\"antonyms\":[]},{\"definition\":\"A particular lect or language variety.\",\"example\":\"\",\"synonyms\":[],\"antonyms\":[]},{\"definition\":\"An emergency requiring situation-trained members of the staff.\",\"example\":\"\",\"synonyms\":[],\"antonyms\":[]}]},{\"partOfSpeech\":\"verb\",\"definitions\":[{\"definition\":\"To write software programs.\",\"example\":\"I learned to code on an early home computer in the 1980s.\",\"synonyms\":[],\"antonyms\":[]},{\"definition\":\"To add codes to a dataset.\",\"example\":\"\",\"synonyms\":[],\"antonyms\":[]},{\"definition\":\"To categorise by assigning identifiers from a schedule, for example CPT coding for medical insurance purposes.\",\"example\":\"\",\"synonyms\":[],\"antonyms\":[]},{\"definition\":\"To encode.\",\"example\":\"We should code the messages we send out on Usenet.\",\"synonyms\":[],\"antonyms\":[]},{\"definition\":\"To encode a protein.\",\"example\":\"\",\"synonyms\":[],\"antonyms\":[]},{\"definition\":\"To call a hospital emergency code.\",\"example\":\"coding in the CT scanner\",\"synonyms\":[],\"antonyms\":[]}]}]},{\"word\":\"code\",\"phonetic\":\"/kəʊd/\",\"phonetics\":[{\"text\":\"/kəʊd/\",\"audio\":\"https://api.dictionaryapi.dev/media/pronunciations/en/code-uk.mp3\"},{\"text\":\"/koʊd/\",\"audio\":\"\"}],\"origin\":\"\",\"meanings\":[{\"partOfSpeech\":\"verb\",\"definitions\":[{\"definition\":\"Of a patient, to suffer a sudden medical emergency (a code blue) such as cardiac arrest.\",\"example\":\"\",\"synonyms\":[],\"antonyms\":[]}]}]}]"
