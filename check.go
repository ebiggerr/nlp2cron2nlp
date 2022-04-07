package main

import (
	"errors"
)

//this is constant
var (
	optionList = []string{"nlp2cron", "cron2nlp"}
)

func parse(input string, opts string) (string, error) {

	var result = ""

	if len(input) == 0 {
		return input, errors.New("error - Invalid expression")
	}

	if !contains(opts, optionList) {
		return input, errors.New("error - Invalid mode")
	}

	parser := parser{}

	if opts == optionList[0] {
		nlp2cron := &nlp2cron{}
		parser.setParsingMethod(nlp2cron)
		result = parser.parsing(input)
	}

	if opts == optionList[1] {
		cron2nlp := &cron2nlp{}
		parser.setParsingMethod(cron2nlp)
		result = parser.parsing(input)
	}

	return result, nil
}

func contains(target string, arr []string) bool {

	for _, e := range arr {
		if e == target {
			return true
		}
	}
	return false
}
