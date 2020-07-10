package util

import (
	"fmt"
	"log"
	"strings"
)

// UnmarshalError logs unmarshal error
func UnmarshalError(err error, exchange string) {
	log.Fatal(fmt.Errorf(
		"crypex: unmarshal response: [%s]: %v",
		strings.ToLower(exchange), err,
	))
}
