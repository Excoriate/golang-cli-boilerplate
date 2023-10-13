package output

import (
	"fmt"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/o11y"
	"github.com/Excoriate/golang-cli-boilerplate/pkg/tui"
)

type Options struct {
	OutputType   string
	Data         [][]string
	Filename     string
	SaveInDisk   bool
	TableHeaders []string
}

const tableOutputType = "table"

type TerminalOutput struct {
	tw     tui.MessageWriter
	logger o11y.LoggerInterface
}

type TerminalOutputWriter interface {
	Show(options Options) error
}

func NewTerminalOutput(tw tui.MessageWriter, l o11y.LoggerInterface) *TerminalOutput {
	return &TerminalOutput{
		tw:     tw,
		logger: l,
	}
}

func (t *TerminalOutput) Show(options Options) error {
	if options.Data == nil {
		t.logger.Error("failed to show terminal output, data is nil")
		return fmt.Errorf("failed to show terminal output, data is nil")
	}

	if options.OutputType == "" {
		options.OutputType = tableOutputType
	}

	if options.SaveInDisk && options.Filename == "" {
		t.logger.Error("failed to show terminal output, filename is empty but 'saveInDisk' is true")
		return fmt.Errorf("failed to show terminal output, filename is empty but 'saveInDisk' is true")
	}

	if options.OutputType == tableOutputType && len(options.TableHeaders) == 0 {
		t.logger.Error("failed to show terminal output, table headers are empty")
		return fmt.Errorf("failed to show terminal output, table headers are empty")
	}

	if options.OutputType == "yaml" {
		options.OutputType = "yml"
	}

	switch options.OutputType {
	case tableOutputType:
		if err := tui.ShowTable(tui.TableOptions{
			Headers: options.TableHeaders,
			Data:    options.Data,
		}); err != nil {
			t.logger.Error("failed to show terminal output, error showing table", err)
			return fmt.Errorf("failed to show terminal output, error showing table: %s", err.Error())
		}

	case "json":
		if err := ShowJSON(StOutOptions{
			Data:       options.Data,
			SaveInDisk: options.SaveInDisk,
			Filename:   options.Filename,
		}); err != nil {
			t.logger.Error("failed to show terminal output, error showing json", err)
			return fmt.Errorf("failed to show terminal output, error showing json: %s", err.Error())
		}

	case "yml":
		if err := ShowYAML(StOutOptions{
			Data:       options.Data,
			SaveInDisk: options.SaveInDisk,
			Filename:   options.Filename,
		}); err != nil {
			t.logger.Error("failed to show terminal output, error showing yaml", err)
			return fmt.Errorf("failed to show terminal output, error showing yaml: %s", err.Error())
		}
	}

	if options.SaveInDisk {
		t.logger.Info(fmt.Sprintf("The output was saved in disk in the file: %s", options.Filename))
	}

	return nil
}
