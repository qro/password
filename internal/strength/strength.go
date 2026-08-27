package strength

import (
	"crypto/sha1"
	"fmt"
	"io"
	"math"
	"net/http"
	"strings"
	"unicode"
)

type Result struct {
	Entropy  float64
	Rating   string
	Breached bool
}

func Check(password string) Result {
	entropy := calcEntropy(password)
	rating := rateEntropy(entropy)
	breached := checkBreach(password)

	return Result{
		Entropy:  entropy,
		Rating:   rating,
		Breached: breached,
	}
}

func calcEntropy(password string) float64 {
	if len(password) == 0 {
		return 0
	}

	poolSize := 0
	hasLower, hasUpper, hasDigit, hasSymbol := false, false, false, false

	for _, ch := range password {
		switch {
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsDigit(ch):
			hasDigit = true
		default:
			hasSymbol = true
		}
	}

	if hasLower {
		poolSize += 26
	}
	if hasUpper {
		poolSize += 26
	}
	if hasDigit {
		poolSize += 10
	}
	if hasSymbol {
		poolSize += 32
	}

	return float64(len(password)) * math.Log2(float64(poolSize))
}

func rateEntropy(entropy float64) string {
	switch {
	case entropy < 28:
		return "Very Weak"
	case entropy < 36:
		return "Weak"
	case entropy < 60:
		return "Fair"
	case entropy < 80:
		return "Strong"
	default:
		return "Very Strong"
	}
}

func checkBreach(password string) bool {
	hash := fmt.Sprintf("%X", sha1.Sum([]byte(password)))
	prefix := hash[:5]
	suffix := hash[5:]

	resp, err := http.Get("https://api.pwnedpasswords.com/range/" + prefix)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	lines := strings.Split(string(body), "\n")
	for _, line := range lines {
		parts := strings.Split(strings.TrimSpace(line), ":")
		if len(parts) >= 1 && parts[0] == suffix {
			return true
		}
	}

	return false
}
