package reader

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type Reader interface {
	ReadFile(filePath string, chunkSize int) (string, error)
}

type FileReader struct{}

func NewFileReader() Reader {
	return &FileReader{}
}

func (fr *FileReader) ReadFile(filePath string, chunkSize int) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return "", fmt.Errorf("failed to get file info: %w", err)
	}
	fileSize := fileInfo.Size()

	numChunks := int(fileSize) / chunkSize
	if fileSize%int64(chunkSize) != 0 {
		numChunks++
	}

	// Create a slice to store chunks in the correct order
	chunks := make([]string, numChunks)
	var wg sync.WaitGroup

	for i := range numChunks {
		wg.Add(1)
		go func(chunkIndex int) {
			defer wg.Done()

			// Create a new file handle for each goroutine
			chunkFile, err := os.Open(filePath)
			if err != nil {
				fmt.Printf("Error opening file for chunk %d: %v\n", chunkIndex, err)
				return
			}
			defer chunkFile.Close()

			offset := int64(chunkIndex) * int64(chunkSize)
			chunkData, err := fr.readChunk(chunkFile, offset, chunkSize)
			if err != nil {
				fmt.Printf("Error reading chunk %d: %v\n", chunkIndex, err)
				return
			}

			// Store in the correct position
			chunks[chunkIndex] = chunkData
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Combine chunks in the correct order
	var output strings.Builder
	for _, chunk := range chunks {
		output.WriteString(chunk)
	}

	return output.String(), nil
}

func (fr *FileReader) readChunk(file *os.File, offset int64, chunkSize int) (string, error) {
	buffer := make([]byte, chunkSize)

	_, err := file.Seek(offset, io.SeekStart)
	if err != nil {
		return "", fmt.Errorf("failed to seek to offset %d: %w", offset, err)
	}

	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return "", fmt.Errorf("failed to read chunk at offset %d: %w", offset, err)
	}

	return string(bytes.Trim(buffer[:n], "\x00")), nil
}
