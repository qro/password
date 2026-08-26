package generator

import (
	"crypto/rand"
	"errors"
	"math/big"
)

const (
	Uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lowercase = "abcdefghijklmnopqrstuvwxyz"
	Numbers   = "0123456789"
	Symbols   = "!@#$%^&*()-_=+[]{}|;:',.<>?/`~"
)

type Options struct {
	Length  int
	Upper   bool
	Lower   bool
	Digits  bool
	Symbols bool
}

func Generate(opts Options) (string, error) {
	charset := ""
	if opts.Upper {
		charset += Uppercase
	}
	if opts.Lower {
		charset += Lowercase
	}
	if opts.Digits {
		charset += Numbers
	}
	if opts.Symbols {
		charset += Symbols
	}

	if charset == "" {
		return "", errors.New("at least one character set must be selected")
	}
	if opts.Length < 1 {
		return "", errors.New("password length must be at least 1")
	}

	result := make([]byte, opts.Length)
	charsetLen := big.NewInt(int64(len(charset)))

	for i := 0; i < opts.Length; i++ {
		idx, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			return "", err
		}
		result[i] = charset[idx.Int64()]
	}

	return string(result), nil
}
