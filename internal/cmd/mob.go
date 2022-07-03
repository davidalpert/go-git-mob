package cmd

import (
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/authors"
	"github.com/davidalpert/go-git-mob/internal/cfg"
	"github.com/davidalpert/go-git-mob/internal/cmd/utils"
	"github.com/davidalpert/go-git-mob/internal/msg"
	"github.com/davidalpert/go-git-mob/internal/revParse"
	"github.com/davidalpert/go-git-mob/internal/version"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"sort"
	"strings"
)

type MobOptions struct {
	*utils.PrinterOptions
	utils.IOStreams
	Initials               []string
	ListOnly               bool
	PrintVersion           bool
	CurrentGitUser         *authors.Author
	AllCoAuthorsByInitials map[string]authors.Author
}

func NewMobOptions(ioStreams utils.IOStreams) *MobOptions {
	return &MobOptions{
		IOStreams:      ioStreams,
		PrinterOptions: utils.NewPrinterOptions().WithDefaultTableWriter().WithDefaultOutput("text"),
	}
}

func NewCmdMob(ioStreams utils.IOStreams) *cobra.Command {
	o := NewMobOptions(ioStreams)
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

	o.PrinterOptions.AddPrinterFlags(cmd)

	cmd.Flags().BoolVarP(&o.ListOnly, "list", "l", false, "list which co-authors are available")
	cmd.Flags().BoolVarP(&o.PrintVersion, "version", "v", false, "print git-mob version")

	cmd.AddCommand(NewCmdMobInit(ioStreams))
	cmd.AddCommand(NewCmdMobHooks(ioStreams))
	cmd.AddCommand(NewCmdSolo(ioStreams))
	cmd.AddCommand(NewCmdCoauthors(ioStreams))
	cmd.AddCommand(NewCmdExplode(ioStreams))
	cmd.AddCommand(NewCmdImplode(ioStreams))
	cmd.AddCommand(NewCmdVersion(ioStreams))
	cmd.AddCommand(NewCmdPrint(ioStreams))

	return cmd
}

// Complete the options
func (o *MobOptions) Complete(cmd *cobra.Command, args []string) error {
	o.Initials = args

	if allCoAuthorsByInitials, err := cfg.ReadAllCoAuthorsFromFile(); err != nil {
		return err
	} else {
		o.AllCoAuthorsByInitials = allCoAuthorsByInitials
	}

	return nil
}

// Validate the options
func (o *MobOptions) Validate() error {
	if (o.ListOnly || o.PrintVersion) && 0 < len(o.Initials) {
		return fmt.Errorf("cannot configure a mob while listing availble coauthors or printing the version")
	}

	if !o.ListOnly && !o.PrintVersion {
		if !revParse.InsideWorkTree() {
			return fmt.Errorf("not inside a git repository working tree")
		}
		if a, err := cfg.GetUser(); err != nil {
			return err
		} else {
			o.CurrentGitUser = a
		}
	}

	return o.PrinterOptions.Validate()
}

// Run the command
func (o *MobOptions) Run() error {
	if o.ListOnly {
		return o.listCoAuthors()
	}

	if o.PrintVersion {
		versionCmd := NewCmdVersion(o.IOStreams)
		versionCmd.SetArgs([]string{}) // the version command doesn't accept the -v flag
		return versionCmd.Execute()
	}

	return o.setMob()
}

func (o *MobOptions) listCoAuthors() error {

	initials := make([]string, 0)
	for ii, _ := range o.AllCoAuthorsByInitials {
		initials = append(initials, ii)
	}

	sort.Strings(initials)

	if o.FormatCategory() == "text" {
		for _, ii := range initials {
			a := o.AllCoAuthorsByInitials[ii]
			fmt.Printf("%s %s %s\n", ii, a.Name, a.Email)
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
	if err := resetMob(); err != nil {
		return fmt.Errorf("resetMob: %v", err)
	}
	if err := cfg.AddCoAuthors(coauthors...); err != nil {
		return fmt.Errorf("AddCoAuthors: %v", err)
	}
	if err := msg.WriteGitMessage(coauthors...); err != nil {
		return fmt.Errorf("WriteGitMessage: %v", err)
	}

	parts := make([]string, len(o.Initials))
	meTag := o.CurrentGitUser.String()
	for i, initial := range o.Initials {
		for ii, a := range o.AllCoAuthorsByInitials {
			if strings.EqualFold(initial, ii) {
				parts[i] = a.String()
				break
			}
		}
	}

	return o.WriteStringln(strings.Join(append([]string{meTag}, parts...), "\n"))
}

// resetMob clears out the co-authors from global git config
func resetMob() error {
	return cfg.RemoveAllGlobal("git-mob.co-author")
}

// setCommitTemplate sets the local commit.template config setting to take advantage of `.gitmessage`
func setCommitTemplate() error {
	p, err := msg.CommitTemplatePath()
	if err != nil {
		return err
	}
	if !cfg.Has("commit.template") {
		return cfg.Set("commit.template", p)
	}
	return nil
}
