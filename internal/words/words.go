package words

import (
	"fmt"
	"strings"

	"github.com/susg/autowordle/internal/reader"
)

var SupportedWordLengths = []int{5}
var chunkSize = 1024

type WordManager interface {
	GetWords(wordLength int) ([]string, error)
}

type WordManagerImpl struct {
	wordsCache map[int][]string
}

func StartWordManager(r reader.Reader) WordManager {
	wmi := &WordManagerImpl{
		wordsCache: make(map[int][]string),
	}

	for _, wordLen := range SupportedWordLengths {
		filePath := createFilePath(wordLen)
		fileContent, err := r.ReadFile(filePath, chunkSize)
		if err != nil {
			panic(err)
		}
		words := strings.Split(fileContent, "\n")
		wmi.wordsCache[wordLen] = words
	}
	return wmi
}

func (wmi *WordManagerImpl) GetWords(wordLength int) ([]string, error) {
	if words, ok := wmi.wordsCache[wordLength]; ok {
		return words, nil
	}
	return nil, fmt.Errorf("words not found for length %d", wordLength)
}

func createFilePath(wordLength int) string {
	return fmt.Sprintf("/Users/sushant.gupta/Documents/NotBackedUp/Workspace/auto-wordle/data/prod/%d.txt", wordLength)
}
