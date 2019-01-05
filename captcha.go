package captcha

import (
	"strconv"
	"strings"
)

var operand = [3]string{"+", "-", "x"}
var stringNum = [10]string{
	"ZERO",
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
		return joinString(strconv.Itoa(n1), operand[op-1], stringNum[n2])
	} else if pat == 2 {
		return joinString(stringNum[n2], operand[op-1], strconv.Itoa(n1))
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
	if n < 0 || n > 9 {
		return false
	}
	return true
}
