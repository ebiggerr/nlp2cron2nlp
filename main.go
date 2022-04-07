// The tool does NLP. Converts natural language to cron expression
package main

import (
	"flag"
	"fmt"
	"strings"
)

var (
	mode  *string
	input []string
)

func init() {
	mode = flag.String("mode", "nlp2cron", "the mode, either 'nlp2cron' OR 'cron2nlp'")
}

func main() {

	flag.Parse()
	input = flag.Args()

	switch len(input) {
	case 0:
		flag.PrintDefaults()
	case 1:
		fmt.Println("Your cron expression is too short.")
	}

	result, err := parse(strings.Join(input, " "), *mode)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
