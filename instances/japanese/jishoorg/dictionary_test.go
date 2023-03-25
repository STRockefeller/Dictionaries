package jishoorg

import (
	"net/http"
	"testing"

	"github.com/STRockefeller/dictionaries/instances/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

/* -------------------------------------------------------------------------- */
/*                             connect to real api                            */
/* -------------------------------------------------------------------------- */

func TestDictionary(t *testing.T) {
	assert := assert.New(t)
	res, err := NewDictionary().Search("コード")
	assert.NoError(err)
	assert.Equal("コード", res.Data[0].Japanese[0].Reading)
}

/* -------------------------------------------------------------------------- */
/*                             connect to mock api                            */
/* -------------------------------------------------------------------------- */

type Suite struct {
	testutils.Suite
}

func NewSuite() *Suite {
	return &Suite{
		Suite: *testutils.NewSuite(testutils.PatchHttpRequest(urlAssemble(testWord), http.StatusOK, mockResponseBody)),
	}
}

func TestSuite(t *testing.T) {
	suite.Run(t, NewSuite())
}

func (s *Suite) TestDictionary() {
	assert := s.Assertions
	res, err := NewDictionary().Search("コード")
	assert.NoError(err)
	assert.Equal("コード", res.Data[0].Japanese[0].Reading)
}

const testWord = "コード"

const mockResponseBody = "{\"meta\":{\"status\":200},\"data\":[{\"slug\":\"コード\",\"is_common\":true,\"tags\":[],\"jlpt\":[\"jlpt-n3\"],\"japanese\":[{\"reading\":\"コード\"}],\"senses\":[{\"english_definitions\":[\"code\",\"program\"],\"parts_of_speech\":[\"Noun\"],\"links\":[],\"tags\":[\"Computing\"],\"restrictions\":[],\"see_also\":[],\"antonyms\":[],\"source\":[],\"info\":[],\"sentences\":null},{\"english_definitions\":[\"code\",\"regulation\",\"rules\"],\"parts_of_speech\":[\"Noun\"],\"links\":[],\"tags\":[],\"restrictions\":[],\"see_also\":[],\"antonyms\":[],\"source\":[],\"info\":[],\"sentences\":null},{\"english_definitions\":[\"code\",\"symbol set\"],\"parts_of_speech\":[\"Noun\"],\"links\":[],\"tags\":[],\"restrictions\":[],\"see_also\":[],\"antonyms\":[],\"source\":[],\"info\":[],\"sentences\":null}],\"attribution\":{\"jmdict\":true,\"jmnedict\":false,\"dbpedia\":false}},{\"slug\":\"コード-1\",\"is_common\":false,\"tags\":[],\"jlpt\":[],\"japanese\":[{\"reading\":\"コード\"}],\"senses\":[{\"english_definitions\":[\"cord\",\"flex\"],\"parts_of_speech\":[\"Noun\"],\"links\":[],\"tags\":[],\"restrictions\":[],\"see_also\":[],\"antonyms\":[],\"source\":[],\"info\":[],\"sentences\":null}],\"attribution\":{\"jmdict\":true,\"jmnedict\":false,\"dbpedia\":false}},{\"slug\":\"コード-2\",\"is_common\":false,\"tags\":[],\"jlpt\":[],\"japanese\":[{\"reading\":\"コード\"}],\"senses\":[{\"english_definitions\":[\"chord\"],\"parts_of_speech\":[\"Noun\"],\"links\":[],\"tags\":[\"Music\"],\"restrictions\":[],\"see_also\":[\"和音\"],\"antonyms\":[],\"source\":[],\"info\":[],\"sentences\":null},{\"english_definitions\":[\"string (on an instrument)\"],\"parts_of_speech\":[\"Noun\"],\"links\":[],\"tags\":[\"Music\"],\"restrictions\":[],\"see_also\":[],\"antonyms\":[],\"source\":[],\"info\":[],\"sentences\":null}],\"attribution\":{\"jmdict\":true,\"jmnedict\":false,\"dbpedia\":false}},{\"slug\":\"コードレス\",\"is_common\":true,\"tags\":[],\"jlpt\":[],\"japanese\":[{\"reading\":\"コードレス\"}],\"senses\":[{\"english_definitions\":[\"cordless\"],\"parts_of_speech\":[\"Noun which may take the genitive case particle 'no'\"],\"links\":[],\"tags\":[],\"restrictions\":[],\"see_also\":[],\"antonyms\":[],\"source\":[],\"info\":[],\"sentences\":null}],\"attribution\":{\"jmdict\":true,\"jmnedict\":false,\"dbpedia\":false}},{\"slug\":\"コードレスホン\",\"is_common\":true,\"tags\":[],\"jlpt\":[],\"japanese\":[{\"reading\":\"コードレスホン\"},{\"reading\":\"コードレス・ホン\"}],\"senses\":[{\"english_definitions\":[\"cordless phone\"],\"parts_of_speech\":[\"Noun\"],\"links\":[],\"tags\":[],\"restrictions\":[],\"see_also\":[],\"antonyms\":[],\"source\":[],\"info\":[],\"sentences\":null}],\"attribution\":{\"jmdict\":true,\"jmnedict\":false,\"dbpedia\":false}},{\"slug\":\"コードページ\",\"is_common\":false,\"tags\":[],\"jlpt\":[],\"japanese\":[{\"reading\":\"コードページ\"},{\"reading\":\"コード・ページ\"}],\"senses\":[{\"english_definitions\":[\"code page\"],\"parts_of_speech\":[\"Noun\"],\"links\":[],\"tags\":[\"Computing\"],\"restrictions\":[],\"see_also\":[],\"antonyms\":[],\"source\":[],\"info\":[],\"sentences\":null},{\"english_definitions\":[\"Code page\"],\"parts_of_speech\":[\"Wikipedia definition\"],\"links\":[{\"text\":\"Read “Code page” on English Wikipedia\",\"url\":\"http://en.wikipedia.org/wiki/Code_page?oldid=492193204\"},{\"text\":\"Read “コードページ” on Japanese Wikipedia\",\"url\":\"http://ja.wikipedia.org/wiki/コードページ?oldid=37042512\"}],\"tags\":[],\"restrictions\":[],\"see_also\":[],\"antonyms\":[],\"source\":[],\"info\":[],\"sentences\":[]}],\"attribution\":{\"jmdict\":true,\"jmnedict\":false,\"dbpedia\":\"http://dbpedia.org/resource/Code_page\"}},{\"slug\":\"コードバン\",\"is_common\":false,\"tags\":[],\"jlpt\":[],\"japanese\":[{\"reading\":\"コードバン\"}],\"senses\":[{\"english_definitions\":[\"cordovan (type of leather)\"],\"parts_of_speech\":[\"Noun\"],\"links\":[],\"tags\":[],\"restrictions\":[],\"see_also\":[],\"antonyms\":[],\"source\":[],\"info\":[],\"sentences\":null},{\"english_definitions\":[\"Shell cordovan\"],\"parts_of_speech\":[\"Wikipedia definition\"],\"links\":[{\"text\":\"Read “Shell cordovan” on English Wikipedia\",\"url\":\"http://en.wikipedia.org/wiki/Shell_cordovan?oldid=471647726\"},{\"text\":\"Read “コードバン” on Japanese Wikipedia\",\"url\":\"http://ja.wikipedia.org/wiki/コードバン?oldid=40642535\"}],\"tags\":[],\"restrictions\":[],\"see_also\":[],\"antonyms\":[],\"source\":[],\"info\":[],\"sentences\":[]}],\"attribution...+2338 more"
