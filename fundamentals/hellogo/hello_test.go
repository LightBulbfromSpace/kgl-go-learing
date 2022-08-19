package hellogo_test

import (
	"github.com/LightBulbfromSpace/kld-go-learning/hellogo"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHello(t *testing.T) {
	t.Parallel()

	require.Equal(t, hellogo.Hello(), "Hello, GoLang")
}
