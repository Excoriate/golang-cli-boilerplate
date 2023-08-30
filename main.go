package main

import (
	"os"

	"github.com/Excoriate/golang-cli-boilerplate/cmd"
)

func main() {
	_ = os.Setenv("CLI_NAME", "golang-cli-boilerplate")
	_ = os.Setenv("ROOT_CMD_NAME", "golang-cli-boilerplate")
	cmd.Execute()
}
