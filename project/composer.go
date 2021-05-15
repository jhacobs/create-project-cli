package project

import "os/exec"

func ComposerUpdate() error {
	cmd := exec.Command("composer", "update")
	return cmd.Run()
}
