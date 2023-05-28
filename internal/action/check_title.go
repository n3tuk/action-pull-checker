package action

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Define an Interface which maps to the GetTitle function on the type so that
// we can simplify the requirements of the value passed and improve testability
type PullRequestTitle interface {
	GetTitle() string
}

// Check the title of the pull request to validate that it is at least the
// minimum length required, and return an error if it is not
func CheckTitle(log *logrus.Logger, pull PullRequestTitle, minimum int) error {
	title := pull.GetTitle()

	log.
		WithFields(logrus.Fields{
			"title":   title,
			"minimum": minimum,
		}).
		Debug("check the title")

	if len(title) < minimum {
		return fmt.Errorf("the title of the pull request ('%s') has less than %d characters", title, minimum)
	}

	log.
		Info("the pull request title check has passed")

	return nil
}
