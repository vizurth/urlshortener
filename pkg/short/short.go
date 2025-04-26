package short

import "math/rand"

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateShortKey(length int) string {
	newUrl := make([]byte, length)
	for i := range newUrl {
		newUrl[i] = charset[rand.Intn(len(charset))]
	}
	return string(newUrl)
}
