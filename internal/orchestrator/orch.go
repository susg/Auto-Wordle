package orchestrator

import (
	"github.com/susg/autowordle/internal/filter"
	"github.com/susg/autowordle/internal/filter/rules"
	"github.com/susg/autowordle/internal/validate"
	"github.com/susg/autowordle/internal/words"
)

type WordleOrchestrator interface {
	GenerateWords(input []string) ([]string, error)
}

type WordleOrchestratorImpl struct {
	wf    filter.WordFilterer
	v     validate.Validator
	words []string
}

func NewWordleOrchestratorImpl(wordLength int, wm words.WordManager, v validate.Validator) WordleOrchestrator {
	words, err := wm.GetWords(wordLength)
	if err != nil {
		panic(err)
	}
	wf := filter.NewWordFiltererImpl(wordLength, &rules.RulesCheckerImpl{})
	return &WordleOrchestratorImpl{
		wf:    wf,
		v:     v,
		words: words,
	}
}

func (w *WordleOrchestratorImpl) GenerateWords(input []string) ([]string, error) {
	if err := w.v.Validate(input); err != nil {
		return nil, err
	}
	words := w.wf.FilterWords(input, w.words)
	w.words = words
	return words, nil
}
