package util

import (
	"reflect"
	"testing"
)

func TestTableName(t *testing.T) {
	type args struct {
		model string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"TableName1", args{"TableName"}, "table_name"},
		{"TableName2", args{"tableName"}, "table_name"},
		{"TableName3", args{"TableNameTable"}, "table_name_table"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TableName(tt.args.model); got != tt.want {
				t.Errorf("TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVarName(t *testing.T) {
	type args struct {
		typeName string
		plural   bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"VarName1", args{"TypeName", false}, "typeName"},
		{"VarName1", args{"TypeName", true}, "typeNames"},
		{"VarName1", args{"typeName", false}, "typ"},
		{"VarName1", args{"error", false}, "err"},
		{"VarName1", args{"er", false}, "er"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VarName(tt.args.typeName, tt.args.plural); got != tt.want {
				t.Errorf("VarName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcat(t *testing.T) {
	type args struct {
		strings [][]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Concat1", args{[][]string{{"123", "456"}, {"789"}}}, []string{"123", "456", "789"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Concat(tt.args.strings...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Concat() = %v, want %v", got, tt.want)
			}
		})
	}
}
