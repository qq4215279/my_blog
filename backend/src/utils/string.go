package utils

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
)

// 默认分隔符
var defaultSeparator = ","

// 是否为空字符串
func IsBlank(str string) (IsBlank bool) {
	return str == string("") || strings.Trim(str, " ") == string("")
}

// 是否为空字符串
func IsNotBlank(str string) (IsNotBlank bool) {
	return !IsBlank(str)
}

// 是否有一个字符串是空串
func IsAnyBlank(strArr ...string) (IsAnyBlank bool) {
	for _, str := range strArr {
		if IsBlank(str) {
			return true
		}
	}
	return false
}

// 是否全部为空
func IsAllBlank(strArr ...string) (IsAllBlank bool) {
	for _, str := range strArr {
		if !IsBlank(str) {
			return false
		}
	}
	return true
}

// 按默认分隔符分割成[]string
func toArray(str string) (res []string) {
	return strings.Split(str, defaultSeparator)
}

// 按默认分隔符分割成[]int
func ToIntArray(str string) (res []int, error error) {
	if IsBlank(str) {
		return res, nil
	}
	stringArray := toArray(str)
	error = nil
	for _, intStr := range stringArray {
		intValue, err := strconv.Atoi(intStr)
		if err != nil {
			log.Printf("string 转 int 失败 %s", err.Error())
			error = err
		}
		res = append(res, intValue)
	}
	return res, error
}

// 按分隔符分割成[]int
func ToIntArrayBySeparator(str string, separator string) (res []int) {
	if IsBlank(str) {
		return res
	}
	stringArray := strings.Split(str, separator)
	for _, intStr := range stringArray {
		intValue, err := strconv.Atoi(intStr)
		if err != nil {
			panic(err)
		}
		res = append(res, intValue)
	}
	return res
}

func IntArray2String(arr []int, separator string) (res string) {
	res = ""
	for _, v := range arr {
		if len(res) > 0 {
			res = res + separator
		}
		res = res + strconv.Itoa(v)
	}
	return res
}

// 字符串转成map
func Bytes2Map(bytes []byte) (res map[string]interface{}) {
	res = make(map[string]interface{})
	_ = json.Unmarshal(bytes, &res)
	return res
}

// string转map
func String2Map(str string) (res map[string]interface{}) {
	res = make(map[string]interface{})
	_ = json.Unmarshal([]byte(str), &res)
	return Bytes2Map([]byte(str))
}

// map转string
func Map2String(str map[string]interface{}) (res string) {
	bytes, err := json.Marshal(&str)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
