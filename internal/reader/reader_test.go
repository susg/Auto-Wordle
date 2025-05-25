package reader

import (
	"os"
	"testing"
)

func TestNewFileReader(t *testing.T) {
	fr := NewFileReader()
	if _, ok := fr.(*FileReader); !ok {
		t.Errorf("NewFileReader() did not return *FileReader, got %T", fr)
	}
}

func TestFileReader_ReadFile(t *testing.T) {
	type args struct {
		filePath  string
		chunkSize int
	}
	tests := []struct {
		name    string
		fr      *FileReader
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "failed to open file",
			fr:   &FileReader{},
			args: args{
				filePath:  "data/test/6.txt",
				chunkSize: 10,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "successful file read",
			fr:   &FileReader{},
			args: args{
				filePath:  "../../data/test/5.txt",
				chunkSize: 10,
			},
			want:    getExpectedContent(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fr := &FileReader{}
			got, err := fr.ReadFile(tt.args.filePath, tt.args.chunkSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileReader.ReadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FileReader.ReadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileReader_readChunk(t *testing.T) {
	file, err := os.Open("../../data/test/5.txt")
	if err != nil {
		t.Fatalf("failed to open test file: %v", err)
	}
	defer file.Close()

	type args struct {
		offset    int64
		chunkSize int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "read first 5 bytes",
			args:    args{offset: 0, chunkSize: 5},
			want:    "which",
			wantErr: false,
		},
		{
			name:    "read middle chunk",
			args:    args{offset: 12, chunkSize: 5},
			want:    "their",
			wantErr: false,
		},
		{
			name:    "large chunksize",
			args:    args{offset: 90, chunkSize: 100},
			want:    "think",
			wantErr: false,
		},
		{
			name:    "negative offset",
			args:    args{offset: -1, chunkSize: 5},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fr := &FileReader{}
			got, err := fr.readChunk(file, tt.args.offset, tt.args.chunkSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("readChunk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("readChunk() = %q, want %q", got, tt.want)
			}
		})
	}
}

func getExpectedContent() string {
	content := `which
there
their
about
would
these
other
words
could
write
first
water
after
where
right
think`
	return content
}
