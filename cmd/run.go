package cmd

import (
	"fmt"
	"strings"

	"github.com/n3tuk/action-pull-requester/internal/action"
	"github.com/n3tuk/action-pull-requester/internal/github"

	"github.com/spf13/cobra"
)

var (
	repository string
	number     int

	dryRun bool

	titleMinimum     int = 25
	labelPrefixes    string
	labelPrefixesAny bool

	// runCmd represents the run command
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run the checks and automations on the pull request",
		RunE:  RunChecks,
	}
)

func init() {
	runCmd.Flags().StringVarP(&repository, "repository", "r", "n3tuk/action-pull-requester", "The name of the repository to check")
	runCmd.Flags().IntVar(&number, "number", 0, "The number of the pull request to check")
	runCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "Only show what actions would be taken")
	runCmd.Flags().IntVar(&titleMinimum, "title-minimum", titleMinimum, "The minimum number of characters a title should contain")
	runCmd.Flags().StringVar(&labelPrefixes, "label-prefixes", "", "A comma-separated list of label prefixes to check for on a pull request")
	runCmd.Flags().BoolVar(&labelPrefixesAny, "label-prefixes-any", false, "Set that any label prefix can match to pass, rather than all")
	rootCmd.AddCommand(runCmd)
}

func RunChecks(cmd *cobra.Command, args []string) error {
	options := &action.Options{
		TitleMinimum:     titleMinimum,
		LabelPrefixes:    labelPrefixes,
		LabelPrefixesAny: labelPrefixesAny,
	}

	repo := strings.Split(repository, "/")

	pr, err := github.NewPullRequest(logger, repo[0], repo[1], number)
	if err != nil {
		return fmt.Errorf("failed to fetch the pull request: %w", err)
	}

	if err := action.RunChecks(logger, pr, options); err != nil {
		return fmt.Errorf("checks failed to run: %w", err)
	}

	return nil
}
