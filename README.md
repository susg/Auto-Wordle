![Go Version](https://img.shields.io/badge/go-1.22+-blue)
# 🟩🟨⬜️ Auto Wordle

Auto Wordle is a command-line application that helps you solve Wordle puzzles efficiently. Given your guesses and their feedback (green, yellow, black), Auto Wordle filters and suggests possible solutions.

## Features

- Interactive command-line interface
- Supports multiple word lengths (currently 5-letter words)
- Fast filtering of possible words based on your input
- Designed for speed and ease of use

## Usage

1. **Run the application:**
   ```sh
   go run main.go
   ```

2. **Follow the prompts:**

1) Enter the word length (e.g., 5).
2) For each letter in your guess, enter a two-character string:
    - The first character is the letter.
    - The second character is the color feedback:
        - g for green (correct letter, correct position)
        - y for yellow (correct letter, wrong position)
        - b for black/gray (letter not in the word)
    - Example input for a 5-letter word (apple): ag pb pb ly eb
3) View filtered word suggestions.
4) Enter q at any prompt to quit or start a new game.

## Example

```
*******************************************
*************** Auto Wordle ***************
*******************************************

Enter the word length (enter 'q' to quit): 5

Enter the input string (enter 'q' to start new game): ag pb pb ly eb

Filtered words: [along altar allow ...]
```

## Project Structure

- [`main.go`](main.go): Entry point and CLI logic
- [`internal/words`](internal/words): Loads and manages word lists
- [`internal/filter`](internal/filter): Filters words based on rules
- [`internal/validate`](internal/validate): Validates user input
- [`internal/reader`](internal/reader): Efficient file reading
- [`data/prod/5.txt`](data/prod/5.txt): Word list for 5-letter words

## Requirements
- Go 1.22+
- Word lists in [`data/prod`](data/prod)

## License
MIT License

---
Contributions and suggestions are welcome!
