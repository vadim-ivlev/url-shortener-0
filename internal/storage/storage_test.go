package storage

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	Create()
	os.Exit(m.Run())
}

func TestSet(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty value",
			args: args{key: "", value: ""},
			want: "",
		},
		{
			name: "non-empty value",
			args: args{key: "short1", value: "https://www.google.com"},
			want: "short1",
		},
		{
			name: "another non-empty value",
			args: args{key: "short2", value: "https://www.youtube.com"},
			want: "short2",
		},
		{
			name: "another non-empty value",
			args: args{key: "short4", value: "https://www.youtube.com"},
			want: "short2",
		},
		{
			name: "empty key",
			args: args{key: "", value: "https://www.youtube.com"},
			want: "short2",
		},
		{
			name: "empty value",
			args: args{key: "short3", value: ""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Set(tt.args.key, tt.args.value)
			fmt.Printf("Name = %v, Key = %v, Value = %v, Want = %v, Got = %v\n", tt.name, tt.args.key, tt.args.value, tt.want, got)
			if got != tt.want {
				t.Errorf("Set() = %v, want %v", got, tt.want)
			}
		})
	}
}
