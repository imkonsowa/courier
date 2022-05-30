package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EnvStringNotExists(t *testing.T) {
	v := String("_DUMMY_", "not-exists")
	assert.Equal(t, v, "not-exists")
}

func Test_EnvStringExists(t *testing.T) {
	os.Setenv("_DUMMY_", "dummy")

	v := String("_DUMMY_", "not-exists")

	assert.Equal(t, v, "dummy")
}
