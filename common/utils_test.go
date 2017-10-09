package common

import "testing"

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
