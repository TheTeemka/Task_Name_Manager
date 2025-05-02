package utils

import (
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plain string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}
	return hash, nil
}

func CompareHashAndPassword(plain, hash []byte) (bool, error) {
	log.Println(string(plain), string(hash))
	err := bcrypt.CompareHashAndPassword(hash, plain)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
