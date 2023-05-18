package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	Branch    string
	Commit    string
	Version   string
	BuildDate string

	versionCmd = &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Show the version of this application",
		Run: func(cmd *cobra.Command, args []string) {
			//nolint:forbidigo // This is a genuine output to the console
			fmt.Printf(
				strings.Join([]string{
					"v%s",
					"  built %s",
					"  commit %s (branch %s)",
				}, "\n")+"\n",
				Version, BuildDate, Commit, Branch)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
