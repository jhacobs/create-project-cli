package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const (
	infoColor = "\u001b[36m"
	resetColor = "\u001b[0m"
	successColor = "\u001b[32;1m"
)

var rootCommand = &cobra.Command{
	Use: "create-project",
	Short: "Create a new project",
	Long: "create a new project based on a blueprint",
}

func Execute() {
	err := rootCommand.Execute()

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
