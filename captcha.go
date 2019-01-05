package captcha

import (
	"strconv"
	"strings"
)

var operator = [3]string{"+", "-", "x"}
var operand = [9]string{
	"ONE",
	"TWO",
	"THREE",
	"FOUR",
	"FIVE",
	"SIX",
	"SEVEN",
	"EIGHT",
	"NINE",
}

// Captcha is function create captcha
func Captcha(pat, n1, op, n2 int) string {
	if pat < 1 || pat > 2 {
		return ""
	}
	if op < 1 || op > 3 {
		return ""
	}
	if !validateNumber(n1) || !validateNumber(n2) {
		return ""
	}

	if pat == 1 {
		return joinString(strconv.Itoa(n1), operator[op-1], operand[n2-1])
	} else if pat == 2 {
		return joinString(operand[n2-1], operator[op-1], strconv.Itoa(n1))
	}
	return ""
}

func joinString(s ...string) string {
	res := []string{}
	for _, val := range s {
		res = append(res, val)
	}
	return strings.Join(res, " ")
}

func validateNumber(n int) bool {
	if n < 1 || n > 9 {
		return false
	}
	return true
}
