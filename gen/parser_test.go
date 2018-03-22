package gen

import (
	"reflect"
	"testing"
)

func TestGpaParser_trimPrefix(t *testing.T) {

	type args struct {
		fullName string
		prefix   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"trimPrefix1", args{"findAll", "FindFunc"}, "All", false},
		{"trimPrefix2", args{"FindAll", "FindFunc"}, "All", false},
		{"trimPrefix3", args{"findAll", "Find1"}, "", true},
		{"trimPrefix4", args{"finAll", "FindFunc"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GpaParser{}
			got, err := g.trimPrefix(tt.args.fullName, tt.args.prefix)
			if (err != nil) != tt.wantErr {
				t.Errorf("GpaParser.trimPrefix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GpaParser.trimPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

type (
	Model struct {
		Id       string
		Name     string
		LastName string
	}
)

var (
	func1 = &Func{
		FileName: "parser_test.go",
		FullName: "DeleteById",
		Params:   []*Field{{Name: "id", IsPointer: false, Typ: Type{TypeName: "string", IsPointer: false}}},
		Results:  []*Field{{Name: "model", IsPointer: false, Typ: Type{TypeName: "Model", IsPointer: true}}},
		Receiver: &Field{Name: "m", IsPointer: false, Typ: Type{TypeName: "Model", IsPointer: true}},
	}
	func2 = &Func{
		FileName: "parser_test.go",
		FullName: "DeleteByName",
		Params:   []*Field{{Name: "name", IsPointer: false, Typ: Type{TypeName: "string", IsPointer: false}}},
		Results:  []*Field{{Name: "model", IsPointer: false, Typ: Type{TypeName: "Model", IsPointer: true}}},
		Receiver: &Field{Name: "m", IsPointer: false, Typ: Type{TypeName: "Model", IsPointer: true}},
	}
	func3 = &Func{
		FileName: "parser_test.go",
		FullName: "DeleteByNameAndLastName",
		Params: []*Field{
			{Name: "name", IsPointer: false, Typ: Type{TypeName: "string", IsPointer: false}},
			{Name: "lastName", IsPointer: false, Typ: Type{TypeName: "string", IsPointer: false}},
		},
		Results:  []*Field{{Name: "model", IsPointer: false, Typ: Type{TypeName: "Model", IsPointer: true}}},
		Receiver: &Field{Name: "m", IsPointer: false, Typ: Type{TypeName: "Model", IsPointer: true}},
	}

	predicates1 = []*Predicate{
		&Predicate{
			Field: "Id", OpCode: OpEqual, OpText: "", Logic: "", ParamCount: 1,
		},
	}
	predicates2 = []*Predicate{
		&Predicate{
			Field: "Name", OpCode: OpEqual, OpText: "", Logic: "", ParamCount: 1,
		},
	}
	predicates3 = []*Predicate{
		&Predicate{
			Field: "Name", OpCode: OpEqual, OpText: "", Logic: "", ParamCount: 1,
		},
		&Predicate{
			Field: "LastName", OpCode: OpEqual, OpText: "", Logic: "And", ParamCount: 1,
		},
	}
)

func TestGpaParser_ParseCreate(t *testing.T) {

	type args struct {
		fd *Func
	}
	tests := []struct {
		name       string
		args       args
		wantCreate *CreateFunc
		wantErr    bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GpaParser{}
			gotCreate, err := g.ParseCreate(tt.args.fd)
			if (err != nil) != tt.wantErr {
				t.Errorf("GpaParser.ParseCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCreate, tt.wantCreate) {
				t.Errorf("GpaParser.ParseCreate() = %v, want %v", gotCreate, tt.wantCreate)
			}
		})
	}
}

func TestGpaParser_ParseDelete(t *testing.T) {

	type args struct {
		fd *Func
	}
	tests := []struct {
		name       string
		args       args
		wantDelete *DeleteFunc
		wantErr    bool
	}{
		{"ParseDelete1", args{fd: func1}, &DeleteFunc{Table: "Model", Func: func1, Predicates: predicates1}, false},
		{"ParseDelete2", args{fd: func2}, &DeleteFunc{Table: "Model", Func: func2, Predicates: predicates2}, false},
		{"ParseDelete3", args{fd: func3}, &DeleteFunc{Table: "Model", Func: func3, Predicates: predicates3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GpaParser{}
			gotDelete, err := g.ParseDelete(tt.args.fd)
			if (err != nil) != tt.wantErr {
				t.Errorf("GpaParser.ParseDelete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDelete, tt.wantDelete) {
				t.Errorf("GpaParser.ParseDelete() = %v, want %v", gotDelete, tt.wantDelete)
			}
		})
	}
}

func TestGpaParser_extractPredicate(t *testing.T) {

	type args struct {
		str string
	}
	tests := []struct {
		name           string
		args           args
		wantPredicates []*Predicate
		wantErr        bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GpaParser{}
			gotPredicates, _, err := g.extractPredicate(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("GpaParser.extractPredicate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPredicates, tt.wantPredicates) {
				t.Errorf("GpaParser.extractPredicate() = %v, want %v", gotPredicates, tt.wantPredicates)
			}
		})
	}
}

func TestGpaParser_extractTitle(t *testing.T) {

	type args struct {
		str string
	}
	tests := []struct {
		name      string
		args      args
		wantTitle string
		wantErr   bool
	}{
		{"extractTitle1", args{str: "TitleTitle"}, "Title", false},
		{"extractTitle2", args{str: "titleTitle"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GpaParser{}
			gotTitle, err := g.extractTitle(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("GpaParser.extractTitle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTitle != tt.wantTitle {
				t.Errorf("GpaParser.extractTitle() = %v, want %v", gotTitle, tt.wantTitle)
			}
		})
	}
}

func TestGpaParser_ParseUpdate(t *testing.T) {

	type args struct {
		fd *Func
	}
	tests := []struct {
		name       string
		args       args
		wantUpdate *UpdateFunc
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GpaParser{}
			gotUpdate, err := g.ParseUpdate(tt.args.fd)
			if (err != nil) != tt.wantErr {
				t.Errorf("GpaParser.ParseUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUpdate, tt.wantUpdate) {
				t.Errorf("GpaParser.ParseUpdate() = %v, want %v", gotUpdate, tt.wantUpdate)
			}
		})
	}
}

func TestGpaParser_ParseFind(t *testing.T) {

	type args struct {
		fd *Func
	}
	tests := []struct {
		name     string
		args     args
		wantFind *FindFunc
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GpaParser{}
			gotFind, err := g.ParseFind(tt.args.fd)
			if (err != nil) != tt.wantErr {
				t.Errorf("GpaParser.ParseFind() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFind, tt.wantFind) {
				t.Errorf("GpaParser.ParseFind() = %v, want %v", gotFind, tt.wantFind)
			}
		})
	}
}
