package words

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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
	rootDir := findProjectRoot()
	return filepath.Join(rootDir, fmt.Sprintf("data/prod/%d.txt", wordLength))
}

func findProjectRoot() string {
	// Start with the current file's directory
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	// Navigate up until you find go.mod
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			fmt.Printf("Found go.mod in directory: %s\n", dir)
			return dir
		}

		// Move up one directory
		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached filesystem root without finding go.mod
			// As a fallback, use the current working directory
			cwd, _ := os.Getwd()
			return cwd
		}
		dir = parent
	}
}
