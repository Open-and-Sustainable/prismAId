package review

import (
	"reflect"
	"testing"
)

func TestNewQuery(t *testing.T) {
	tests := []struct {
		name    string
		prompts []string
		keys    []string
		want    Query
		wantErr bool
	}{
		{
			name:    "Valid inputs",
			prompts: []string{"What is AI?", "Describe machine learning."},
			keys:    []string{"AI", "ML"},
			want: Query{
				Prompts: []string{"What is AI?", "Describe machine learning."},
				Keys:    []string{"AI", "ML"},
			},
			wantErr: false,
		},
		{
			name:    "Empty inputs",
			prompts: []string{},
			keys:    []string{},
			want: Query{
				Prompts: []string{},
				Keys:    []string{},
			},
			wantErr: false,
		},
		{
			name:    "Mismatched inputs length",
			prompts: []string{"What is AI?"},
			keys:    []string{"AI", "ML"},
			want: Query{
				Prompts: []string{"What is AI?"},
				Keys:    []string{"AI", "ML"},
			},
			wantErr: false,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewQuery(tt.prompts, tt.keys)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQuery() got = %v, want %v", got, tt.want)
			}
		})
	}
}
