package errs

import "fmt"

func getErrPrefix(errName string) string {
	return fmt.Sprintf("%s error: ", errName)
}

type Opts struct {
	Error   error
	Details string
}
