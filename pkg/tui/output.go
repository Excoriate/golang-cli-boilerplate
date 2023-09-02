package tui

import (
	"fmt"
	"os"

	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers/y"
	"github.com/alecthomas/chroma/styles"
)

func ShowFileInTerminal(content []byte) error {
	// Use the Chroma yaml lexer
	lexer := y.YAML
	// Use the Chroma default style
	style := styles.Get("github")
	if style == nil {
		style = styles.Fallback
	}
	// Use the standard out as the output buffer
	formatter := formatters.Get("terminal256")

	iterator, err := lexer.Tokenise(nil, string(content))
	if err != nil {
		return err
	}

	err = formatter.Format(os.Stdout, style, iterator)
	if err != nil {
		return err
	}

	fmt.Println()
	return nil
}
