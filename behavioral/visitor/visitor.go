package visitor

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
)

type FileInterface interface {
	accept(visitor Visitor)
}

type JsonFile struct {
	msg string
}

func (s *JsonFile) accept(visitor Visitor) {
	visitor.visitorJsonFile(s)
}

type YamlFile struct {
	msg string
}

func (s *YamlFile) accept(visitor Visitor) {
	visitor.visitorYamlFile(s)
}

type Visitor interface {
	visitorJsonFile(s *JsonFile)
	visitorYamlFile(s *YamlFile)
}

func aesEncrypt(input string, key string) string {
	inputByte := []byte(input)
	keyByte := []byte(key)

	block, err := aes.NewCipher(keyByte)
	if err != nil {
		log.Fatal(err)
	}

	blockSize := block.BlockSize()
	inputByte = PKCS7Padding(inputByte, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, keyByte[:blockSize])

	encrypted := make([]byte, len(inputByte))
	blockMode.CryptBlocks(encrypted, inputByte)
	return base64.StdEncoding.EncodeToString(encrypted)
}

func PKCS7Padding(cipherText []byte, blockSize int) []byte {
	// 当待加密数据的长度刚好是密钥的长度的时候，要填充的数据最多。因此，无论待加密数据长度是什么，都会需要进行padding
	padding := blockSize - len(cipherText)%blockSize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, paddingText...)
}

type Encryptor struct{}

func (e *Encryptor) visitorJsonFile(s *JsonFile) {
	//json的加密使用密钥1234567812345678
	s.msg = aesEncrypt(s.msg, string("1234567812345678"))
}

func (e *Encryptor) visitorYamlFile(s *YamlFile) {
	//yaml的加密使用密钥8765432187654321
	s.msg = aesEncrypt(s.msg, string("8765432187654321"))
}

type Compressor struct{}

func (c *Compressor) visitorJsonFile(s *JsonFile) {
	// json使用gzip压缩
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	defer gz.Close()
	if _, err := gz.Write([]byte(s.msg)); err != nil {
		log.Fatal(err)
	}
	s.msg = b.String()
}

func (c *Compressor) visitorYamlFile(s *YamlFile) {
	// yaml使用flate压缩
	var b bytes.Buffer
	zw, err := flate.NewWriter(&b, flate.BestCompression)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := zw.Write([]byte(s.msg)); err != nil {
		log.Fatal(err)
	}
	s.msg = b.String()
}

func main() {

}
