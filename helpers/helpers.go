package helpers

import (
	"net/mail"
	"net/url"
)

func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func ValidateUrl(address string) bool {
	url, err := url.Parse(address)

	if url.Scheme == "" {
		return false
	}

	if url.Host == "" {
		return false
	}

	return err == nil
}
