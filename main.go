package main

//go:generate go run ./internal/version_gen.go git-mob

import (
	_ "embed"
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/cmd"
	"github.com/davidalpert/go-git-mob/internal/cmd/utils"
	"os"
)

func main() {
	rootCmd := cmd.NewRootCmd(utils.DefaultOSStreams())

	rootCmd.SetArgs(os.Args[1:]) // without program

	err := rootCmd.Execute()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
