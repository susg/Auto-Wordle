package filter

import (
	"slices"
	"testing"

	"github.com/susg/autowordle/internal/config"
	"github.com/susg/autowordle/internal/filter/rules"
)

func TestNewWordFiltererImpl(t *testing.T) {
	cfg := config.Config{}
	rc := &rules.RulesCheckerImpl{}

	wf := NewWordFiltererImpl(5, rc, cfg)

	if wf == nil {
		t.Error("Expected WordFilterer instance, got nil")
	}

	impl, ok := wf.(*WordFiltererImpl)
	if !ok {
		t.Error("Expected WordFiltererImpl type")
	}

	if impl.wi == nil {
		t.Error("Expected WordleInfo to be initialized")
	}

	if impl.rc != rc {
		t.Error("Expected RulesChecker to be set correctly")
	}
}

func TestFilterWords(t *testing.T) {
	tests := []struct {
		name           string
		input          []string
		words          []string
		satisfiedWords map[string]bool
		expected       []string
	}{
		{
			name:     "empty words list",
			input:    []string{"ag", "pb", "pb", "ly", "eb"},
			words:    []string{},
			expected: []string{},
		},
		{
			name:     "all words satisfy rules",
			input:    []string{"ag", "pb", "pb", "ly", "eb"},
			words:    []string{"aloft", "allow", "altar"},
			expected: []string{"aloft", "allow", "altar"},
		},
		{
			name:     "some words satisfy rules",
			input:    []string{"ag", "pb", "pb", "ly", "eb"},
			words:    []string{"plain", "aloft", "plant", "black", "allow", "altar", "apple"},
			expected: []string{"aloft", "allow", "altar"},
		},
		{
			name:     "no words satisfy rules",
			input:    []string{"ag", "pb", "pb", "ly", "eb"},
			words:    []string{"plain", "plant", "black", "apple"},
			expected: []string{},
		},
		{
			name:  "large word list for concurrency test",
			input: []string{"ag", "pb", "pb", "ly", "eb"},
			words: func() []string {
				words := make([]string, 250)
				words[0] = "aloft"
				words[1] = "allow"
				words[2] = "altar"
				for i := 3; i < 250; i++ {
					words[i] = "word" + string(rune('a'+i%26))
				}
				return words
			}(),
			expected: []string{"aloft", "allow", "altar"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.GetConfig()
			rc := &rules.RulesCheckerImpl{}
			wf := NewWordFiltererImpl(5, rc, cfg)

			result := wf.FilterWords(tt.input, tt.words)

			if len(result) != len(tt.expected) {
				t.Errorf("Expected %d words, got %d", len(tt.expected), len(result))
			}

			for _, word := range result {
				if !slices.Contains(tt.expected, word) {
					t.Errorf("Unexpected word %s found in result", word)
				}
			}
		})
	}
}

func TestFilterWordsCore(t *testing.T) {
	cfg := config.GetConfig()
	rc := &rules.RulesCheckerImpl{}

	wf := NewWordFiltererImpl(5, rc, cfg).(*WordFiltererImpl)
	wf.wi.Update([]string{"ag", "pb", "pb", "ly", "eb"})

	words := []string{"plain", "aloft", "plant", "black", "allow", "altar", "apple"}
	resultChan := make(chan []string, 1)

	wf.filterWordsCore(words, resultChan)

	result := <-resultChan
	expected := []string{"aloft", "allow", "altar"}

	if len(result) != len(expected) {
		t.Errorf("Expected %d words, got %d", len(expected), len(result))
	}

	for i, word := range result {
		if word != expected[i] {
			t.Errorf("Expected word %s at position %d, got %s", expected[i], i, word)
		}
	}
}
