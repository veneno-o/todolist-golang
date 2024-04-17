package utils

import "strconv"

// string to uint
func StrToUint(str string) (uint, error) {
	num, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(num), nil
}
