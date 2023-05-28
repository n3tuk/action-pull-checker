package action

import (
	"fmt"
	"strings"

	"github.com/n3tuk/action-pull-requester/internal/github"

	"github.com/sirupsen/logrus"
)

type Options struct {
	TitleMinimum  int
	LabelPrefixes string
}

func RunChecks(logger *logrus.Logger, pull *github.PullRequest, options *Options) error {
	logger.
		WithFields(logrus.Fields{
			"owner":      pull.Owner,
			"repository": pull.Repository,
			"number":     pull.Number,
		}).
		Debug("running checks")

	if err := CheckTitle(logger, pull, options.TitleMinimum); err != nil {
		return fmt.Errorf("check on title failed: %w", err)
	}

	if err := CheckDescription(logger, pull); err != nil {
		return fmt.Errorf("check on description failed: %w", err)
	}

	prefixes := strings.Split(options.LabelPrefixes, ",")

	if err := CheckLabels(logger, pull, prefixes); err != nil {
		return fmt.Errorf("check on labels failed: %w", err)
	}

	return nil
}

func RunAutomations(logger *logrus.Logger, pull *github.PullRequest) error {
	logger.
		WithFields(logrus.Fields{
			"owner":      pull.Owner,
			"repository": pull.Repository,
			"number":     pull.Number,
		}).
		Debug("running automations")

	if err := CheckAssignee(logger, pull); err != nil {
		return fmt.Errorf("check on assignee failed: %w", err)
	}

	return nil
}

func CheckDescription(log *logrus.Logger, pullRequest *github.PullRequest) error {
	log.
		Debug("checking the description")

	log.
		Error("description check not yet supported")

	return nil
}

func CheckAssignee(log *logrus.Logger, pullRequest *github.PullRequest) error {
	log.
		Debug("checking the assignee")

	log.
		Error("assignee automation not yet supported")

	return nil
}
