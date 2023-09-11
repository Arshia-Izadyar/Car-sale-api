package common

import (
	"unicode"

	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
)

func HasLetter(s string) bool {
	for _, c := range s {
		if unicode.IsLetter(c) {
			return true
		}
	}
	return false
}

func HasDigits(s string) bool {
	for _, c := range s {
		if unicode.IsDigit(c) {
			return true
		}
	}
	return false
}

func HasLower(s string) bool {
	for _, c := range s {
		if unicode.IsLower(c) && unicode.IsLetter(c) {
			return true
		}
	}
	return false
}

func HasUpper(s string) bool {
	for _, c := range s {
		if unicode.IsUpper(c) && unicode.IsLetter(c) {
			return true
		}
	}
	return false
}

func CheckPassword(pass string) bool {
	cfg := config.GetConfig()
	if len(pass) < cfg.Password.MinLength || len(pass) > cfg.Password.MaxLength {
		return false
	}
	if cfg.Password.IncludeChars && !HasLetter(pass) {
		return false
	}
	if cfg.Password.IncludeDigits && !HasDigits(pass) {
		return false
	}
	if cfg.Password.IncludeUppercase && !HasUpper(pass) {
		return false
	}
	if cfg.Password.IncludeLowercase && !HasLower(pass) {
		return false
	}

	return true
}
