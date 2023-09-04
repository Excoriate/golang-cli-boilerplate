package cliutils

import (
	"fmt"
	"strings"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/utils"
)

type ExampleTemplateOptions struct {
	CLIName     string
	Command     string
	Options     string
	Explanation string
	Title       string
}

func GenerateExampleInCMD(examples []ExampleTemplateOptions) string {
	var allExamples string
	for i, example := range examples {
		lineLength := 65
		exampleTitle := fmt.Sprintf("Example %d: %s", i+1, example.Title)
		padding := strings.Repeat("=", lineLength)
		paddingToTitle := strings.Repeat(" ", (lineLength-len(exampleTitle))/2)

		wrappedExplanation := utils.WrapAtLength(example.Explanation,
			lineLength-6) // wrap the explanation

		// Preparation of the command usage string.
		usage := example.CLIName
		if example.Command != "" {
			usage = fmt.Sprintf("%s %s", usage, example.Command)
		}
		if example.Options != "" {
			usage = fmt.Sprintf("%s %s", usage, example.Options)
		}

		exampleStr := fmt.Sprintf(`
%s
%s%s%s

  1. Usage:
  -----------------------
  > %s

  2. Explanation:
  -----------------------
  > %s

%s
`, padding, paddingToTitle, exampleTitle, paddingToTitle, usage, wrappedExplanation, padding)
		allExamples += exampleStr
	}
	return allExamples
}
