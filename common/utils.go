package common

import (
	"os"
	"strings"
	"fmt"
)

//func visitFile(path string, f os.FileInfo, err error) error {
//	if err == nil && IsGoFile(f) {
//		err = processFile(path, nil, os.Stdout, false)
//	}
//	// Don't complain if a file was deleted in the meantime (i.e.
//	// the directory changed concurrently while running gofmt).
//	if err != nil && !os.IsNotExist(err) {
//		panic(err)
//	}
//	return nil
//}

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
