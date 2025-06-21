package words

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/susg/autowordle/internal/config"
	"github.com/susg/autowordle/internal/reader"
	"github.com/susg/autowordle/utils"
)

type WordManager interface {
	GetWords(wordLength int) ([]string, error)
}

type WordManagerImpl struct {
	wordsCache map[int][]string
}

func StartWordManager(r reader.Reader, cfg config.Config) WordManager {
	wmi := &WordManagerImpl{
		wordsCache: make(map[int][]string),
	}

	for _, wordLen := range cfg.SupportedWordLengths {
		filePath := createFilePath(wordLen, cfg)
		fileContent, err := r.ReadFile(filePath, cfg.FileChunkSize)
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

func createFilePath(wordLength int, cfg config.Config) string {
	rootDir := utils.FindProjectRoot()
	return filepath.Join(rootDir, fmt.Sprintf("%s%d.txt", cfg.BaseWordsPath, wordLength))
}
