package action

import (
	"fmt"
	"strings"

	"github.com/n3tuk/action-pull-requester/internal/github"

	"github.com/sirupsen/logrus"
)

type (
	PullRequest interface {
		PullRequestTitle
		PullRequestLabels
		PullRequestBody

		GetOwner() string
		GetRepository() string
		GetNumber() int
	}

	Options struct {
		TitleMinimum    int
		BodySplit       string
		BodyMinimum     int
		LabelPrefixes   string
		LabelPrefixMode string
	}
)

func RunChecks(logger *logrus.Logger, pull PullRequest, options *Options) error {
	logger.
		WithFields(logrus.Fields{
			"owner":      pull.GetOwner(),
			"repository": pull.GetRepository(),
			"number":     pull.GetNumber(),
		}).
		Debug("running checks")

	if err := CheckTitle(logger, pull, options.TitleMinimum); err != nil {
		return fmt.Errorf("check on title failed: %w", err)
	}

	if err := CheckBody(logger, pull, options.BodySplit, options.BodyMinimum); err != nil {
		return fmt.Errorf("check on body failed: %w", err)
	}

	prefixes := strings.Split(options.LabelPrefixes, ",")

	if err := CheckLabels(logger, pull, prefixes, options.LabelPrefixMode); err != nil {
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

func CheckAssignee(log *logrus.Logger, pullRequest *github.PullRequest) error {
	log.
		Debug("checking the assignee")

	log.
		Error("assignee automation not yet supported")

	return nil
}
