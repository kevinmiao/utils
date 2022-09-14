package utils

import (
	"math/rand"
	"strings"
	"time"
)

var (
	letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func GenSecret(n int, args ...string) string {
	b := make([]rune, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Intn(62)]
	}

	var lower, upper bool
	if len(args) > 0 {
		format := args[0]
		if format == "lower" {
			lower = true
		} else if format == "upper" {
			upper = true
		}
	}
	s := string(b)
	if lower {
		s = strings.ToLower(s)
	} else if upper {
		s = strings.ToUpper(s)
	}
	return s
}
