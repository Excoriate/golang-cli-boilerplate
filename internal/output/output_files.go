package output

import (
	"fmt"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/tui"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/parser"
	"github.com/Excoriate/golang-cli-boilerplate/pkg/utils"
)

type StOutOptions struct {
	Filename   string
	SaveInDisk bool
	Data       interface{}
}

func ShowYAML(options StOutOptions) error {
	// Convert the struct into a yaml file.
	yaml, err := parser.ConvertStructToYAML(options.Data)
	if err != nil {
		return fmt.Errorf("failed to show yaml: %s", err.Error())
	}

	if err := showInTerminal(yaml); err != nil {
		return err
	}

	if options.SaveInDisk {
		return saveInDisk(options.Filename, yaml, "yaml")
	}

	return nil
}

func ShowJSON(options StOutOptions) error {
	// Convert the struct into a json file.
	json, err := parser.ConvertStuctToJSON(options.Data)
	if err != nil {
		return fmt.Errorf("failed to show json: %s", err.Error())
	}

	// Show the json in the terminal.
	if err := showInTerminal(json); err != nil {
		return err
	}

	if options.SaveInDisk {
		return saveInDisk(options.Filename, json, "json")
	}

	return nil
}

func showInTerminal(content []byte) error {
	if err := tui.ShowFileInTerminal(content); err != nil {
		return fmt.Errorf("failed to show file in terminal: %s", err.Error())
	}
	return nil
}

func saveInDisk(filename string, content []byte, fileExtension string) error {
	if filename == "" {
		filename = fmt.Sprintf("output.%s", fileExtension)
	}

	if err := utils.WriteFileWithContent(filename, content); err != nil {
		return fmt.Errorf("failed to write file to disk: %s", err.Error())
	}

	return nil
}
