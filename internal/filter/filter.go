package filter

import (
	"github.com/susg/autowordle/internal/filter/rules"
	"github.com/susg/autowordle/internal/models"
)

type WordFilterer interface {
	FilterWords(input []string, words []string) []string
}

type WordFiltererImpl struct {
	wi *models.WordleInfo
	rc rules.RulesChecker
}

func NewWordFiltererImpl(wordLength int, rc rules.RulesChecker) WordFilterer {
	return &WordFiltererImpl{
		wi: models.NewWordleInfo(wordLength),
		rc: rc,
	}
}

func (wfi *WordFiltererImpl) FilterWords(input []string, words []string) []string {
	wfi.wi.Update(input)
	iterations := (len(words) + 99) / 100

	channel := make(chan []string, iterations)
	for i := range iterations {
		start := i * 100
		end := min(start+100, len(words))
		go wfi.filterWordsCore(words[start:end], channel)
	}

	var result []string
	for range iterations {
		res := <-channel
		result = append(result, res...)
	}
	return result
}

func (wfi *WordFiltererImpl) filterWordsCore(words []string, resultChan chan []string) {
	var result []string
	for _, word := range words {
		if wfi.rc.AreRulesSatisfied(*wfi.wi, word) {
			result = append(result, word)
		}
	}
	resultChan <- result
}
