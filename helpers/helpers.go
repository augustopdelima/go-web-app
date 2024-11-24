package helpers

import (
	"net/mail"
	"net/url"
	"strings"
	"unicode"
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

func ValidateCellphoneNumber(cellphoneNumber string) bool {
	cellphoneNumber = strings.ReplaceAll(cellphoneNumber, " ", "")
	cellphoneNumber = strings.ReplaceAll(cellphoneNumber, "-", "")
	cellphoneNumber = strings.ReplaceAll(cellphoneNumber, "(", "")
	cellphoneNumber = strings.ReplaceAll(cellphoneNumber, ")", "")

	// Verifica se o número contém apenas dígitos
	for _, char := range cellphoneNumber {
		if !unicode.IsDigit(char) {
			return false
		}
	}

	// Verifica se o número tem tamanho válido (11 dígitos)
	length := len(cellphoneNumber)

	return length == 11
}
