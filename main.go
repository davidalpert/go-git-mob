package main

//go:generate go run ./internal/version_gen.go git-mob

import (
	_ "embed"
	"github.com/davidalpert/go-git-mob/internal/cmd"
)

func main() {
	cmd.Execute()
}
