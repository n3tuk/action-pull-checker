package cmd

import (
	"fmt"

	"github.com/n3tuk/action-pull-requester/internal/action"

	"github.com/spf13/cobra"
)

var (
	dryRun bool

	// runCmd represents the run command
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run the checks and automations on the pull request",
		RunE:  RunChecks,
	}
)

func init() {
	runCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "Only show what actions would be taken")
	rootCmd.AddCommand(runCmd)
}

func RunChecks(cmd *cobra.Command, args []string) error {
	if err := action.RunChecks(logger); err != nil {
		return fmt.Errorf("checks failed to run: %w", err)
	}

	return nil
}
