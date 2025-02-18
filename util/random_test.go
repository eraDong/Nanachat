package util

import (
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomString(t *testing.T) {
	t.Run("test random string", func(t *testing.T) {
		len := rand.IntN(100)
		randStr := RandomString(len)
		require.Len(t, randStr, len)
	})
}
