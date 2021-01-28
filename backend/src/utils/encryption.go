/**
 * @Author xieed
 * @Description //加解密
 * @Date 2020/10/10 16:08
 **/
package utils

import (
	"backend/src/constants"
	"backend/src/global"
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
)

// md5加密
func MD5Encryption(str string, salt string) (result string) {
	write := md5.New()
	_, _ = io.WriteString(write, constants.StaticSalt+str+salt)
	//w.Sum(nil)将w的hash转成[]byte格式
	return fmt.Sprintf("%x", write.Sum(nil))
}

// 密码加密
func EncryptPassword(password string) string {
	return DesCBCEncrypt(password, constants.PasswordKeyPrefix+DesCBCDecrypt(global.Config.Crypto.CipherTextSuffix, global.Config.Crypto.DecodeKeySuffix))
}

// 密码解密
func DecryptPassword(cipherText string) string {
	return DesCBCDecrypt(cipherText, constants.PasswordKeyPrefix+DesCBCDecrypt(global.Config.Crypto.CipherTextSuffix, global.Config.Crypto.DecodeKeySuffix))
}

// des加密
func DesCBCEncrypt(contentStr, keyStr string) string {
	content := []byte(contentStr)
	key := []byte(keyStr)
	//1：创建并返回一个DES算法的cipher.block接口
	block, err := des.NewCipher(key)
	//2：判断是否创建成功
	if err != nil {
		panic(err)
	}
	//3：对最后一个明文分组进行数据填充
	content = paddingPKCS5(content, block.BlockSize())
	//4：创建一个密码分组为链接模式的，底层使用DES加密的BLockMode接口
	//    参数iv的长度，必须等于b的块尺寸
	blockMode := cipher.NewCBCEncrypter(block, key)
	//5:加密连续的数据块
	dst := make([]byte, len(content))
	blockMode.CryptBlocks(dst, content)
	//fmt.Println("加密之后的数据:",dst)
	//6:将加密后的数据返回
	return base64.StdEncoding.EncodeToString(dst)
}

//des解密
func DesCBCDecrypt(cipherTextStr, keyStr string) string {
	cipherText, _ := base64.StdEncoding.DecodeString(cipherTextStr)
	key := []byte(keyStr)
	//1:创建并返回一个使用DES算法的cipher.block接口
	block, err := des.NewCipher(key)
	//2:判断是否创建成功
	if err != nil {
		panic(err)
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	//解密数据
	dst := make([]byte, len(cipherText))
	blockMode.CryptBlocks(dst, cipherText)
	//5:去掉最后一组填充数据
	dst = unPaddingPKCS5(dst)

	//返回结果
	return string(dst)
}

//使用pkcs5的方式填充

func paddingPKCS5(cipherText []byte, blockSize int) []byte {
	//1:计算最后一个分组缺多少字节
	padding := blockSize - (len(cipherText) % blockSize)
	//2:创建一个大小为padding的切片,每个字节的值为padding
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	//3:将padText添加到原始数据的后边，将最后一个分组缺少的字节数补齐
	newText := append(cipherText, padText...)

	return newText
}

//删除pkcs5填充的尾部数据
func unPaddingPKCS5(origData []byte) []byte {
	//1:计算数据的总长度
	length := len(origData)
	//2:根据填充的字节值得到填充的次数
	number := int(origData[length-1])
	//3:将尾部填充的number个字节去掉
	return origData[:(length - number)]
}
