package models

import "github.com/susg/autowordle/set"

type WordleInfo struct {
	FixedLetters    map[string]set.Set
	UnfixedLetters  map[string]set.Set
	ExcludedLetters set.Set
	WordLength      int
}

func NewWordleInfo(wordLength int) *WordleInfo {
	return &WordleInfo{
		FixedLetters:    make(map[string]set.Set),
		UnfixedLetters:  make(map[string]set.Set),
		ExcludedLetters: set.New(),
		WordLength:      wordLength,
	}
}

func (wi *WordleInfo) Update(input []string) {
	for i, str := range input {
		letter := string(str[0])
		if str[1] == 'g' {
			if _, exists := wi.FixedLetters[letter]; !exists {
				wi.FixedLetters[letter] = set.New()
			}
			wi.FixedLetters[letter].Insert(i)
			if wi.ExcludedLetters.Contains(letter) {
				wi.ExcludedLetters.Remove(letter)
			}
		} else if str[1] == 'y' {
			if _, exists := wi.UnfixedLetters[letter]; !exists {
				wi.UnfixedLetters[letter] = set.New()
			}
			wi.UnfixedLetters[string(str[0])].Insert(i)
			if wi.ExcludedLetters.Contains(letter) {
				wi.ExcludedLetters.Remove(letter)
			}
		} else {
			_, fixed := wi.FixedLetters[letter]
			_, unfixed := wi.UnfixedLetters[letter]
			if !fixed && !unfixed {
				wi.ExcludedLetters.Insert(letter)
			}
		}
	}
}
