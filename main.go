package main

import (
	"fmt"
	"log"
	"os"
)

func usage() {
	// TODO: add commands dynamically and print usage as expected per command
	log.Println("Usage: btsnap <command> <arguments>")
}

// func cmd() error {
// 	argument := os.Args[2]
//
// 	if _, err := os.Stat(argument); os.IsNotExist(err) {
// 		return fmt.Errorf("%s\n\n%s is not a valid filepath", usage(), argument)
// 	} else if err != nil {
// 		// TODO: test if filepath is part of BTRFS subvolume
// 		return fmt.Errorf("unable to check filepath %s: %w", argument, err)
// 	}
//
// 	return nil
// }

func main() {
	var err error
	command := os.Args[1]
	b := &btsnap{}

	switch command {
	case "shot":
		err = b.shot()
	default:
		usage()
		err = fmt.Errorf("unknown command: %s", command)
	}

	if err != nil {
		log.Fatal(err)
	}
}
