package main

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	password := "password"
	hash := sha256.Sum256([]byte(password))
	hash512 := sha512.Sum512([]byte(password))
	fmt.Println("Password:", password)
	fmt.Printf("SHA-256 Hash: %x\n", hash)
	fmt.Printf("SHA-512 Hash: %x\n", hash512)

	salt, err := generageSalt()
	if err != nil {
		fmt.Println("Error generating salt:", err)
		return
	}
	hash1 := hashPassword(password, salt)
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Printf("Salt: %s\n", saltStr)
	fmt.Println("Hashed Password with Salt:", hash1)
}

func generageSalt() ([]byte, error) {
	salt := make([]byte, 16) // Generate a 16-byte salt
	_, error := io.ReadFull(rand.Reader, salt)
	if error != nil {
		return nil, fmt.Errorf("failed to generate salt: %v", error)
	}
	return salt, nil
}

func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}