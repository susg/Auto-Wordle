package validate

import (
	"testing"

	"github.com/susg/autowordle/internal/config"
)

func TestWordleValidator_Validate(t *testing.T) {
	type fields struct {
		wordLength int
	}
	type args struct {
		input []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "valid input",
			fields:  fields{wordLength: 5},
			args:    args{input: []string{"ag", "gg", "by", "bb", "yy"}},
			wantErr: false,
		},
		{
			name:    "invalid input length",
			fields:  fields{wordLength: 5},
			args:    args{input: []string{"ag", "gg", "by", "bb"}},
			wantErr: true,
		},
		{
			name:    "invalid input",
			fields:  fields{wordLength: 5},
			args:    args{input: []string{"aga", "gg", "by", "bb", "yy"}},
			wantErr: true,
		},
		{
			name:    "invalid letter",
			fields:  fields{wordLength: 5},
			args:    args{input: []string{"1g", "gg", "by", "bb", "yy"}},
			wantErr: true,
		},
		{
			name:    "invalid colour",
			fields:  fields{wordLength: 5},
			args:    args{input: []string{"aa", "gg", "by", "bb", "yy"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wv := &WordleValidator{
				wordLength: tt.fields.wordLength,
				cfg:        config.GetConfig(),
			}
			if err := wv.Validate(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("WordleValidator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewWordleValidator(t *testing.T) {
	cfg := config.GetConfig()
	cfg.SupportedWordLengths = []int{4, 5, 6}

	tests := []struct {
		name       string
		wordLength int
		wantErr    bool
	}{
		{
			name:       "supported word length 5",
			wordLength: 5,
			wantErr:    false,
		},
		{
			name:       "supported word length 4",
			wordLength: 4,
			wantErr:    false,
		},
		{
			name:       "unsupported word length 7",
			wordLength: 7,
			wantErr:    true,
		},
		{
			name:       "unsupported word length 0",
			wordLength: 0,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator, err := NewWordleValidator(tt.wordLength, cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWordleValidator() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && validator == nil {
				t.Errorf("Expected non-nil validator for supported word length %d", tt.wordLength)
			}
			if tt.wantErr && validator != nil {
				t.Errorf("Expected nil validator for unsupported word length %d", tt.wordLength)
			}
		})
	}
}
