package validator

import (
	"fmt"
	"regexp"
	"strconv"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
var integerRegex = regexp.MustCompile(`^-?[0-9]+$`)

type stringRule func(string) (bool, string)

func (v *Validator) CheckWithRules(key, data string, rules ...stringRule) {
	for _, rule := range rules {
		ok, msg := rule(data)
		if !ok {
			v.AddError(key, msg)
		}
	}
}

func IsValidLength(min, max int) stringRule {
	return func(data string) (bool, string) {
		if min > max {
			return false, "internal err: set min length is bigger than set max lenght"
		}

		if len(data) < min {
			return false, fmt.Sprintf("length must be at least %d characters", min)
		}

		if len(data) > max {
			return false, fmt.Sprintf("length must not exceed %d characters", max)
		}
		return true, ""
	}
}

func IsNotEmpty(data string) (bool, string) {
	if len(data) == 0 {
		return false, "must not be empty"
	}
	return true, ""
}

func IsValidEmail(email string) (bool, string) {
	if !emailRegex.MatchString(email) {
		return false, "invalid email format"
	}
	return true, ""
}

// bitSize 0 means any size; 0 <= bitSize <= 64
func IsInt(bitSize int) stringRule {
	return func(data string) (bool, string) {
		if !integerRegex.MatchString(data) {
			return false, "must be a valid integer"
		}

		if bitSize > 0 {
			_, err := strconv.ParseInt(data, 10, 64)
			if err != nil {
				return false, fmt.Sprintf("number is out of %d bit integer range", bitSize)
			}
		}

		return true, ""
	}
}
