package helper

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

func GenerateUUID() string {
	return strings.ReplaceAll(uuid.NewV4().String(), "-", "")
}
