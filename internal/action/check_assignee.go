package action

import (
	"fmt"

	"github.com/google/go-github/v52/github"
	"github.com/sirupsen/logrus"
)

// Define an Interface which maps to the GetAssignee function on the type so that
// we can simplify the requirements of the value passed and improve testability
type PullRequestAssignee interface {
	GetUser() *github.User
	GetAssignees() []*github.User
	SetAssignee([]string) error
}

// Check the title of the pull request to validate that it is at least the
// minimum length required, and return an error if it is not
func CheckAssignees(log *logrus.Logger, pull PullRequestAssignee) error {
	var assigneeNames []string

	assignees := pull.GetAssignees()
	user := pull.GetUser()

	for _, assignee := range assignees {
		assigneeNames = append(assigneeNames, *assignee.Login)
	}

	log.
		WithFields(logrus.Fields{
			"assignees": assigneeNames,
			"user":      *user.Login,
		}).
		Debug("check the assignees")

	if len(assignees) == 0 {
		err := pull.SetAssignee([]string{*user.Login})
		if err != nil {
			return fmt.Errorf("unable to set the assignee: %w", err)
		}

		log.
			Info("the pull request has been assigned to the owner")
	} else {
		log.
			Info("the pull request assignee check has passed")
	}

	return nil
}
