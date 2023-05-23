package action

import (
	"fmt"

	"github.com/n3tuk/action-pull-requester/internal/github"

	"github.com/sirupsen/logrus"
)

type Options struct {
	TitleMinimum int
}

func RunChecks(logger *logrus.Logger, pullRequest *github.PullRequest, options *Options) error {
	logger.
		WithFields(logrus.Fields{
			"owner":      pullRequest.Owner,
			"repository": pullRequest.Repository,
			"number":     pullRequest.Number,
		}).
		Debug("running checks")

	if err := CheckTitle(logger, pullRequest, options.TitleMinimum); err != nil {
		return fmt.Errorf("check on title failed: %w", err)
	}

	if err := CheckDescription(logger, pullRequest); err != nil {
		return fmt.Errorf("check on description failed: %w", err)
	}

	if err := CheckLabels(logger, pullRequest); err != nil {
		return fmt.Errorf("check on labels failed: %w", err)
	}

	return nil
}

func RunAutomations(logger *logrus.Logger, pullRequest *github.PullRequest) error {
	logger.
		WithFields(logrus.Fields{
			"owner":      pullRequest.Owner,
			"repository": pullRequest.Repository,
			"number":     pullRequest.Number,
		}).
		Debug("running automations")

	if err := CheckAssignee(logger, pullRequest); err != nil {
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

func CheckLabels(log *logrus.Logger, pullRequest *github.PullRequest) error {
	log.
		Debug("checking the labels")

	log.
		Error("label check not yet supported")

	return nil
}

func CheckAssignee(log *logrus.Logger, pullRequest *github.PullRequest) error {
	log.
		Debug("checking the assignee")

	log.
		Error("assignee automation not yet supported")

	return nil
}
