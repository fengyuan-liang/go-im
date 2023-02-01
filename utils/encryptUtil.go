package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"strings"
)

// @Description: md5工具类
// @Version: 1.0.0
// @Date: 2023/01/31 13:22
// @Author: fengyuan-liang@foxmail.com

//==================== 散列值加密 ==========================

// Check
//
//	@Description: 传入一个未加密的字符串和与加密后的数据，进行对比，如果正确就返回true
//	@args content 未加密的字符
//	@args encrypted 加密后的数据
//	@return bool 正确就返回true
func Check(content, encrypted string) bool {
	return strings.EqualFold(Encode(content), encrypted)
}

// Encode
//
//	@Description: 散列值加密
//	@args data 要加密的字符串
//	@return string 加密后的数据
func Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// EncodeBySalt md5并且加盐混淆
//
//	@args plainPwd 明文密码
//	@args salt 盐
func EncodeBySalt(plainPwd string, salt string) string {
	return Encode(plainPwd + salt)
}

// CheckBySalt 检查加盐后密码是否一样
func CheckBySalt(plainPwd string, salt string, password string) bool {
	return EncodeBySalt(plainPwd, salt) == password
}

//==================== 对称加密 ==========================

// AesEncrypt
//
//	@Description: aes对称加密
//	@args orig 要加密的字段
//	@args key 密钥
//	@return string 加密后的字段
func AesEncrypt(orig string, key string) string {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	return base64.StdEncoding.EncodeToString(cryted)

}

// AesDecrypt
//
//	@Description: aes解密
//	@args cryted 加密的密文
//	@args key 密钥
//	@return string 解密的字符串
func AesDecrypt(cryted string, key string) string {
	// 转成字节数组
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	k := []byte(key)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return string(orig)
}

// PKCS7Padding 补码
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS7UnPadding 去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
