package main

import (
	"github.com/ISauve/nikos/cmd"
)

func main() {
	cmd.SetupCommands()
	cmd.RootCmd.Execute()
}
