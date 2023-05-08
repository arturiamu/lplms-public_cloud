package uuid

import (
	"fmt"
	"testing"
)

//
//
//
//
//
//
//
//
//
//
//745de68b-05fe-4db0-8524-713bfb9e5b3d
//79effecf-3cf2-4d66-af0d-0b2280880fed
//75d2c573-7c61-49fc-80a2-ab94391fed37
//4efe1a73-e829-4678-b377-c6cacdf68a0f
//d1763d6f-1896-4f3e-8f68-25d2bb41755f
//2f50d48e-8b37-4e9d-ab3d-4a980795cd16
//c14a4ea7-c205-4d18-94a2-bea49e0f5458
func TestGenerateUUID(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(GenerateUUID())
	}
}
