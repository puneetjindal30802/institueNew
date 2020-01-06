package utils

import (
	"math/rand"
	"time"
)

var r *rand.Rand // Rand for this package.

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// /*
//  *	Generate 6 digit unique number
//  */
// func EncodeToString(max int) string {
// 	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
// 	result := make([]byte, max)
// 	for i := range result {
// 		result[i] = chars[r.Intn(len(chars))]
// 	}
// 	return string(result)
// }

/*
 * Function to generate a random string token to send in mails for password reset link.
 *
 * Returns result type string.
 */
func RandomStringToken(strlen int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, strlen)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	return string(result)
}
