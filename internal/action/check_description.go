package action

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

// Define an Interface which maps to the GetBody function on the type so that we
// can simplify the requirements of the value passed and improve testability
type PullRequestBody interface {
	GetBody() string
}

// Check the body of the pull request to validate that it is at least the
// minimum length required, and return an error if it is not
func CheckBody(log *logrus.Logger, pull PullRequestBody, split string, minimum int) error {
	body := pull.GetBody()
	check := body

	log.
		WithFields(logrus.Fields{
			"body":    body,
			"minimum": minimum,
			"split":   split,
		}).
		Debug("check the body")

	if split != "" {
		splits := strings.Split(body, split)
		check = splits[0]
	}

	if len(check) < minimum {
		return fmt.Errorf("the body of the pull request has less than %d characters", minimum)
	}

	log.
		Info("the pull request body check has passed")

	return nil
}
