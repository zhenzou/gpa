package gen

import "testing"

var (
	gpa = NewDebugGpa(&Generator{})
)

func TestGpa_Process(t *testing.T) {
	gpa.Process("/media/Media/Projects/Go/GOPATH/src/gpa/gen/gpa_test.go")
}
