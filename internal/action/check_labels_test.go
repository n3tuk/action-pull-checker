package action_test

import (
	"testing"

	"github.com/n3tuk/action-pull-requester/internal/action"

	"github.com/google/go-github/v52/github"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

// Define the structure of the table for TestCheckLabels, with the expected
// input, tests, and if the CheckLabels() function should pass or fail the test
// (i.e. the required labels are added to the pull request)
type CheckLabelsTest struct {
	Name             string
	Labels           []*github.Label
	RequiredPrefixes []string
	Mode             string
	Pass             bool
}

var (
	// Pre-create the strings for the names as they must be passed as pointers
	nameTestOne     = "test/one"
	nameTestTwo     = "test/two"
	nameReleaseOne  = "release/one"
	nameReleaseTwo  = "release/two"
	namePriorityOne = "priority/one"
	namePriorityTwo = "priority/two"

	prefixTest     = "test/"
	prefixRelease  = "release/"
	prefixPriority = "priority/"

	labelTestOne     = &github.Label{Name: &nameTestOne}
	labelTestTwo     = &github.Label{Name: &nameTestTwo}
	labelReleaseOne  = &github.Label{Name: &nameReleaseOne}
	labelReleaseTwo  = &github.Label{Name: &nameReleaseTwo}
	labelPriorityOne = &github.Label{Name: &namePriorityOne}
	labelPriorityTwo = &github.Label{Name: &namePriorityTwo}

	// Define the expected tests for TestCheckLabels()
	CheckLabelsTests = []*CheckLabelsTest{
		{
			Name:             "all-types-match-and",
			Labels:           []*github.Label{labelTestOne, labelTestTwo, labelReleaseOne, labelReleaseTwo, labelPriorityOne, labelPriorityTwo},
			RequiredPrefixes: []string{prefixTest, prefixRelease, prefixPriority},
			Mode:             "all",
			Pass:             true,
		},
		{
			Name:             "all-types-match-all",
			Labels:           []*github.Label{labelTestOne, labelTestTwo, labelReleaseOne, labelReleaseTwo, labelPriorityOne, labelPriorityTwo},
			RequiredPrefixes: []string{prefixTest, prefixRelease, prefixPriority},
			Mode:             "all",
			Pass:             true,
		},
		{
			Name:             "simple-types-match-and",
			Labels:           []*github.Label{labelTestOne},
			RequiredPrefixes: []string{prefixTest},
			Mode:             "all",
			Pass:             true,
		},
		{
			Name:             "simple-types-match-any",
			Labels:           []*github.Label{labelTestOne},
			RequiredPrefixes: []string{prefixTest},
			Mode:             "all",
			Pass:             true,
		},
		{
			Name:             "empty-prefixes-case",
			Labels:           []*github.Label{labelTestOne, labelTestTwo, labelReleaseOne, labelReleaseTwo, labelPriorityOne, labelPriorityTwo},
			RequiredPrefixes: []string{},
			Mode:             "all",
			Pass:             true,
		},
		{
			Name:             "empty-labels-case-and",
			Labels:           []*github.Label{},
			RequiredPrefixes: []string{prefixTest, prefixRelease, prefixPriority},
			Mode:             "all",
			Pass:             false,
		},
		{
			Name:             "empty-labels-case-any",
			Labels:           []*github.Label{},
			RequiredPrefixes: []string{prefixTest, prefixRelease, prefixPriority},
			Mode:             "any",
			Pass:             false,
		},
		{
			Name:             "missing-labels",
			Labels:           []*github.Label{labelTestOne, labelTestTwo},
			RequiredPrefixes: []string{prefixPriority},
			Mode:             "all",
			Pass:             false,
		},
		{
			Name:             "partial-missing-labels-and",
			Labels:           []*github.Label{labelTestOne, labelReleaseOne, labelReleaseTwo},
			RequiredPrefixes: []string{prefixRelease, prefixPriority},
			Mode:             "all",
			Pass:             false,
		},
		{
			Name:             "partial-missing-labels-any",
			Labels:           []*github.Label{labelTestOne, labelReleaseOne, labelReleaseTwo},
			RequiredPrefixes: []string{prefixRelease, prefixPriority},
			Mode:             "any",
			Pass:             true,
		},
		{
			Name:             "suffix-test-1",
			Labels:           []*github.Label{labelTestOne},
			RequiredPrefixes: []string{prefixTest},
			Mode:             "all",
			Pass:             true,
		},
		{
			Name:             "suffix-test-2",
			Labels:           []*github.Label{labelTestTwo},
			RequiredPrefixes: []string{prefixTest},
			Mode:             "all",
			Pass:             true,
		},
		{
			Name:             "invalid-mode",
			Labels:           []*github.Label{labelTestTwo},
			RequiredPrefixes: []string{prefixTest},
			Mode:             "everything",
			Pass:             false,
		},
	}
)

// Provide the GitLabels() function against the CheckLabelsTest type so that it
// matches the PullRequestLabels interface required for CheckLabels()
func (c *CheckLabelsTest) GetLabels() []*github.Label {
	return c.Labels
}

// Test the CheckLabels() function for testing the presence of required labels
// on pull requests in GitHub
func TestCheckLabels(t *testing.T) {
	logger, _ := test.NewNullLogger()

	for _, check := range CheckLabelsTests {
		t.Run(check.Name, func(t *testing.T) {
			err := action.CheckLabels(logger, check, check.RequiredPrefixes, check.Mode)
			if check.Pass {
				assert.NoError(t, err, "CheckLabels() returned an error when nil was expected")
			} else {
				assert.Error(t, err, "CheckLabels() did not return an error when one was expected")
			}
		})
	}
}
