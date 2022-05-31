package cmd

import (
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/cmd/utils"
	"github.com/davidalpert/go-git-mob/internal/version"
	"github.com/spf13/cobra"
	"os"
)

// cfgFile is an optional path to a configuration file used to initialize viper
var cfgFile string

// Execute builds the default root command and invokes it with os.Args
func Execute() {
	rootCmd := NewRootCmd(utils.DefaultOSStreams())

	rootCmd.SetArgs(os.Args[1:]) // without program

	err := rootCmd.Execute()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// NewRootCmd creates the 'root' command and configures it's nested children
func NewRootCmd(ioStreams utils.IOStreams) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "git-mob",
		Short: "A git plugin to help manage git coauthors.",
		Long: fmt.Sprintf(`git-mob %s

A git plugin to help manage git coauthors.
`, version.Detail.Version),
		SilenceUsage:  true,
		SilenceErrors: true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd:   false,
			DisableNoDescFlag:   false,
			DisableDescriptions: false,
			HiddenDefaultCmd:    true,
		},
		// Uncomment the following line if your bare application
		// has an action associated with it:
		//	Run: func(cmd *cobra.Command, args []string) { },
		//  RunE: func(cmd *cobra.Command, args []string) error { },
		Aliases: []string{},
	}

	// Register subcommands
	rootCmd.AddCommand(NewCmdVersion(ioStreams))
	rootCmd.AddCommand(NewCmdPrint(ioStreams))

	//rootCmd.PersistentFlags().BoolP("verbose", "vv", false, "enable verbose output")

	return rootCmd
}

func init() {
}
