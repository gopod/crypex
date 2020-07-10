package util

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

// GenerateUUID used as order id generator in NewOrder, ReplaceOrder
func GenerateUUID() string {
	return strings.ReplaceAll(uuid.NewV4().String(), "-", "")
}
