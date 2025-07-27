package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"io"

	"passm/internal/helper"
	"passm/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

func deriveTo32Bytes() string {
	mainPass, _ := repository.FinMainPasswordById(helper.Session.Values["user_id"].(int))

	// Step 2: Derive a 32-byte key using SHA-25. Returns a 32-byte array
	hash := sha256.Sum256([]byte(mainPass))
	// Step 3: Convert hash to hexadecimal string
	fullHexString := hex.EncodeToString(hash[:])
	// Step 4: Truncate or adjust to 32 characters. Use the first 32 characters
	aesKey := fullHexString[:32]

	return aesKey
}

func Encrypt(plainPassword string) string {
	secretKey := deriveTo32Bytes()

	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return ""
	}

	plainText := []byte(plainPassword)

	// The IV needs to be unique, but not secure. Therefore, it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plainText))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return ""
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plainText)

	return base64.RawStdEncoding.EncodeToString(ciphertext)
}

func Decrypt(encryptPassword string) string {
	secretKey := deriveTo32Bytes()

	ciphertext, err := base64.RawStdEncoding.DecodeString(encryptPassword)
	if err != nil {
		return ""
	}

	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return ""
	}

	// The IV needs to be unique, but not secure. Therefore, it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return ""
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes), err
}
