package shield

import (
  "io/ioutil"
	"regexp"
	"strings"
)

type enTokenizer struct {
}

type ptBRTokenizer struct {
}

func NewEnglishTokenizer() Tokenizer {
	return &enTokenizer{}
}

func NewPortugueseTokenizer() Tokenizer {
  return &ptBRTokenizer{}
}

func PerformTokenization(text string, splitToken *regexp.Regexp) (words map[string]int64) {
  words = make(map[string]int64)
  for _, w := range splitToken.Split(text, -1) {
    if len(w) > 2 {
      words[strings.ToLower(w)]++
    }
  }
  return
}

func LoadStopListForLocale(locale string) []string {
  fileBytes, _ := ioutil.ReadFile("./stoplists/" + locale + ".txt")
  fileContent := string(fileBytes)

  return strings.Split(fileContent, "\n")
}

// Spamassassin stoplist
//
// http://wiki.apache.org/spamassassin/BayesStopList
func (t *enTokenizer) Tokenize(text string) map[string]int64 {
  return PerformTokenization(text, enToken)
}

func (t *ptBRTokenizer) Tokenize(text string) map[string]int64 {
  return PerformTokenization(text, ptBRToken)
}

var enWords = LoadStopListForLocale("en")
var enToken = regexp.MustCompile(`\b([^\w]+|` + strings.Join(enWords, "|") + `)\b`)

var ptBRWords = LoadStopListForLocale("pt-BR")
var ptBRToken = regexp.MustCompile(`\b([^\w]+|` + strings.Join(ptBRWords, "|") + `)\b`)
