package shortener

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestShorten(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		wantKey string
	}{
		{
			name:    "empty value",
			args:    args{value: ""},
			wantKey: "short1",
		},
		{
			name:    "non-empty value",
			args:    args{value: "https://www.google.com"},
			wantKey: "short2",
		},
		{
			name:    "another non-empty value",
			args:    args{value: "https://www.youtube.com"},
			wantKey: "short3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotKey := Shorten(tt.args.value); gotKey != tt.wantKey {
				t.Errorf("Shorten() = %v, want %v", gotKey, tt.wantKey)
			} else {
				t.Logf("Value = %v, Shorten() = %v, want %v", tt.args.value, gotKey, tt.wantKey)
			}
		})
	}
}
