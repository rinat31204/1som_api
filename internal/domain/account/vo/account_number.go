package vo

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"regexp"
	"strings"
)

const (
	letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits  = "0123456789"
)

var regex = regexp.MustCompile(`^[A-Z][0-9]{2}-[A-Z][0-9]{2}$`)

func randomChar(source string) (byte, error) {
	max := big.NewInt(int64(len(source)))
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}
	return source[n.Int64()], nil
}

func generateBlock() (string, error) {
	l, err := randomChar(letters)
	if err != nil {
		return "", err
	}

	d1, err := randomChar(digits)
	if err != nil {
		return "", err
	}

	d2, err := randomChar(digits)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%c%c%c", l, d1, d2), nil
}

func GenerateAccountNumber() (string, error) {
	left, err := generateBlock()
	if err != nil {
		return "", err
	}

	right, err := generateBlock()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s-%s", left, right), nil
}

func NormalizeAccountNumber(input string) (string, error) {
	accountNumber := strings.TrimSpace(input)
	accountNumber = strings.ToUpper(accountNumber)
	accountNumber = strings.ReplaceAll(accountNumber, " ", "")
	accountNumber = strings.ReplaceAll(accountNumber, "–", "-")
	accountNumber = strings.ReplaceAll(accountNumber, "—", "-")

	if !regex.MatchString(accountNumber) {
		return "", fmt.Errorf("invalid account number format")
	}

	return accountNumber, nil
}
