package api

import (
	"errors"
	"log"
	"math/rand"
	"strings"
)

var ErrGenerationPostIdFailed = errors.New("post_id_generation_failed")

func (h *Handlers) generatePostId() (string, error) {
	const alphabet = "abdcdefghijklmnopqrtuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-"
	const idLength = 5
	const maxAttempts = 5

	for attempt := 0; attempt < maxAttempts; attempt++ {
		var idBuilder strings.Builder

		for i := 0; i < idLength; i++ {
			char := alphabet[rand.Intn(len(alphabet))]
			_ = idBuilder.WriteByte(char)
		}

		id := idBuilder.String()

		idAlreadyUsed := h.Database.Check(id)
		if idAlreadyUsed {
			log.Printf("Got collision for id #{id}. Retry")
			continue
		}

		return id, nil
	}

	return "", ErrGenerationPostIdFailed
}
