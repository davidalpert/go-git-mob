package cmd

import (
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/cmd/utils"
	"github.com/davidalpert/go-git-mob/internal/revParse"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

type MobInitOptions struct {
	*utils.PrinterOptions
	utils.IOStreams
}

func NewMobInitOptions(ioStreams utils.IOStreams) *MobInitOptions {
	return &MobInitOptions{
		IOStreams:      ioStreams,
		PrinterOptions: utils.NewPrinterOptions().WithDefaultOutput("text"),
	}
}

func NewCmdMobInit(ioStreams utils.IOStreams) *cobra.Command {
	o := NewMobInitOptions(ioStreams)
	var cmd = &cobra.Command{
		Use:     "init",
		Short:   "initializes git-mob for the current repo",
		Aliases: []string{"i", "initialize"},
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := o.Complete(cmd, args); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			return o.Run()
		},
	}

	o.PrinterOptions.AddPrinterFlags(cmd)

	return cmd
}

// Complete the options
func (o *MobInitOptions) Complete(cmd *cobra.Command, args []string) error {
	return nil
}

// Validate the options
func (o *MobInitOptions) Validate() error {
	if !revParse.InsideWorkTree() {
		return fmt.Errorf("the init command only works inside a git repository's working folder")
	}

	return o.PrinterOptions.Validate()
}

// Run the command
func (o *MobInitOptions) Run() error {
	f := filepath.Join("hooks", "prepare-commit-msg")
	fileName, err := revParse.GitPath(f)
	if err != nil {
		return fmt.Errorf("the 'init' command is only valid inside a local git working directory: %v", err)
	}
	fileNameRel := revParse.GitPathRelativeToTopLevelDirectory(f)
	fileContents := `#!/bin/sh

COMMIT_MSG_FILE=$1
COMMIT_SOURCE=$2
SHA1=$3

set -e

git mob hooks prepare-commit-msg "$COMMIT_MSG_FILE" $COMMIT_SOURCE $SHA1
`
	if err := os.WriteFile(fileName, []byte(fileContents), 0755); err != nil {
		return fmt.Errorf("writing git hook: %v", err)
	}
	return o.IOStreams.WriteOutput(fmt.Sprintf("initialized local git hook: '%s'\ngit-mob will now help prepare commit messages in this repo\n", fileNameRel), o.PrinterOptions)
}
