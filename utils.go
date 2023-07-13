package main

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Generate merchantToken / Any SHA256 Tokens
func generateSHA256Token(builder string) string {
	// Create a new SHA256 hash object
	hash := sha256.New()

	// Write the input string to the hash object
	hash.Write([]byte(builder))

	// Get the computed hash sum
	hashSum := hash.Sum(nil)

	return hex.EncodeToString(hashSum)
}

// Buat generate random number berdasarkan range
func generateRandomNumber(maxLength int) string {
	// Generate a random integer within the maximum length
	randomInt := rand.Intn(pow10(maxLength))

	// Convert the random integer to a string
	randomString := strconv.Itoa(randomInt)

	return randomString
}

// Range number dari 0 sampai 10^max-1
func pow10(exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= 10
	}
	return result
}

// Untuk generator maxCcYyMm
func generate5Years() string {
	// currentYear := strconv.Itoa(time.Now().Year())
	// randomYear := currentYear + offset
	currentYear := time.Now().Year() % 100
	maxOffset := 5
	offset := rand.Intn(maxOffset + 1)

	randomYear := currentYear + offset
	randomMonth := rand.Intn(12) + 1

	yearStr := strconv.Itoa(randomYear)
	monthStr := strconv.Itoa(randomMonth)

	if len(yearStr) < 2 {
		yearStr = "0" + yearStr
	}
	if len(monthStr) < 2 {
		monthStr = "0" + monthStr
	}

	randomYearMonth := yearStr + monthStr

	return randomYearMonth
}

// Extraksi Form
func extractForm(htmlBody string) (url.Values, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	formValue := url.Values{}
	doc.Find("form input").Each(func(_ int, s *goquery.Selection) {

		name, _ := s.Attr("name")
		value, _ := s.Attr("value")
		formValue.Add(name, value)
	})

	return formValue, nil
}
