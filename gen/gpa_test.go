package gen

import (
	"testing"
	"os"
	"path/filepath"

	"github.com/zhenzou/gpa/common"
)

var (
	gpa = NewDebugGpa(NewGenerator(&GormTransformer{}, &common.GpaParser{}))
)

func TestGpa_Process(t *testing.T) {
	gopath := os.Getenv("GOPATH")
	fp := filepath.Join(gopath, "src", "github.com", "zhenzou", "gpa", "example", "example.go")
	gpa.Process(fp)
}
