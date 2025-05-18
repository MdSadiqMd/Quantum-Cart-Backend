package helpers

import (
	"crypto/rand"
	"strconv"
)

func RandomNumbers(length int) (int, error) {
	const numbers = "0123456789"
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return 0, err
	}

	numLength := len(numbers)
	for i := range buffer {
		buffer[i] = numbers[int(buffer[i])%numLength]
	}

	if buffer[0] == '0' {
		buffer[0] = numbers[1+int(buffer[0])%(numLength-1)]
	}
	return strconv.Atoi(string(buffer))
}
