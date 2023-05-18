package action

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func RunChecks(log *logrus.Logger) error {
	log.
		Info("running checks")

	if err := CheckTitle(log); err != nil {
		return fmt.Errorf("check on title failed: %w", err)
	}

	if err := CheckDescription(log); err != nil {
		return fmt.Errorf("check on description failed: %w", err)
	}

	if err := CheckLabels(log); err != nil {
		return fmt.Errorf("check on labels failed: %w", err)
	}

	if err := CheckAssignee(log); err != nil {
		return fmt.Errorf("check on assignee failed: %w", err)
	}

	if err := CheckTitle(log); err != nil {
		return fmt.Errorf("check on title failed: %w", err)
	}

	return nil
}

func CheckTitle(log *logrus.Logger) error {
	log.
		Info("checking the title")

	return nil
}

func CheckDescription(log *logrus.Logger) error {
	log.
		Info("checking the description")

	return nil
}

func CheckLabels(log *logrus.Logger) error {
	log.
		Info("checking the labels")

	return nil
}

func CheckAssignee(log *logrus.Logger) error {
	log.
		Info("checking the assignee")

	return nil
}
