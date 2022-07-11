package command

import (
	"github.com/Irsad99/LectronicApp/src/config/server"

	"github.com/spf13/cobra"
)

var initCommand = cobra.Command{
	Short: "simple api with golang",
}

func init() {
	initCommand.AddCommand(server.ServeCmd)
}

func Run(args []string) error {
	initCommand.SetArgs(args)

	return initCommand.Execute()
}
