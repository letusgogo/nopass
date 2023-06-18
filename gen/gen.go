package gen

import (
	"github.com/letusgogo/nopass/algo"
	"github.com/letusgogo/nopass/rule"
	"hash/fnv"
)

const (
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits           = "0123456789"
	specialSymbols   = "!@#$%^&*()-_+=<>?"
)

func hashToIndex(str string, charset string) int {
	hash := fnv.New32a()
	_, err := hash.Write([]byte(str))
	if err != nil {
		panic(err)
	}
	sum := hash.Sum32()
	return int(sum) % len(charset)
}

func generateHardPassword(password string) string {
	hardPassword := make([]byte, len(password))

	// Iterate through each character of the original password
	for i := 0; i < len(password); i++ {
		switch i % 4 {
		case 0: // Map to lowercase letters
			index := hashToIndex(password[i:], lowercaseLetters)
			hardPassword[i] = lowercaseLetters[index]
		case 1: // Map to uppercase letters
			index := hashToIndex(password[i:], uppercaseLetters)
			hardPassword[i] = uppercaseLetters[index]
		case 2: // Map to digits
			index := hashToIndex(password[i:], digits)
			hardPassword[i] = digits[index]
		case 3: // Map to special symbols
			index := hashToIndex(password[i:], specialSymbols)
			hardPassword[i] = specialSymbols[index]
		}
	}

	return string(hardPassword)
}

func GeneratePassword(rule *rule.Rule, algo algo.Algorithm) string {
	originalPassword := rule.Compose()

	password := algo.Generate(originalPassword)

	hardPassword := generateHardPassword(password)
	return hardPassword
}
