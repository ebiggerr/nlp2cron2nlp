package main

import (
	"errors"
)

var (
	//this is constant
	optionList = []string{"nlp2cron", "cron2nlp"}
)

func parse(input string, opts string) (string, error) {

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
		if invalid := parser.validate(input); invalid != "" {
			return "Invalid characters/symbol found", errors.New("your input contains invalid characters/symbol " + invalid)
		} else {
			return parser.parsing(input)
		}
	}

	if opts == optionList[1] {
		cron2nlp := &cron2nlp{}
		parser.setParsingMethod(cron2nlp)
		if invalid := parser.validate(input); invalid != "" {
			return "Invalid characters/symbol found", errors.New("your input contains invalid characters/symbol " + invalid)
		} else {
			return parser.parsing(input)
		}
	}

	return "Something went wrong", errors.New("something went wrong")
}

func contains(target string, arr []string) bool {

	for _, e := range arr {
		if e == target {
			return true
		}
	}
	return false
}
