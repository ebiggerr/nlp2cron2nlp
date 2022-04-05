// The tool does NLP. Converts natural language to cron expression
package main

import (
	"flag"
	"fmt"
	"os"
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

	if len(input) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	result, err := parse(strings.Join(input, " "), *mode)
	if err != nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
		os.Exit(1)
	}

}
