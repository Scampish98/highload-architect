package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"highload-architect/internal/apperror"
)

type Password string

func (p Password) Hash() (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		if errors.Is(err, bcrypt.ErrPasswordTooLong) {
			return "", apperror.ErrIncorrectPassword.WithDetails("password too long")
		}

		return "", err
	}

	return string(h), nil
}

func (p Password) Compare(hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(p))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return apperror.ErrIncorrectPassword
		}

		return err
	}

	return nil
}
