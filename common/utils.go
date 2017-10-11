package common

import (
	"os"
	"strings"
	"fmt"
	"bytes"
	"unicode"
	"math"
)

func IsGoFile(f os.FileInfo) bool {
	name := f.Name()
	return !f.IsDir() && !strings.HasPrefix(name, ".") && strings.HasSuffix(name, ".go")
}

// 忽略大小写的一系列方法
// some utils func for strings without case sensitive
func HasPrefix(s, prefix string) bool {
	s = strings.ToLower(s)
	prefix = strings.ToLower(prefix)
	return strings.HasPrefix(s, prefix)
}

func TrimPrefix(s, prefix string) string {
	s = strings.ToLower(s)
	prefix = strings.ToLower(prefix)
	return strings.TrimPrefix(s, prefix)
}

func IsInStringSlice(strs []string, s string) bool {
	for _, str := range strs {
		if str == s {
			return true
		}
	}
	return false
}

func AssertIn(strs []string, s string) {
	if !IsInStringSlice(strs, s) {
		panic(fmt.Sprintf("%s is not in %v", s, strs))
	}
}

func AssertPrefix(str string, prefix string) {
	if !strings.HasPrefix(str, prefix) {
		panic(fmt.Sprintf("%s must start with %s", str, prefix))
	}
}

func Concat(strings ...[]string) []string {
	ss := []string{}
	for _, str := range strings {
		ss = append(ss, str...)
	}
	return ss
}

func AnyPrefix(s string, prefixs []string) (prefix string, ok bool) {
	for _, prefix = range prefixs {
		if strings.HasPrefix(s, prefix) {
			ok = true
			return
		}
	}
	return
}

// TestModel->test_model
func TableName(model string) string {
	return toSnake(model)
}

func toSnake(s string) string {
	buf := bytes.NewBufferString("")
	for i, v := range s {
		if i > 0 && v >= 'A' && v <= 'Z' {
			buf.WriteRune('_')
		}
		buf.WriteRune(v)
	}
	return strings.ToLower(buf.String())
}

func MapIndexed(s string, f func(int, rune) bool) string {
	for i, r := range s {
		if !f(i, r) {
			break
		}
	}
	return s
}

// 如果以大写开头，则将首字母改为小写，如果以小写字幕开头则取前三个字幕
// TypeName->typeName
// error->err
func VarName(typeName string, plural bool) string {
	bs := make([]byte, 0, len(typeName))
	buf := bytes.NewBuffer(bs)
	for i, r := range typeName {
		if i == 0 {
			if unicode.IsLower(r) {
				buf.Write([]byte(typeName[:int(math.Min(float64(len(typeName)), 3))]))
				break
			}
			buf.WriteRune(unicode.ToLower(r))
		} else {
			buf.WriteRune(r)
		}
	}
	if plural {
		buf.WriteRune('s')
	}
	return buf.String()
}
