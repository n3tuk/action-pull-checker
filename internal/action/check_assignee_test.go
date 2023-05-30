package action_test

import (
	"fmt"
	"testing"

	"github.com/n3tuk/action-pull-requester/internal/action"

	"github.com/google/go-github/v52/github"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

// Define the structure of the table for TestCheckAssignee, with the expected
// inputs, tests, and if the CheckAssignee() function should pass or fail the
// test (i.e. the Assignee has been set)
type CheckAssigneeTest struct {
	Name        string
	User        *github.User
	Assignees   []*github.User
	UserPresent bool
	GitHubError error
	Pass        bool
}

var (
	// Pre-create the strings for the logins as they must be passed as pointers
	userLogin     = "user-01"
	assigneLogin1 = "assignee-01"
	assigneLogin2 = "assignee-02"
	assigneLogin3 = "assignee-03"

	user      = &github.User{Login: &userLogin}
	assignee1 = &github.User{Login: &assigneLogin1}
	assignee2 = &github.User{Login: &assigneLogin2}
	assignee3 = &github.User{Login: &assigneLogin3}

	// Define the expected tests for TestCheckAssignee()
	CheckAssigneeTests = []*CheckAssigneeTest{
		{
			Name:        "unassigned",
			User:        user,
			Assignees:   nil,
			UserPresent: true,
			GitHubError: nil,
			Pass:        true,
		},
		{
			Name:        "assigned-with-user",
			User:        user,
			Assignees:   []*github.User{user},
			UserPresent: true,
			GitHubError: nil,
			Pass:        true,
		},
		{
			Name:        "assigned-with-user-and-others",
			User:        user,
			Assignees:   []*github.User{user, assignee1, assignee3},
			UserPresent: true,
			GitHubError: nil,
			Pass:        true,
		},
		{
			Name:        "assigned-without-user-with-others",
			User:        user,
			Assignees:   []*github.User{assignee1, assignee2, assignee3},
			UserPresent: false,
			GitHubError: nil,
			Pass:        true,
		},
		{
			Name:        "unassigned-with-error",
			User:        user,
			Assignees:   nil,
			UserPresent: false,
			GitHubError: fmt.Errorf("unknown error"),
			Pass:        false,
		},
	}
)

// Provide the GetUser() function against the CheckAssigneeTest type so that
// it matches the PullRequestAssignee interface required for CheckAssignee()
func (c *CheckAssigneeTest) GetUser() *github.User {
	return c.User
}

// Provide the GetAssignee() function against the CheckAssigneeTest type so that
// it matches the PullRequestAssignee interface required for CheckAssignee()
func (c *CheckAssigneeTest) GetAssignees() []*github.User {
	return c.Assignees
}

// Provide the SetAssignee() function against the CheckAssigneeTest type so that it
// matches the PullRequest interface required for CheckAssignee()
func (c *CheckAssigneeTest) SetAssignee(users []string) error {
	if c.GitHubError != nil {
		return fmt.Errorf("unable to set assignee on the pull request: %w", c.GitHubError)
	}

	for _, user := range users {
		login := user // take a copy of the login
		c.Assignees = append(c.Assignees, &github.User{Login: &login})
	}

	return nil
}

// Test the CheckAssignee() function for testing the status of the Assignees on
// the pull request, and see if the user is assigned to it if empty
func TestCheckAssignee(t *testing.T) {
	logger, _ := test.NewNullLogger()

	for _, check := range CheckAssigneeTests {
		t.Run(check.Name, func(t *testing.T) {
			err := action.CheckAssignees(logger, check)
			if check.Pass {
				assert.NoError(t, err, "The CheckAssignee() returned an error when nil was expected")
			} else {
				assert.Error(t, err, "The CheckAssignee() did not return an error when one was expected")
			}

			found := false
			for _, user := range check.GetAssignees() {
				if *user.Login == userLogin {
					found = true
				}
			}

			assert.Equal(t, check.UserPresent, found, "User %s in Assignees is expected to be %t but check was %t", userLogin, check.UserPresent, found)
		})
	}
}
