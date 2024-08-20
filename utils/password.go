package utils

import (
	"log"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// Password must be 8 characters long, must include minimum one uppercase, one lowercase, one number and one special character
func IsStrongPassword(password string)bool{
    if len(password) < 8 {
        return false
    }

    hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
    hasLowercase := regexp.MustCompile(`[a-z]`).MatchString(password)
    hasDigit := regexp.MustCompile(`\d`).MatchString(password)
    hasSpecialChar := regexp.MustCompile(`[!"#$%&'()*+,-./:;<=>?@[\]^_{|}~]`).MatchString(password)

    return hasUppercase && hasLowercase && hasDigit && hasSpecialChar
}


func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}


func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = "login or password is incorrect"
		check = false
	}

	return check, msg
}