package shield

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	tokenizer := NewEnglishTokenizer()
	text := "lorem    ipsum able hello erik    can do hi there  \t  spaaace! lorem"
	m := tokenizer.Tokenize(text)

  expectedMap := map[string]int64{ "lorem": 2, "ipsum": 1, "hello": 1, "erik": 1, "spaaace": 1 }
  if errorMessage := compareMaps(m, expectedMap); errorMessage != "" {
    t.Fatal(errorMessage)
  }
}

func TestTokenizeWithPartialMatch(t *testing.T) {
  tokenizer := NewEnglishTokenizer()
  text := "lorem    ipsum able hello erik   cannon can do hi there  \t  spaaace! lorem"
  m := tokenizer.Tokenize(text)

  expectedMap := map[string]int64{ "lorem": 2, "ipsum": 1, "hello": 1, "erik": 1, "spaaace": 1, "cannon": 1 }
  if errorMessage := compareMaps(m, expectedMap); errorMessage != "" {
    t.Fatal(errorMessage)
  }
}
