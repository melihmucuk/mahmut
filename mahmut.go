package mahmut

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// Bagla encrypts the text with the given key.
func Bagla(key string, text string) string {

	k := len(key)
	switch k {
	default:
		panic("key length must be 16, 24 or 32")
	case 16, 24, 32:
		break
	}

	plaintext := []byte(text)

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext)
}

// Coz decrypts the cryptoText with the given key.
func Coz(key string, cryptoText string) string {
	k := len(key)
	switch k {
	default:
		panic("key length must be 16, 24 or 32")
	case 16, 24, 32:
		break
	}

	ciphertext, _ := base64.StdEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}
