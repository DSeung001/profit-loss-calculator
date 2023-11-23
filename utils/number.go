package utils

import (
	"strconv"
	"unicode"
)

// InNumberCharRemove : 숫자가 아닌 값 제거해서 반환
func InNumberCharRemove(input string) (int, error) {
	var result string

	for _, char := range input {
		// 유니코드 포인트가 숫자인지 확인
		if unicode.IsDigit(char) {
			result += string(char)
		}
	}

	number, err := strconv.Atoi(result)
	if err != nil {
		return 0, err
	}

	return number, nil
}
