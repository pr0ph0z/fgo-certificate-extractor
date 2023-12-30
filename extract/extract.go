package extract

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"fmt"
)

var (
	key = "b5nHjsMrqaeNliSs3jyOzgpD"
	iv  = []byte("wuD6keVr")
)

func Extract(certificate string) string {
	decryptedCertificate, err := decryptDES(certificate)
	if err != nil {
		panic(err)
	}

	return decryptedCertificate
}

func decryptDES(certificate string) (string, error) {
	// Decode base64-encoded ciphertext
	ciphertext, err := base64.StdEncoding.DecodeString(certificate)
	if err != nil {
		return "", err
	}

	// Create a TripleDES block cipher
	block, err := des.NewTripleDESCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// Check if the ciphertext length is a multiple of the block size
	if len(ciphertext)%block.BlockSize() != 0 {
		return "", fmt.Errorf("ciphertext is not a multiple of the block size")
	}

	// Create a mode with the given IV
	mode := cipher.NewCBCDecrypter(block, iv)

	// Decrypt the ciphertext
	mode.CryptBlocks(ciphertext, ciphertext)

	// Remove PKCS7 padding (assuming PKCS7 padding is used)
	padLength := int(ciphertext[len(ciphertext)-1])
	ciphertext = ciphertext[:len(ciphertext)-padLength]

	return string(ciphertext), nil
}
