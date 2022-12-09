package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"io"
)


func MD5HashFromString(input string) []byte {
	MD5 := md5.New()
	MD5.Write([]byte(input))
	return MD5.Sum(nil)
}

func AESEncryptFromReader(source io.Reader, key string) ([]byte, error) {
	data, err := io.ReadAll(source)
	if err != nil {
		return nil, err;
	}

	return AESEncrypt(data,key)
}

func AESEncrypt(source []byte, key string) ([]byte, error) {
	// create cyper
	blockCipher, err := aes.NewCipher(MD5HashFromString(key))

	if err != nil {
		return nil, err;
	}

	// new cipher 
	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err;
	}

	// create a nonce
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader,nonce)
	if err != nil {
		return nil, err;
	}
	// encrypt with starting of the cipher text containing the nonce
	return gcm.Seal(nonce,nonce,source,nil), nil
}

func AESDecryptFromReader(source io.Reader, key string) ([]byte, error) {
	data, err := io.ReadAll(source)
	if err != nil {
		return nil, err;
	}

	return AESDecrypt(data,key)
}

func AESDecrypt(cipherText []byte, key string) ([]byte, error) {
	// create cyper
	blockCipher, err := aes.NewCipher(MD5HashFromString(key))

	if err != nil {
		return nil, err;
	}

	// new cipher 
	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err;
	}

	noncesize := gcm.NonceSize()

	if len(cipherText) < noncesize {
		return nil, CIPHER_TEXT_IMPROPER_SIZE
	}

	nonce , cipherData := cipherText[:noncesize], cipherText[noncesize:]

	return gcm.Open(nil, nonce,cipherData, nil)
}