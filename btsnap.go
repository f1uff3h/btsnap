package main

import (
	"fmt"
	"os"
)

type btsnap struct {
	command        string
	totalArguments int
}

func (b *btsnap) checkCommand() error {
	if os.Args[1] != b.command {
		return fmt.Errorf("\n\nUnknown command %s", os.Args[1])
	}

	if len(os.Args) < b.totalArguments {
		fmt.Println("AA")
		return fmt.Errorf("\n\nInsufficient arguments provided")
	}

	return nil
}

func (b btsnap) shot() error {
	b.command = "shot"
	b.totalArguments = 3
	if err := b.checkCommand(); err != nil {
		return err
	}

	return nil
}

// func boot() error {
// 	bootCmd := &btsnap{
// 		command:        "boot",
// 		totalArguments: 2,
// 	}
//
// 	return bootCmd.cmd()
// }
