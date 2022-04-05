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
		{"3PM on the first three day of each year ", "0 15 1,2,3 1 *"},
		{"10PM on every weekday", "0 22 * * 1-5"},
		{"once every 2 hours", "0 0-23/2 * * *"},
		{"once every 30 minutes", "*/30 * * * *"},
	}

	for _, testcase := range testcases {
		cron := parse(testcase.input, "nlp2cron")
		if cron != testcase.output {
			t.Errorf("Input: %q, the cron notation: %q", cron, testcase.output)
		}
	}
}
