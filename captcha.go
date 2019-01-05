package captcha

import (
	"strconv"
	"strings"
)

var (
	operator = [3]string{"+", "-", "x"}
	operand  = [9]string{
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
)

// Captcha is function create captcha
func Captcha(pattern, n1, op, n2 int) string {
	if pattern < 1 || pattern > 2 {
		return ""
	}
	if op < 1 || op > 3 {
		return ""
	}
	if !validateNumber(n1) || !validateNumber(n2) {
		return ""
	}

	if pattern == 1 {
		return joinWithSpace(strconv.Itoa(n1), operator[op-1], operand[n2-1])
	} else if pattern == 2 {
		return joinWithSpace(operand[n2-1], operator[op-1], strconv.Itoa(n1))
	}
	return ""
}

func joinWithSpace(s ...string) string {
	return strings.Join(s, " ")
}

func validateNumber(n int) bool {
	if n < 1 || n > 9 {
		return false
	}
	return true
}
