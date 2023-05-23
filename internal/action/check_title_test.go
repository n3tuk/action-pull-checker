package action_test

import (
	"testing"

	"github.com/n3tuk/action-pull-requester/internal/action"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

// Define the structure of the table for TestCheckTitle, with the expected
// input, the test for the length, and if the CheckTitle() function should pass
// or fail the test
type CheckTitleTest struct {
	Input   string
	Minimum int
	Pass    bool
}

// Define the expected tests for TestCheckTitle()
var CheckTitleTests = []*CheckTitleTest{
	{
		Input:   "This is a Pull Request Title",
		Minimum: 20,
		Pass:    true,
	},
	{
		Input:   "This is a Pull Request Title",
		Minimum: 50,
		Pass:    false,
	},
	{
		Input:   "Invalid",
		Minimum: 20,
		Pass:    false,
	},
}

// Provide the GitTitle() function against the CheckTitleTest type so that it
// matches the PullRequest interface required for CheckTitle()
func (c *CheckTitleTest) GetTitle() string {
	return c.Input
}

// Test the CheckTitle() function for testing the length of the titles on pull
// requests in GitHub
func TestCheckTitle(t *testing.T) {
	logger, _ := test.NewNullLogger()

	t.Run("minimum", func(t *testing.T) {
		for _, check := range CheckTitleTests {
			err := action.CheckTitle(logger, check, check.Minimum)
			if check.Pass {
				assert.NoError(t, err, "The CheckTitle returned an error when nil was expected")
			} else {
				assert.Error(t, err, "The CheckTitle did not return an error when expected")
			}
		}
	})
}
