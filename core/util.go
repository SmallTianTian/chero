package core

import (
	"os"
	"strings"
)

// IsDir 路径是否是文件夹.
func IsDir(path string) bool {
	fi, _ := os.Stat(path)
	if fi == nil {
		return false
	}
	return fi.IsDir()
}

// IsExist 路径是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsNotExist(err)
}

// StringUnderscored 字符串中驼峰命名转下划线分割。
func StringUnderscored(str string) string {
	size := len(str)
	sb := strings.Builder{}
	for i := 0; i < size; i++ {
		var b byte
		if b = str[i]; b >= 'A' && b <= 'Z' {
			b += 'a' - 'A'
		}
		sb.WriteByte(b)
		if b != '_' && i+1 < size && str[i+1] >= 'A' && str[i+1] <= 'Z' {
			sb.WriteByte('_')
		}
	}
	return sb.String()
}

// UrlUnderscored 网址中驼峰命名转下划线分割。
func UrlUnderscored(str string) string {
	strs := make([]string, 0)
	for _, v := range strings.Split(str, "/") {
		strs = append(strs, StringUnderscored(v))
	}
	return strings.Join(strs, "/")
}

func AbsHttpPath(dir, bDir string) string {
	sub := dir[len(bDir):]
	sub = strings.TrimSpace(sub)
	switch {
	case sub == "":
		fallthrough
	case sub == ".":
		fallthrough
	case sub[:2] == "./":
		return "/"
	case sub[0] == '/':
		return sub
	default:
		return "/" + sub
	}
}
