package action_test

import (
	"testing"

	"github.com/n3tuk/action-pull-requester/internal/action"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

// Define the structure of the table for TestCheckTitle, with the expected
// input, tests, and if the CheckTitle() function should pass or fail the test
// (i.e. the title is over the minimum length)
type CheckTitleTest struct {
	Title   string
	Minimum int
	Pass    bool
}

// Define the expected tests for TestCheckTitle()
var CheckTitleTests = []*CheckTitleTest{
	{
		Title:   "This is a Pull Request Title",
		Minimum: 20,
		Pass:    true,
	},
	{
		Title:   "This is a Pull Request Title",
		Minimum: 50,
		Pass:    false,
	},
	{
		Title:   "Invalid",
		Minimum: 20,
		Pass:    false,
	},
}

// Provide the GitTitle() function against the CheckTitleTest type so that it
// matches the PullRequestTitle interface required for CheckTitle()
func (c *CheckTitleTest) GetTitle() string {
	return c.Title
}

// Test the CheckTitle() function for testing the length of the titles on pull
// requests in GitHub
func TestCheckTitle(t *testing.T) {
	logger, _ := test.NewNullLogger()

	t.Run("minimum", func(t *testing.T) {
		for _, check := range CheckTitleTests {
			err := action.CheckTitle(logger, check, check.Minimum)
			if check.Pass {
				assert.NoError(t, err, "CheckTitle() returned an error when nil was expected")
			} else {
				assert.Error(t, err, "CheckTitle() did not return an error when one was expected")
			}
		}
	})
}
