package main

import (
	"reflect"
	"testing"
)

func Test_readFDInfo(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readFDInfo(tt.args.fileName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readFDInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkFlags(t *testing.T) {
	type args struct {
		hex int64
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkFlags(tt.args.hex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
