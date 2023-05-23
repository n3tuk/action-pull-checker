package action

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func RunChecks(log *logrus.Logger) error {
	log.
		Info("starting checks and automations")

	if err := CheckTitle(log); err != nil {
		return fmt.Errorf("check on title failed: %w", err)
	}

	if err := CheckDescription(log); err != nil {
		return fmt.Errorf("check on description failed: %w", err)
	}

	if err := CheckLabels(log); err != nil {
		return fmt.Errorf("check on labels failed: %w", err)
	}

	return nil
}

func RunAutomations(log *logrus.Logger) error {
	log.
		Info("starting checks")

	if err := CheckAssignee(log); err != nil {
		return fmt.Errorf("check on assignee failed: %w", err)
	}

	return nil
}

func CheckTitle(log *logrus.Logger) error {
	log.
		Info("checking the title")

	log.
		Error("title check not yet supported")

	return nil
}

func CheckDescription(log *logrus.Logger) error {
	log.
		Info("checking the description")

	log.
		Error("description check not yet supported")

	return nil
}

func CheckLabels(log *logrus.Logger) error {
	log.
		Info("checking the labels")

	log.
		Error("label check not yet supported")

	return nil
}

func CheckAssignee(log *logrus.Logger) error {
	log.
		Info("checking the assignee")

	log.
		Error("assignee automation not yet supported")

	return nil
}
