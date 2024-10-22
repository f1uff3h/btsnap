package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/dennwc/btrfs"
)

var err error

// TODO: add --help flag
func main() {
	command, args, err := getCommandArguments()
	if err != nil {
		log.Fatal(err)
	}

	switch command {
	case "shot":
		// TODO: maybe move these checks somewhere else?
		if len(args) < 2 {
			usageShot()
			log.Fatal("not enough arguments provided")
		}
		if len(args) > 2 {
			usageShot()
			log.Fatal()
		}

		err = shot(args[0], args[1])
	default:
		err = fmt.Errorf("unknown command: %s", command)
	}

	if err != nil {
		log.Fatal(err)
	}
}

func getCommandArguments() (string, []string, error) {
	if len(os.Args) < 2 {
		usage()
		return "", nil, errors.New("no command provided")
	}
	if len(os.Args) < 3 {
		usage()
		return "", nil, errors.New("no arguments provided")
	}

	return os.Args[1], os.Args[2:], nil
}

func usage() {
	fmt.Println(`Usage: btsnap <command> <args>

Available commands:
  shot	take a snapshot of a BTRFS subvolume
		`)
}

// FIXME: destSubvolume should be destPath and should allow destPath/someNonExistantDir as everything after / will be created
func shot(sourceSubvolume, destSubvolume string) error {
	// NOTE: abstract checks to helper func ?
	if isSu, err := btrfs.IsSubVolume(sourceSubvolume); err != nil {
		return fmt.Errorf("unable to check if %s is a BTRFS subvolume: %w", sourceSubvolume, err)
	} else if !isSu {
		usageShot()
		return fmt.Errorf("path %s is not a BTRFS subvolume", sourceSubvolume)
	}
	if isSu, err := btrfs.IsSubVolume(destSubvolume); err != nil {
		return fmt.Errorf("unable to check if %s is a BTRFS subvolume: %w", destSubvolume, err)
	} else if !isSu {
		usageShot()
		return fmt.Errorf("path %s is not a BTRFS subvolume", destSubvolume)
	}

	// NOTE: provide option for ro?
	btrfs.SnapshotSubVolume(sourceSubvolume, destSubvolume, false)

	return err
}

func usageShot() {
	fmt.Println(`Usage: btsnap shot <source subvolume path> <destination subvolume path>

Takes a snapshot of <source subvolume path> and saves it at <destination subvolume path>
		`)
}
