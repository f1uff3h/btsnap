package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	t.Run("no command", func(t *testing.T) {
		err := run([]string{})
		assert.Error(t, err)
		assert.Equal(t, "No command provided", err.Error())
	})

	t.Run("unknown flag", func(t *testing.T) {
		err := run([]string{"btsnap", "-asdf"})
		assert.Error(t, err)
		assert.Equal(t, "unknown flag -asdf", err.Error())
	})

	t.Run("no subvol to snapshot", func(t *testing.T) {
		err := run([]string{"btsnap", "shot"})
		assert.Error(t, err)
		assert.Equal(t, "Missing path of volume to snapshot", err.Error())
	})

	t.Run("no dest path", func(t *testing.T) {
		err := run([]string{"btsnap", "shot", "/root"})
		assert.Error(t, err)
		assert.Equal(t, "Missing snapshot destination path", err.Error())
	})

	t.Run("unknown command", func(t *testing.T) {
		err := run([]string{"btsnap", "asdf"})
		assert.Error(t, err)
		assert.Equal(t, "unknown command asdf", err.Error())
	})
}
