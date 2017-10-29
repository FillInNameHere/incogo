package datatype

import (
	"regexp"
	"errors"
)

type Email struct {
	ID uint
	Email string
}

func NewEmail(mail string) *Email {
	if !isValidMail(mail) {
		errors.New("invalid mail format")
		return nil
	}
	return &Email{Email:mail}
}

func isValidMail(email string) bool {
	var rx = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return rx.MatchString(email)
}