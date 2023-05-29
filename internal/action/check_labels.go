package action

import (
	"fmt"
	"strings"

	"github.com/google/go-github/v52/github"
	"github.com/sirupsen/logrus"
)

// Define an Interface which maps to the GetLabels function on the type so that
// we can simplify the requirements of the value passed and improve testability
type PullRequestLabels interface {
	GetLabels() []*github.Label
}

// Check the labels of the pull request to validate that the prefixes requested
// have been attached, and return an error if it is not
func CheckLabels(log *logrus.Logger, pull PullRequestLabels, prefixes []string, mode string) error {
	var attached, missing []string

	switch mode {
	case "any":
	case "all":
	default:
		return fmt.Errorf("the mode (%s) is not one of any or all", mode)
	}

	labels := pull.GetLabels()
	for _, label := range labels {
		attached = append(attached, *label.Name)
	}

	log.
		WithFields(logrus.Fields{
			"attached": strings.Join(attached, ","),
			"prefixes": strings.Join(prefixes, ","),
			"mode":     mode,
		}).
		Debug("checking the labels")

	for _, prefix := range prefixes {
		found := false

		log.
			WithFields(logrus.Fields{
				"prefix": prefix,
			}).
			Debug("checking for label prefix")

		for _, label := range attached {
			if strings.HasPrefix(label, prefix) {
				if mode == "any" {
					return nil // quick exit
				}

				found = true
			}
		}

		if !found {
			missing = append(missing, prefix)
		}
	}

	if len(missing) > 0 {
		return fmt.Errorf("labels with the prefix(es) %s have not been found", strings.Join(missing, ","))
	}

	log.
		Info("the pull request label check has passed")

	return nil
}
