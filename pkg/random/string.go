package random

import "math/rand"

func NewString(length int) string {
	dictionary := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"0123456789"

	result := make([]byte, length)
	for i := range result {
		result[i] = dictionary[rand.Intn(len(dictionary))]
	}
	return string(result)
}
