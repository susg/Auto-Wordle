package validate

import (
	"fmt"
	"slices"

	"github.com/susg/autowordle/internal/words"
	//"github.com/susg/autowordle/util"
)

type Validator interface {
	Validate(input []string) error
}

type WordleValidator struct {
	wordLength int
}

func NewWordleValidator(wordLength int) (Validator, error) {
	if !slices.Contains(words.SupportedWordLengths, wordLength) {
		return nil, fmt.Errorf("word length %d is not supported", wordLength)
	}
	return &WordleValidator{wordLength: wordLength}, nil
}

func (wv *WordleValidator) Validate(input []string) error {
	if len(input) != wv.wordLength {
		return fmt.Errorf("input length must be %d characters", wv.wordLength)
	}

	for _, str := range input {
		if len(str) != 2 {
			return fmt.Errorf("invalid input %v", str)
		}

		letter := str[0]
		colour := str[1]
		if letter < 'a' || letter > 'z' {
			return fmt.Errorf("input letter %v does not lie between a-z", string(letter))
		}

		if colour != 'g' && colour != 'y' && colour != 'b' {
			return fmt.Errorf("input colour %v is not valid", string(colour))
		}
	}
	return nil
}
