package rules

import (
	"strings"

	"github.com/susg/autowordle/internal/models"
)

type RulesChecker interface {
	AreRulesSatisfied(wi models.WordleInfo, word string) bool
}

type RulesCheckerImpl struct{}

func (rci *RulesCheckerImpl) AreRulesSatisfied(wi models.WordleInfo, word string) bool {
	if len(word) != wi.WordLength {
		return false
	}

	return rci.excludedLettersRule(wi, word) && rci.fixedLettersRule(wi, word) && rci.unfixedLettersRule(wi, word)
}

func (rci *RulesCheckerImpl) excludedLettersRule(wi models.WordleInfo, word string) bool {
	for _, letter := range word {
		if wi.ExcludedLetters.Contains(string(letter)) {
			return false
		}
	}
	return true
}

func (rci *RulesCheckerImpl) fixedLettersRule(wi models.WordleInfo, word string) bool {
	for letter, positionsSet := range wi.FixedLetters {
		positions := positionsSet.GetAll()
		for _, pos := range positions {
			if string(word[pos.(int)]) != letter {
				return false
			}
		}
	}
	return true
}

func (rci *RulesCheckerImpl) unfixedLettersRule(wi models.WordleInfo, word string) bool {
	for letter, positionsSet := range wi.UnfixedLetters {
		if !strings.Contains(word, letter) {
			return false
		}

		positions := positionsSet.GetAll()
		for _, pos := range positions {
			if string(word[pos.(int)]) == letter {
				return false
			}
		}
	}
	return true
}
