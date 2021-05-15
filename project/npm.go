package project

import "os/exec"

func NpmUpdate() error {
	cmd := exec.Command("npm", "update")
	return cmd.Run()
}

func NpmInstall() error {
	cmd := exec.Command("npm", "install")
	return cmd.Run()
}

func NpmAuditFix() error {
	cmd := exec.Command("npm", "audit", "fix")
	return cmd.Run()
}
