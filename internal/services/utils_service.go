package service

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/devstackq/real-time-forum/internal/models"
)

func Randomaizer() string {
	rand.Seed(time.Now().Unix())
	charSet := "abcdedfghijklmnopqrst"
	var output strings.Builder
	length := 10
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}
	return output.String()
}
func IsEmailValid(user *models.User) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,6}$`)
	return Re.MatchString(user.Email)
}

func IsPasswordValid(user *models.User) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(user.Password) >= 7 {
		hasMinLen = true
	}
	for _, char := range user.Password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}
