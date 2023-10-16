package identifiers

import (
	"math/rand"
	"strings"

	"github.com/google/uuid"
)

const (
	chars = "1234567890"
)

func GenerateUUID() string {
	return uuid.New().String()
}

func GenerateLinkID() string {
	var sb strings.Builder

	for i := 0; i < 10; i++ {
		ind := rand.Intn(len(chars))
		err := sb.WriteByte(chars[ind])
		if err != nil {
			i--
		}
	}

	return sb.String()
}
