package main

//go:generate go run ./.tools/version_gen.go git-mob

import (
	_ "embed"
	"github.com/davidalpert/go-git-mob/internal/cmd"
)

func main() {
	cmd.Execute()
}
