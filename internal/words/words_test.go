package words

import (
	"errors"
	"reflect"
	"testing"

	"github.com/susg/autowordle/internal/config"
	"github.com/susg/autowordle/internal/reader"
	"github.com/susg/autowordle/utils"

	mock_reader "github.com/susg/autowordle/internal/reader/mock"
	"go.uber.org/mock/gomock"
)

func TestStartWordManager(t *testing.T) {
	cfg := config.GetConfig()
	cfg.SupportedWordLengths = []int{5, 6}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_reader.NewMockReader(ctrl)
	path1 := utils.FindProjectRoot() + "/data/test/5.txt"
	path2 := utils.FindProjectRoot() + "/data/test/6.txt"
	m.EXPECT().ReadFile(path1, 1024).Return("hello\nworld", nil)
	m.EXPECT().ReadFile(path2, 1024).Return("lively\nstring", nil)

	m2 := mock_reader.NewMockReader(ctrl)
	m2.EXPECT().ReadFile(gomock.Any(), gomock.Any()).Return("", errors.New("file not found"))
	type args struct {
		r reader.Reader
	}
	tests := []struct {
		name string
		args args
		want WordManager
	}{
		{
			name: "success",
			args: args{
				r: m,
			},
			want: &WordManagerImpl{
				wordsCache: map[int][]string{
					5: {"hello", "world"},
					6: {"lively", "string"},
				},
			},
		},
		{
			name: "failure",
			args: args{
				r: m2,
			},
			want: &WordManagerImpl{
				wordsCache: map[int][]string{
					5: {"hello", "world"},
					6: {"lively", "string"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil && tt.name == "failure" {
					t.Errorf("StartWordManager() did not panic")
				} else if r != nil && tt.name == "success" {
					t.Errorf("StartWordManager() panicked: %v", r)
				}
			}()
			if got := StartWordManager(tt.args.r, cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StartWordManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWordManagerImpl_GetWords(t *testing.T) {
	type fields struct {
		wordsCache map[int][]string
	}
	type args struct {
		wordLength int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				wordsCache: map[int][]string{5: {"hello", "world"}},
			},
			args:    args{wordLength: 5},
			want:    []string{"hello", "world"},
			wantErr: false,
		},
		{
			name: "failure",
			fields: fields{
				wordsCache: map[int][]string{},
			},
			args:    args{wordLength: 5},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wmi := &WordManagerImpl{
				wordsCache: tt.fields.wordsCache,
			}
			got, err := wmi.GetWords(tt.args.wordLength)
			if (err != nil) != tt.wantErr {
				t.Errorf("WordManagerImpl.GetWords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WordManagerImpl.GetWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createFilePath(t *testing.T) {
	type args struct {
		wordLength int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{wordLength: 5},
			want: utils.FindProjectRoot() + "/data/test/5.txt",
		},
	}

	cfg := config.GetConfig()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createFilePath(tt.args.wordLength, cfg); got != tt.want {
				t.Errorf("createFilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
