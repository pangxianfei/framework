package utils

import (
	"math/rand"
	"strings"
	"time"
)

func GetRandStr(n int) (randStr string) {
	// 默认去掉了容易混淆的字符oOLl和数字01，要添加请使用addChars参数
	chars := "ABCDEFGHIJKMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz23456789"
	charsLen := len(chars)
	if n > 10 {
		n = 10
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		randIndex := rand.Intn(charsLen)
		randStr += chars[randIndex : randIndex+1]
	}
	return randStr
}

// Substr字符串截取
func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

// Explode 将字符串按字符拆成数组
func Explode(delimiter, datastr string) (arr []string) {
	ret := strings.Split(datastr, delimiter)

	for _, item := range ret {
		if item != "" {
			arr = append(arr, item)
		}
	}

	return arr
}
