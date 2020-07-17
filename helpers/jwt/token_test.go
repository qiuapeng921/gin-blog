package jwt

import (
	"fmt"
	"testing"
)

var token string

func TestGenerateToken(t *testing.T) {
	result, err := GenerateToken(1, "admin", "api")
	if err != nil {
		panic(err.Error())
	}
	token = result
	fmt.Println(token)
}

func TestParseToken(t *testing.T) {
	result, err := ParseToken(token)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(result)
}