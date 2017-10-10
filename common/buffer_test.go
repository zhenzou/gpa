package common

import (
	"testing"
)

func TestOutputBuffer_WriteStringAt(t *testing.T) {
	type args struct {
		str string
		off int
	}
	tests := []struct {
		name       string
		args       args
		wantString string
	}{
		{"WriteStringAt1", args{"_string", 2}, "st_stringring"},
		{"WriteStringAt2", args{"_string", 4}, "stri_stringng"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := NewOutputBuffer([]byte("string"))
			o.WriteStringAt(tt.args.str, tt.args.off)

			if o.String() != tt.wantString {
				t.Errorf("OutputBuffer.WriteStringAt() = %v, want %v", o.String(), tt.wantString)
			}
		})
	}
}

func TestOutputBuffer_WriteStringAt1(t *testing.T) {
	o := NewOutputBuffer([]byte("string"))
	type args struct {
		str string
		off int
	}
	tests := []struct {
		name       string
		args       args
		wantString string
	}{
		{"WriteStringAt1", args{"_string", 2}, "st_stringring"},
		{"WriteStringAt2", args{"_string", 4}, "st_s_stringtringring"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o.WriteStringAt(tt.args.str, tt.args.off)

			if o.String() != tt.wantString {
				t.Errorf("OutputBuffer.WriteStringAt() = %v, want %v", o.String(), tt.wantString)
			}
		})
	}
}
