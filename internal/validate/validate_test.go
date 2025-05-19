package validate

import "testing"

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
			args:    args{input: []string{"aga", "gg", "by", "bb"}},
			wantErr: true,
		},
		{
			name:    "invalid letter",
			fields:  fields{wordLength: 5},
			args:    args{input: []string{"1g", "gg", "by", "bb"}},
			wantErr: true,
		},
		{
			name:    "invalid colour",
			fields:  fields{wordLength: 5},
			args:    args{input: []string{"aa", "gg", "by", "bb"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wv := &WordleValidator{
				wordLength: tt.fields.wordLength,
			}
			if err := wv.Validate(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("WordleValidator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
