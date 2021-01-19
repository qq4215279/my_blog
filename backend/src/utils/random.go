/**
 * @Author xieed
 * @Description //随机数
 * @Date 2020/10/10 16:15
 **/
package utils

import (
	"math/rand"
	"strconv"
)

var (
	// 大小写字母和数字0-9
	randomStrArray = [62]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "a", "b",
		"c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w",
		"x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R",
		"S", "T", "U", "V", "W", "X", "Y", "Z"}
)

// 获取n位数字随机数
func getRandom(n int) (result string) {
	for i := 0; i < n; i++ {
		result = result + strconv.Itoa(rand.Intn(10))
	}
	return result
}

// 获取密码md5salt
func GetPasswordSalt() (result string) {
	return getRandom(6)
}

// 获取n位随机字符串
func GetRandomString(n int) (result string) {
	for i := 0; i < n; i++ {
		result = result + randomStrArray[rand.Intn(len(randomStrArray))]
	}
	return result
}

// 按权重随机下标
func GetRandomIndex(weightArray []float64) int {
	total := 0.0
	for _, weight := range weightArray {
		total += weight
	}

	n := rand.Float64() * total
	v := 0.0
	for index, w := range weightArray {
		v += w
		if n < v {
			return index
		}
	}
	return -1
}
