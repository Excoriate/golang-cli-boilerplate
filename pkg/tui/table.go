package tui

import (
	"fmt"

	"github.com/pterm/pterm"
)

type TableOptions struct {
	Headers   []string
	Data      [][]string
	WithBoxed bool
}

func ShowTable(options TableOptions) error {
	headers := options.Headers
	data := options.Data

	if len(headers) == 0 {
		return fmt.Errorf("failed to show table: no headers provided")
	}

	if len(data) == 0 {
		return fmt.Errorf("failed to show table: no data provided")
	}

	if len(headers) != len(data[0]) {
		return fmt.Errorf("failed to show table: number of headers does not match number of columns")
	}

	tableData := pterm.TableData{headers}
	tableData = append(tableData, data...) // Include the data as it is provided, not just the indices

	if options.WithBoxed {
		_ = pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).Render()
		return nil
	}
	_ = pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()
	return nil
}
