package project

import (
	"io/ioutil"
	"os"
	"strings"
)

func UpdateAngularJson(project string) error {
	content, err := ioutil.ReadFile("angular.json")
	parsedContent := string(content)

	if err != nil {
		return err
	}

	parsedContent = strings.Replace(parsedContent, "angular-blueprint", project, -1)

	return ioutil.WriteFile("angular.json", []byte(parsedContent), os.ModePerm)
}
