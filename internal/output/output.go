package output

import (
	"fmt"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/o11y"
	"github.com/Excoriate/golang-cli-boilerplate/pkg/tui"
)

type OutputOptions struct {
	OutputType   string
	Data         [][]string
	Filename     string
	SaveInDisk   bool
	TableHeaders []string
}

type TerminalOutput struct {
	tw     tui.MessageWriter
	logger o11y.LoggerInterface
}

type TerminalOutputWriter interface {
	Show(options OutputOptions) error
}

func NewTerminalOutput(tw tui.MessageWriter, l o11y.LoggerInterface) *TerminalOutput {
	return &TerminalOutput{
		tw:     tw,
		logger: l,
	}
}

func (t *TerminalOutput) Show(options OutputOptions) error {
	if options.Data == nil {
		t.logger.Error("failed to show terminal output, data is nil")
		return fmt.Errorf("failed to show terminal output, data is nil")
	}

	if options.OutputType == "" {
		options.OutputType = "table"
	}

	if options.SaveInDisk && options.Filename == "" {
		t.logger.Error("failed to show terminal output, filename is empty but 'saveInDisk' is true")
		return fmt.Errorf("failed to show terminal output, filename is empty but 'saveInDisk' is true")
	}

	if options.OutputType == "table" && len(options.TableHeaders) == 0 {
		t.logger.Error("failed to show terminal output, table headers are empty")
		return fmt.Errorf("failed to show terminal output, table headers are empty")
	}

	if options.OutputType == "yaml" {
		options.OutputType = "yml"
	}

	switch options.OutputType {
	case "table":
		if err := tui.ShowTable(tui.TableOptions{
			Headers: options.TableHeaders,
			Data:    options.Data,
		}); err != nil {
			t.logger.Error("failed to show terminal output, error showing table", o11y.LoggerArg{
				Key:   "error",
				Value: err.Error(),
			})
			return fmt.Errorf("failed to show terminal output, error showing table: %s", err.Error())
		}

	case "json":
		if err := ShowJSON(StOutOptions{
			Data:       options.Data,
			SaveInDisk: options.SaveInDisk,
			Filename:   options.Filename,
		}); err != nil {
			t.logger.Error("failed to show terminal output, error showing json", o11y.LoggerArg{
				Key:   "error",
				Value: err.Error(),
			})
			return fmt.Errorf("failed to show terminal output, error showing json: %s", err.Error())
		}

	case "yml":
		if err := ShowYAML(StOutOptions{
			Data:       options.Data,
			SaveInDisk: options.SaveInDisk,
			Filename:   options.Filename,
		}); err != nil {
			t.logger.Error("failed to show terminal output, error showing yaml", o11y.LoggerArg{
				Key:   "error",
				Value: err.Error(),
			})
			return fmt.Errorf("failed to show terminal output, error showing yaml: %s", err.Error())
		}
	}

	if options.SaveInDisk {
		t.logger.Info("file successfully saved", o11y.LoggerArg{
			Key:   "filename",
			Value: options.Filename,
		})
	}

	return nil
}
