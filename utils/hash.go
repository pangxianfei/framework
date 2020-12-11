package utils

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func Md5ToString(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}

func ShortTag(longstr string, number int) string {
	baseVal := 0x3FFFFFFF
	indexVal := 0x0000003D
	charset := []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

	if number < 1 || number > 4 {
		number = 1
	}
	key := "orange"
	urlhash := Md5ToString(key + longstr)
	len := len(urlhash)

	var hexcc int64
	var short_url []byte
	var result [4]string

	for i := 0; i < 4; i++ {
		urlhash_piece := Substr(urlhash, i*len/4, len/4)
		hexDec, _ := strconv.ParseInt(urlhash_piece, 16, 64)
		hexcc = hexDec & int64(baseVal)

		var index int64
		short_url = []byte{}
		for j := 0; j < 6; j++ {
			//将得到的值与0x0000003d,3d为61，即charset的坐标最大值
			index = hexcc & int64(indexVal)
			short_url = append(short_url, charset[index])
			//循环完以后将hex右移5位
			hexcc = hexcc >> 5
		}
		result[i] = string(short_url)
	}

	return result[number]

}
