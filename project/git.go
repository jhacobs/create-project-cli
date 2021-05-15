package project

import (
	"os"
	"os/exec"
)

const (
	Angular = "git@github.com:jhacobs/angular-blueprint.git"
	Laravel = "git@github.com:jhacobs/laravel-blueprint.git"
)

func Clone(name string, project string) error {
	cmd := exec.Command("git", "clone", project, name)
	return cmd.Run()
}

func RemoveGit() error {
	return os.RemoveAll(".git")
}

func Submodule(project string) error {
	cmd := exec.Command("git", "submodule", "add", project)
	return cmd.Run()
}
