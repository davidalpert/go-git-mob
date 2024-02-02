package cmd

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/davidalpert/go-git-mob/internal/authors"
	"github.com/davidalpert/go-git-mob/internal/gitCommands"
	"github.com/davidalpert/go-git-mob/internal/gitConfig"
	"github.com/davidalpert/go-git-mob/internal/gitMessage"
	"github.com/davidalpert/go-git-mob/internal/gitMobCommands"
	"github.com/davidalpert/go-git-mob/internal/version"
	"github.com/davidalpert/go-printers/v1"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type MobOptions struct {
	*printers.PrinterOptions
	Initials                 []string
	ListOnly                 bool
	PrintMob                 bool
	PrintCoauthorsFilePath   bool
	PrintVersion             bool
	OverrideAuthorByInitials string
	CurrentGitUser           *authors.Author
	AllCoAuthorsByInitials   map[string]authors.Author
}

func NewMobOptions(s printers.IOStreams) *MobOptions {
	return &MobOptions{
		PrinterOptions: printers.NewPrinterOptions().WithDefaultTableWriter().WithDefaultOutput("text"),
	}
}

func NewCmdMob(s printers.IOStreams) *cobra.Command {
	o := NewMobOptions(s)
	var cmd = &cobra.Command{
		Use:   "mob",
		Short: "configure co-authors",
		Long: fmt.Sprintf(`git-mob %s

A git plugin to help manage git coauthors.

Examples:
   $ git mob jd                                      # Set John as co-authors
   $ git solo                                        # Return to working by yourself (i.e. unset all co-authors)
   $ git mob -l                                      # Show a list of all co-authors, John Doe should be there
`, version.Detail.Version),
		Args: cobra.MinimumNArgs(0),
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

	cmd.Flags().BoolVarP(&o.ListOnly, "list", "l", false, "list which co-authors are available")
	cmd.Flags().BoolVarP(&o.PrintCoauthorsFilePath, "print-coauthors-file-path", "p", false, "print path to the coauthors file")
	cmd.Flags().BoolVarP(&o.PrintVersion, "version", "v", false, "print git-mob version")
	cmd.Flags().StringVarP(&o.OverrideAuthorByInitials, "override-author", "a", "", "replace the current author with the co-author matching these initials")

	cmd.AddCommand(NewCmdMobInit(s))
	cmd.AddCommand(NewCmdMobInitAll(s))
	cmd.AddCommand(NewCmdMobHooks(s))
	cmd.AddCommand(NewCmdSolo(s))
	cmd.AddCommand(NewCmdCoauthors(s))
	cmd.AddCommand(NewCmdMobRehash(s))
	cmd.AddCommand(NewCmdMobUninstall(s))
	cmd.AddCommand(NewCmdVersion(s))
	cmd.AddCommand(NewCmdPrint(s))

	return cmd
}

// Complete the options
func (o *MobOptions) Complete(cmd *cobra.Command, args []string) error {
	o.Initials = args

	if allCoAuthorsByInitials, err := completeExistingCoauthorsByInitial(o.PrinterOptions); err != nil {
		return err
	} else {
		o.AllCoAuthorsByInitials = allCoAuthorsByInitials
	}

	if len(args) == 0 {
		o.PrintMob = true
	}

	return nil
}

// Validate the options
func (o *MobOptions) Validate() error {
	if (o.ListOnly || o.PrintVersion) && 0 < len(o.Initials) {
		return fmt.Errorf("cannot configure a mob while listing availble coauthors or printing the version")
	}

	if !o.ListOnly && !o.PrintVersion {
		if a, err := gitMobCommands.GetGitAuthor(); err != nil {
			return err // includes configWarning
		} else {
			o.CurrentGitUser = a
		}
	}

	if o.OverrideAuthorByInitials != "" {
		if _, found := o.AllCoAuthorsByInitials[o.OverrideAuthorByInitials]; !found {
			return fmt.Errorf("cannot find coauthor '%s' to use as an override, to list available coauthors use: git mob --list", o.OverrideAuthorByInitials)
		}
	}

	return o.PrinterOptions.Validate()
}

// Run the command
func (o *MobOptions) Run() error {
	// Help is handled by cobra.Command infrastructure

	if o.PrintVersion {
		versionCmd := NewCmdVersion(o.IOStreams)
		versionCmd.SetArgs([]string{}) // the version command doesn't accept the -v flag
		return versionCmd.Execute()
	}

	if o.PrintCoauthorsFilePath {
		return o.PrinterOptions.WriteOutput(authors.CoAuthorsFilePath)
	}

	if o.ListOnly {
		return o.listCoAuthors()
	}

	if o.OverrideAuthorByInitials != "" {
		a := o.AllCoAuthorsByInitials[o.OverrideAuthorByInitials]
		// override in memory for this command
		o.CurrentGitUser = &a

		// override in git config for other commands
		if err := gitMobCommands.SetGitAuthorGlobal(&a); err != nil {
			return err
		}
	}
	return o.runMob()
}

func (o *MobOptions) runMob() error {
	if o.PrintMob {
		if err := o.printMob(); err != nil {
			return err
		}
		return nil
	}

	return o.setMob()
}

func (o *MobOptions) printMob() error {
	primaryGitAuthor := o.CurrentGitUser

	currentMob := authors.AuthorList{
		Members: []*authors.Author{primaryGitAuthor},
	}

	if gitMobCommands.IsCoAuthorsSet() {
		aa, err := gitMobCommands.GetCoAuthors()
		if err != nil {
			return err
		}
		for _, a := range aa {
			b := a // allocate a local copy so that we don't take the pointer of the iterator
			currentMob.Members = append(currentMob.Members, &b)
		}
	}

	// sort mob members but keep the primary author at the top
	currentMob.SortBy(func(left, right *authors.Author) bool {
		if left.Email == primaryGitAuthor.Email {
			return true
		}
		if right.Email == primaryGitAuthor.Email {
			return false
		}
		return left.Name < right.Name
	})

	if (o.FormatCategory() == "text" || o.FormatCategory() == "table") && !gitMobCommands.UseLocalTemplate() && gitCommands.UsingLocalTemplate() {
		fmt.Fprintf(o.IOStreams.Out, CommitTemplateWorningMessage)
	}

	return o.WithTableWriter("current mob", currentMob.WriteToTable).WriteOutput(currentMob)
}

const CommitTemplateWorningMessage = `
Warning: local commit.template value detected

Using local commit.template could mean your template does not have selected co-authors appended after switching projects.

If you do not use commit.template (e.g. it was added by an earlier version of go-git-mob) you can safely remove it:

    git config --local --unset commit.template

If your team or project uses a local commit.template value you can silence this warning for this repo with:

    git config --local git-mob.use-local-template true

Happy Mobbing!

`

func (o *MobOptions) listCoAuthors() error {
	initials := make([]string, 0)
	for ii, _ := range o.AllCoAuthorsByInitials {
		initials = append(initials, ii)
	}

	sort.Strings(initials)

	if o.FormatCategory() == "text" {
		for _, ii := range initials {
			a := o.AllCoAuthorsByInitials[ii]
			fmt.Fprintf(o.Out, "%s %s %s\n", ii, a.Name, a.Email)
		}
		return nil
	}

	output := make([]authors.AuthorWithInitials, len(initials))
	for i, ii := range initials {
		a := o.AllCoAuthorsByInitials[ii]
		output[i] = authors.AuthorWithInitials{
			Initials: ii,
			Name:     a.Name,
			Email:    a.Email,
		}
	}

	return o.WithTableWriter("available co-authors", func(table *tablewriter.Table) {
		table.SetHeader([]string{"Initials", "Name", "Email"})
		for _, ii := range initials {
			a := o.AllCoAuthorsByInitials[ii]
			table.Append([]string{ii, a.Name, a.Email})
		}

	}).WriteOutput(output)
}

func (o *MobOptions) setMob() error {
	// authorList is preloaded as o.AllCoauthorsByInitial
	coauthors := make([]authors.Author, len(o.Initials))
	for i, initial := range o.Initials {
		for ii, a := range o.AllCoAuthorsByInitials {
			if strings.EqualFold(initial, ii) {
				coauthors[i] = a
				break
			}
		}
		if coauthors[i].Name == "" {
			return fmt.Errorf("author with initials '%s' not found; run 'git mob --list' to see a list of available co-authors", initial)
		}
	}

	if err := setCommitTemplate(); err != nil {
		return fmt.Errorf("setCommitTemplate: %v", err)
	}
	if err := gitMobCommands.ResetMob(); err != nil {
		return fmt.Errorf("ResetMob: %v", err)
	}
	if err := gitMobCommands.AddCoAuthors(coauthors...); err != nil {
		return fmt.Errorf("AddCoAuthors: %v", err)
	}

	return o.printMob()
}

// setCommitTemplate sets the local commit.template config setting to take advantage of `.gitmessage`
func setCommitTemplate() error {
	if !gitCommands.HasTemplatePath() {
		return gitCommands.SetTemplatePath(gitMessage.CommitTemplatePath())
	}
	return nil
}

func completeExistingCoauthorsByInitial(po *printers.PrinterOptions) (map[string]authors.Author, error) {
	allCoAuthorsByInitials, err := gitConfig.ReadAllCoAuthorsFromFile()
	if err != nil {
		var dupesErr authors.DuplicateInitialsError
		switch {
		case errors.As(err, &dupesErr):
			_, _ = fmt.Fprintf(po.ErrOut, "found in: %s\n\n", authors.CoAuthorsFilePath)
			for ii, aa := range dupesErr.DuplicateAuthorsByInitial {
				for _, a := range aa {
					_, _ = fmt.Fprintf(po.ErrOut, "%s %s %s\n", ii, a.Name, a.Email)
				}
			}
			_, _ = fmt.Fprintf(po.ErrOut, "\n")
		default:
		}
		return nil, err
	}
	return allCoAuthorsByInitials, nil
}
