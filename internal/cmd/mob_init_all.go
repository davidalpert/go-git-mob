package cmd

import (
	"fmt"
	"github.com/davidalpert/go-printers/v1"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type MobInitAllOptions struct {
	*printers.PrinterOptions
	BasePath string
	DryRun   bool
}

func NewMobInitAllOptions(s printers.IOStreams) *MobInitAllOptions {
	return &MobInitAllOptions{
		PrinterOptions: printers.NewPrinterOptions().WithStreams(s).WithDefaultTableWriter().WithDefaultOutput("text"),
		BasePath:       ".",
	}
}

func NewCmdMobInitAll(ioStreams printers.IOStreams) *cobra.Command {
	o := NewMobInitAllOptions(ioStreams)
	var cmd = &cobra.Command{
		Use:     "init-all [base-path]",
		Short:   "attempts to initializes git-mob for each of the sub-folders under <base-path> (which defaults to the current directory)",
		Aliases: []string{"ia", "initialize-all"},
		Args:    cobra.RangeArgs(0, 1),
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

	o.AddPrinterFlags(cmd.Flags())
	cmd.Flags().BoolVar(&o.DryRun, "dry-run", false, "dry-run will show you what will be done without doing it")

	return cmd
}

// Complete the options
func (o *MobInitAllOptions) Complete(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		o.BasePath = args[0]
	}

	return nil
}

// Validate the options
func (o *MobInitAllOptions) Validate() error {
	if _, err := os.Stat(o.BasePath); os.IsNotExist(err) {
		return fmt.Errorf("'%s' does not appear to be a valid path", o.BasePath)
	}

	return o.PrinterOptions.Validate()
}

type InitOneResult struct {
	FolderPath string `json:"folder_path"`
	HookPath   string `json:"hook_path"`
	Error      error  `json:"error,omitempty"`
}

// Run the command
func (o *MobInitAllOptions) Run() error {
	files, err := ioutil.ReadDir(o.BasePath)
	if err != nil {
		return fmt.Errorf("reading contents of '%s': %v", o.BasePath, err)
	}

	result := make([]InitOneResult, 0)
	errors := make([]InitOneResult, 0)
	for _, fileInfo := range files {
		if !fileInfo.IsDir() {
			continue
		}

		repoRoot := filepath.Join(o.BasePath, fileInfo.Name())
		if o.BasePath == "." {
			// "./" seems to get stripped off by filepath.Join() but adds context here so let's add it back
			repoRoot = fmt.Sprintf("%s%c%s", o.BasePath, os.PathSeparator, repoRoot)
		}
		hooksFolder := filepath.Join(repoRoot, ".git", "hooks")
		fileName := filepath.Join(hooksFolder, "prepare-commit-msg")
		fileContents := `#!/bin/sh

COMMIT_MSG_FILE=$1
COMMIT_SOURCE=$2
SHA1=$3

set -e

git mob hooks prepare-commit-msg "$COMMIT_MSG_FILE" $COMMIT_SOURCE $SHA1
`
		r := InitOneResult{FolderPath: repoRoot, HookPath: fileName}
		if _, err := os.Stat(hooksFolder); os.IsNotExist(err) {
			r.Error = fmt.Errorf("'%s' does not appear to be a valid git repo", repoRoot)
			errors = append(errors, r)
		} else if o.DryRun {
			result = append(result, r)
		} else if err := os.WriteFile(fileName, []byte(fileContents), 0755); err != nil {
			r.Error = fmt.Errorf("writing git hook: %v", err)
			errors = append(errors, r)
		} else {
			result = append(result, r)
		}
	}

	if o.FormatCategory() == "text" {
		printInitAllResult(o.Out, result, errors, o.DryRun)
		return nil
	}

	return o.WithTableWriter("processed folders", formatInitAllResultsAsTable(result, errors, o.DryRun)).WriteOutput(result)
}

func formatInitAllResultsAsTable(result []InitOneResult, errors []InitOneResult, dryRun bool) func(t *tablewriter.Table) {
	return func(t *tablewriter.Table) {
		t.SetHeader([]string{"Folder Considered", "Result"})
		for _, r := range append(result, errors...) {
			resultText := "initialized"
			if dryRun {
				resultText = "would have initialized"
			}
			if r.Error != nil {
				resultText = r.Error.Error()
			}
			t.Append([]string{r.FolderPath, resultText})
		}
	}
}

func printInitAllResult(w io.Writer, result []InitOneResult, errors []InitOneResult, dryRun bool) {
	if dryRun {
		fmt.Fprintf(w, "[whatif] would initialize prepare-commit-msg git hooks in:\n")
	} else {
		fmt.Fprintf(w, "initialized prepare-commit-msg git hooks in:\n")
	}
	for _, r := range result {
		fmt.Fprintf(w, "- '%s%c'\n", r.FolderPath, os.PathSeparator)
	}
	if dryRun {
		fmt.Fprintf(w, "\nthe following folders would not be not initialized:\n")
	} else {
		fmt.Fprintf(w, "git-mob will now append coauthor annotations to commit messages in those repos\n\nthe following folders were not initialized:\n")
	}
	for _, r := range errors {
		fmt.Fprintf(w, "- '%s%c' (%v)\n", r.FolderPath, os.PathSeparator, r.Error)
	}
	fmt.Fprintf(w, "\nhappy mobbing!\n")
}
