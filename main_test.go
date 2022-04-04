package main

import (
	"testing"
)

func TestParse(t *testing.T) {

	testcases := []struct {
		input, output string
	}{
		{"everyday noon", "0 12 * * *"},
		{"8AM every friday", "0 8 * * 5"},
		{"9PM on the first day of each year ", "0 21 1 1 *"},
	}

	for _, testcase := range testcases {
		cron := parse(testcase.input, "nlp2cron")
		if cron != testcase.output {
			t.Errorf("Input: %q, the cron notation: %q", cron, testcase.output)
		}
	}
}
