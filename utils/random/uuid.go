package random

import (
	uuid "github.com/satori/go.uuid"
	"strconv"
	"strings"
)

var char = []string{"a", "b", "c", "d", "e", "f",
	"g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s",
	"t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5",
	"6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I",
	"J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V",
	"W", "X", "Y", "Z"}

// ShortUuid 随机生成16位字符 字母+数字
func ShortUuid() string {
	var result string
	replace := Uuid()
	for i := 0; i < 8; i++ {
		s := replace[i*4 : i*4+4]
		parseInt, _ := strconv.ParseInt(s, 16, 64)
		result += char[parseInt%62]
	}
	replace2 := Uuid()
	for i := 0; i < 8; i++ {
		s := replace2[i*4 : i*4+4]
		parseInt, _ := strconv.ParseInt(s, 16, 64)
		result += char[parseInt%62]
	}
	return result
}

// Uuid 随机32字符
func Uuid() string {
	v4 := uuid.NewV4()
	return strings.Replace(v4.String(), "-", "", -1)
}
