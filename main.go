// The tool converts NLP to cron expression
package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	mode  *string
	input []string
)

func init() {
	mode = flag.String("mode", "nlp2cron", "mode of the tool: either nlp2cron OR cron2nlp")
}

func main() {

	flag.Parse()
	input = flag.Args()

	if len(input) == 0 {
		fmt.Println("Usage: main.go -mode nlp2cron everyday noon")
		flag.PrintDefaults()
		os.Exit(1)
	}

}
