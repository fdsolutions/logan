package metadata_test

import (
	. "github.com/onsi/ginkgo"
)

const (
	TestGoal            string = "show:version"
	TestGoalWithContext string = "copy:file:unix"
	TestContext         string = "unix"
	TestUnkownContext   string = "unkown"
)

var (
	TestNoContext string

	TestGoalsWithNoContext = []string{
		"show:version",
		"show:help",
	}

	TestGoalsWithContext = []string{
		"show:help:unix",
		"copy:file:unix",
	}
)

var _ = Describe("metadata", func() {

})
