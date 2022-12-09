package encryption_test

import (
	"ai-trustframework/pkg/encryption"
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestEncryptionAESByte(t *testing.T) {

	tt := []struct{
		Name string
		Data string
		Key string
	}{
		{
			Name: "base",
			Data: "This cannot be known by others",
			Key:  "secret",
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name,func(t *testing.T) {
			// encryption
			cipher, err := encryption.AESEncrypt([]byte(tc.Data),tc.Key)
			if err != nil {
				t.Logf("Unable to encrypt : %v",err)
				t.Fail()
			}

			assert.Nil(t, err)
			// decryption
			plain, err := encryption.AESDecrypt(cipher,tc.Key)

			if err != nil {
				t.Logf("Unable to decrypt : %v",err)
				t.Fail()
			}

			if string(plain) != tc.Data {
				t.Logf("Unable to decrypt to same string original: %s decrypted: %s",tc.Data,string(plain))
				t.Fail()
			}
		})
	}

}