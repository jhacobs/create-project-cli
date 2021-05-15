package commands

import (
	"create-project/project"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var angularCommand = &cobra.Command{
	Use: "angular [name]",
	Short: "Create an Angular project",
	Long: "create an Angular project based on the blueprint at https://github.com/jhacobs",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			fmt.Println("Name needs to be provided")
			return nil
		}

		fmt.Println(string(infoColor), "Cloning project", string(resetColor))
		err := project.Clone(args[0], project.Angular)

		if err != nil {
			return err
		}

		fmt.Println(string(infoColor), "Changing directory", string(resetColor))
		err = os.Chdir(args[0])

		if err != nil {
			return err
		}

		fmt.Println(string(infoColor), "Updating angular.json", string(resetColor))
		err = project.UpdateAngularJson(args[0])

		if err != nil {
			return err
		}

		fmt.Println(string(infoColor), "Installing node modules", string(resetColor))
		err = project.NpmInstall()

		if err != nil {
			return err
		}

		fmt.Println(string(infoColor), "Run npm audit fix", string(resetColor))

		err = project.NpmAuditFix()

		if err != nil {
			return err
		}

		fmt.Println(string(infoColor), "Cleanup git", string(resetColor))

		err = project.RemoveGit()

		if err != nil {
			return err
		}

		fmt.Println(string(successColor), "Finish creating angular project", string(resetColor))

		return nil
	},
}

func init() {
	rootCommand.AddCommand(angularCommand)
}
