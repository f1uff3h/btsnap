package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockBtrfs struct {
	snapvol string
	destvol string
	err     string
	ro      bool
}

func (m *mockBtrfs) SnapshotSubVolume(subvol, dst string, ro bool) error {
	var err error

	if m.err == "" {
		err = nil
	} else {
		err = errors.New(m.err)
	}

	m.snapvol = subvol
	m.destvol = dst
	m.ro = ro

	return err
}

func TestShot(t *testing.T) {
	t.Run("Invalid destination path", func(t *testing.T) {
		b := &bt{&mockBtrfs{}}

		err := b.shot("/unavailable", "/not/valid")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Destination path doesn't exist")
	})

	t.Run("Snapshot error", func(t *testing.T) {
		b := &bt{
			&mockBtrfs{
				err: "mockError",
			},
		}
		err := b.shot("/root", "/tmp")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "mockError")
	})

	t.Run("Success", func(t *testing.T) {
		m := &mockBtrfs{}
		b := &bt{m}
		err := b.shot("/root", "/tmp")
		assert.NoError(t, err)
		assert.Equal(t, "/root", m.snapvol)
		assert.Contains(t, m.destvol, "/tmp/root-")
		assert.False(t, m.ro)
	})
}
