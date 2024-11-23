package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func usage() {
	fmt.Println(`Usage: 

  btsnap [flags] <command> [flags] <arguments>

Flags:

  -h | --help    print this text

Commands:

  - shot  takes a snapshot of provided path`)
}

func run(args []string) error {
	var err error

	if len(args) < 2 {
		usage()
		return errors.New("No command provided")
	}

	if strings.HasPrefix(args[1], "-") {
		flag := args[1]

		switch flag {
		case "-h", "--help":
			usage()
			return nil
		default:
			usage()
			return fmt.Errorf("unknown flag %s", flag)
		}
	}

	command := args[1]
	b := newBt()

	switch command {
	case "shot":
		if len(args) < 3 {
			b.usage()
			return errors.New("Missing path of volume to snapshot")
		} else if len(args) < 4 {
			b.usage()
			return errors.New("Missing snapshot destination path")
		}
		if err = b.shot(args[2], args[3]); err != nil {
			return err
		}
	default:
		usage()
		return fmt.Errorf("unknown command %s", command)
	}

	return nil
}

func main() {
	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}
