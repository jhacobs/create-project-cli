package project

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func CopyEnv(project string) error {
	content, err := ioutil.ReadFile(".env.example")
	parsedContent := string(content)

	if err != nil {
		return err
	}

	parsedContent = strings.Replace(parsedContent, "Laravel", project, 1)
	parsedContent = strings.Replace(parsedContent, "laravel_blueprint", project, 1)

	return ioutil.WriteFile(".env", []byte(parsedContent), os.ModePerm)
}

func CopyTestingEnv(project string) error {
	content, err := ioutil.ReadFile(".env.testing.example")
	parsedContent := string(content)

	if err != nil {
		return err
	}

	parsedContent = strings.Replace(parsedContent, "laravel_blueprint_test", project + "_test", 1)

	return ioutil.WriteFile(".env.testing", content, os.ModePerm)
}

func KeyGenerate() error {
	cmd := exec.Command("php", "artisan", "key:generate")
	return cmd.Run()
}
