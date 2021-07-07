package util

import (
	"github.com/shopspring/decimal"
	"strings"
)

// Contains some contains any one of conditions
func Contains(some string, conditions ...string) bool {
	for _, v := range conditions {
		if strings.Contains(some, v) {
			return true
		}
	}
	return false
}

func DivFloat64(a, b float64) (c float64) {
	c, _ = decimal.NewFromFloat(a).DivRound(decimal.NewFromFloat(b), 2).Float64()
	return
}
