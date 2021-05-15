package commands

import (
	"create-project/project"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var laravelAngularCommand = &cobra.Command{
	Use: "laravel-angular [name]",
	Short: "Create a Laravel Angular project",
	Long: "create a Laravel Angular project based on the blueprint at https://github.com/jhacobs",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			fmt.Println(string(infoColor), "Name needs to be provided")
			return nil
		}

		fmt.Println(string(infoColor), "Cloning project", string(resetColor))
		err := project.Clone(args[0], project.Laravel)

		if err != nil {
			return err
		}

		fmt.Println(string(infoColor), "Changing directory", string(resetColor))
		err = os.Chdir(args[0])

		if err != nil {
			return err
		}

		fmt.Println(string(infoColor), "Copy .env", string(resetColor))
		err = project.CopyEnv(args[0])

		if err != nil {
			return err
		}

		fmt.Println(string(infoColor), "Copy .env.testing", string(resetColor))
		err = project.CopyTestingEnv(args[0])

		if err != nil {
			return err
		}

		fmt.Println(string(infoColor), "Generate key", string(resetColor))
		err = project.KeyGenerate()

		if err != nil {
			return err
		}

		fmt.Println(string(infoColor), "Installing composer packages", string(resetColor))
		err = project.ComposerUpdate()

		if err != nil {
			return err
		}

		fmt.Println(string(infoColor), "Changing to angular directory", string(resetColor))
		err = os.Chdir("resources")

		if err != nil {
			return err
		}

		fmt.Println(string(infoColor), "Add angular submodule", string(resetColor))
		err = project.Submodule(project.Angular)

		if err != nil {
			return err
		}

		fmt.Println(string(infoColor), "Changing directory", string(resetColor))
		err = os.Chdir("angular")

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

		fmt.Println(string(successColor), "Finish creating laravel-angular project", string(resetColor))

		return nil
	},
}

func init() {
	rootCommand.AddCommand(laravelAngularCommand)
}
