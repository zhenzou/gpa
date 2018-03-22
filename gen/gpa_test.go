package gen

import (
	"testing"
	"os"
	"path/filepath"
)

var (
	gpa = NewDebugGpa(NewGenerator(&GormTransformer{}, &GpaParser{}))
)

func TestGpa_Process(t *testing.T) {
	gopath := os.Getenv("GOPATH")
	fp := filepath.Join(gopath, "src", "github.com", "zhenzou", "gpa", "example", "example.go")
	gpa.Process(fp)
}
