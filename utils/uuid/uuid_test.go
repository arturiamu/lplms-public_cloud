package uuid

import (
	"fmt"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	fmt.Println(GenerateUUID())
}
