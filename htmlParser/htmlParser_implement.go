package htmlParser

import (
	"os"
	"text/template"
	"github.com/google/uuid"
)

type htmlStructure struct {
	rootPath string
}

func New(rootPath string) HTMLParserInterface {
	return &htmlStructure{rootPath: rootPath}
}

func (a *htmlStructure) Create(templateName string, data interface{}) (string, error) {
	templateGenerator, err := template.ParseFiles(templateName)
	if err != nil {
		return "", err
	}

	fileName := a.rootPath + "/" + uuid.New().String() + ".html"

	fileWriter, err := os.Create(fileName)
	if err != nil {
		return "", err
	}

	if err = templateGenerator.Execute(fileWriter, data); err != nil {
		return "", err
	}

	return fileName, nil
}
