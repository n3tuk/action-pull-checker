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

	titleMinimum    int    = 25
	bodySplit       string = "## Checklist"
	bodyMinimum     int    = 100
	labelPrefixes   string = strings.Join([]string{"release/", "type/", "update/"}, ",")
	labelPrefixMode string = "all"

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
	runCmd.Flags().IntVar(&titleMinimum, "title-minimum", titleMinimum, "The lower bound for the number of characters the title should contain")
	runCmd.Flags().StringVar(&bodySplit, "body-split", bodySplit, "The set of characters which split the body and the pull request template")
	runCmd.Flags().IntVar(&bodyMinimum, "body-minimum", bodyMinimum, "The lower bound for the number of characters the body should contain")
	runCmd.Flags().StringVar(&labelPrefixes, "label-prefixes", "", "A comma-separated list of label prefixes to check for on a pull request")
	runCmd.Flags().StringVar(&labelPrefixMode, "label-prefix-mode", labelPrefixMode, "Set if any one prefix, or all label prefixes, must match to pass")
	rootCmd.AddCommand(runCmd)
}

func RunChecks(cmd *cobra.Command, args []string) error {
	options := &action.Options{
		TitleMinimum:    titleMinimum,
		BodySplit:       bodySplit,
		BodyMinimum:     bodyMinimum,
		LabelPrefixes:   labelPrefixes,
		LabelPrefixMode: labelPrefixMode,
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
