package action_test

import (
	"strings"
	"testing"

	"github.com/n3tuk/action-pull-requester/internal/action"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

// Define the structure of the table for TestCheckBody, with the expected
// input, the test for the length, and if the CheckBody() function should pass
// or fail the test
type CheckBodyTest struct {
	Name    string
	Body    string
	Split   string
	Minimum int
	Pass    bool
}

// Define the expected tests for TestCheckBody()
var CheckBodyTests = []*CheckBodyTest{
	{
		Name: "test-simple-body",
		Body: strings.Join([]string{
			"This is the body text for the pull request.",
		}, "\n"),
		Split:   "---",
		Minimum: 35,
		Pass:    true,
	},
	{
		Name: "test-split-body-pass",
		Body: strings.Join([]string{
			"This is the body text for the pull request.",
			"",
			"---",
			"",
			"This is the footer to not be counted",
		}, "\n"),
		Split:   "---",
		Minimum: 35,
		Pass:    true,
	},
	{
		Name: "test-split-body-fail",
		Body: strings.Join([]string{
			"This is the body text for the pull request.",
			"",
			"---",
			"",
			"This is the footer to not be counted",
		}, "\n"),
		Split:   "---",
		Minimum: 60,
		Pass:    false,
	},
	{
		Name: "test-split-no-split",
		Body: strings.Join([]string{
			"This is the body text for the pull request.",
			"",
			"---",
			"",
			"This is the footer to not be counted",
		}, "\n"),
		Split:   "",
		Minimum: 60,
		Pass:    true,
	},
	{
		Name: "test-split-without-split",
		Body: strings.Join([]string{
			"This is the body text for the pull request.",
			"",
			"This is the footer to not be counted",
		}, "\n"),
		Split:   "---",
		Minimum: 60,
		Pass:    true,
	},
	{
		Name: "test-slim-split",
		Body: strings.Join([]string{
			"This is the body text for the pull request.",
			"---",
			"This is the footer to not be counted",
		}, "\n"),
		Split:   "---",
		Minimum: 35,
		Pass:    true,
	},
	{
		Name: "test-alternate-split",
		Body: strings.Join([]string{
			"This is the body text for the pull request.",
			"",
			"## Checklist",
			"",
			"- [ ] check or not",
		}, "\n"),
		Split:   "---",
		Minimum: 35,
		Pass:    true,
	},
}

// Provide the GitBody() function against the CheckBodyTest type so that it
// matches the PullRequest interface required for CheckBody()
func (c *CheckBodyTest) GetBody() string {
	return c.Body
}

// Test the CheckBody() function for testing the length of the bodys on pull
// requests in GitHub
func TestCheckBody(t *testing.T) {
	logger, _ := test.NewNullLogger()

	for _, check := range CheckBodyTests {
		t.Run(check.Name, func(t *testing.T) {
			err := action.CheckBody(logger, check, check.Split, check.Minimum)
			if check.Pass {
				assert.NoError(t, err, "The CheckBody returned an error when nil was expected")
			} else {
				assert.Error(t, err, "The CheckBody did not return an error when expected")
			}
		})
	}
}
