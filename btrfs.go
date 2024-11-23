package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/dennwc/btrfs"
)

type subvolSnapshotter interface {
	SnapshotSubVolume(subvol, dst string, ro bool) error
}

type btrfsSnap struct{}

func (b *btrfsSnap) SnapshotSubVolume(subvol, dst string, ro bool) error {
	return btrfs.SnapshotSubVolume(subvol, dst, ro)
}

type bt struct {
	snap subvolSnapshotter
}

func newBt() *bt {
	return &bt{
		&btrfsSnap{},
	}
}

func (b *bt) usage() {
	fmt.Println(`Usage:
	
  btsnap shot <subvolume-path> <destination-path>    takes snapshot of btrfs subvolume at <subvolume-path> and saves it at <destination-path>

  Note:
	
    - snapshot will automatically be named as <subvolume-name>-YYYY-MM-DDTHHMMSS`)
}

func (b *bt) shot(snapvol, destvol string) error {
	if _, err := os.Stat(destvol); err != nil {
		return errors.New("Destination path doesn't exist")
	}

	if !strings.HasSuffix(destvol, "/") {
		destvol = destvol + "/"
	}
	dest := destvol + filepath.Base(snapvol) + time.Now().Format("-20060102T151045")

	if err := b.snap.SnapshotSubVolume(snapvol, dest, false); err != nil {
		slog.Error("unable to take snapshot of subvolume: %s and save it at %s", snapvol, destvol)
		return err
	}

	return nil
}
