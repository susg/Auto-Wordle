package orchestrator

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/susg/autowordle/internal/reader"
	"github.com/susg/autowordle/internal/validate"
	"github.com/susg/autowordle/internal/words"
)

func TestGenerateWords_Success(t *testing.T) {
	wordLength := 5
	wm := words.StartWordManager(reader.NewFileReader())
	v, _ := validate.NewWordleValidator(wordLength)
	orch := NewWordleOrchestratorImpl(wordLength, wm, v)
	words, _ := orch.GenerateWords([]string{"lb", "ib", "kb", "eb", "sy"})
	len1 := len(words)
	assert.True(t, len1 > 0)
	assert.True(t, slices.Contains(words, "study"))
	assert.True(t, slices.Contains(words, "ghost"))
	assert.True(t, slices.Contains(words, "frost"))
	assert.True(t, slices.Contains(words, "boast"))
	assert.False(t, slices.Contains(words, "flask"))
	assert.False(t, slices.Contains(words, "apple"))

	words, _ = orch.GenerateWords([]string{"sy", "ty", "ub", "db", "yb"})
	len2 := len(words)
	assert.True(t, len2 > 0)
	assert.True(t, len2 < len1)
	assert.True(t, slices.Contains(words, "ghost"))
	assert.True(t, slices.Contains(words, "boast"))
	assert.True(t, slices.Contains(words, "frost"))
	assert.True(t, slices.Contains(words, "coast"))
	assert.False(t, slices.Contains(words, "flask"))
	assert.False(t, slices.Contains(words, "stood"))

	words, _ = orch.GenerateWords([]string{"bb", "oy", "ab", "sg", "tg"})
	len3 := len(words)
	assert.True(t, len3 > 0)
	assert.True(t, len3 < len2)
	assert.True(t, slices.Contains(words, "ghost"))
	assert.True(t, slices.Contains(words, "frost"))
	assert.False(t, slices.Contains(words, "coast"))
	assert.False(t, slices.Contains(words, "flask"))
	assert.False(t, slices.Contains(words, "stood"))
}

func TestGenerateWords_ValidationError(t *testing.T) {
	wordLength := 5
	wm := words.StartWordManager(reader.NewFileReader())
	v, _ := validate.NewWordleValidator(wordLength)
	orch := NewWordleOrchestratorImpl(wordLength, wm, v)
	_, err := orch.GenerateWords([]string{"invalid", "input"})
	assert.NotNil(t, err)
}

func TestNewWordleOrchestratorImpl_Panic(t *testing.T) {
	wordLength := 5
	wm := words.StartWordManager(reader.NewFileReader())
	v, _ := validate.NewWordleValidator(wordLength)

	assert.Panics(t, func() {
		NewWordleOrchestratorImpl(0, wm, v)
	}, "Expected panic for invalid word length")
}
