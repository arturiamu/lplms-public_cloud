package uuid

import (
	"github.com/google/uuid"
)

// GenerateUUID generate an RFC 4122 UUID -- 55c8f198-e88a-4fb5-9fef-090edca845b3
func GenerateUUID() string {
	return uuid.New().String()
}
