package parser

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/utils"
)

// YamlFileIsValid YamlIsValid checks if the yaml file is valid.
func YamlFileIsValid(yamlFile string) error {
	if yamlFile == "" {
		return fmt.Errorf("the yaml file is required. It was received an empty string")
	}

	validExtensions := []string{".yaml", ".yml"}

	if !strings.Contains(yamlFile, ".") {
		return fmt.Errorf("the yaml file must have an extension. It was received: %s", yamlFile)
	}

	for _, validExtension := range validExtensions {
		if strings.Contains(yamlFile, validExtension) {
			return nil
		}
	}

	if err := utils.FileIsNotEmpty(yamlFile); err != nil {
		return fmt.Errorf("the yaml file must not be empty, error: %s", err.Error())
	}

	return fmt.Errorf("the yaml file must have a valid extension. It was received: %s", yamlFile)
}

// YamlStructureIsValid checks if the yaml file has a valid structure.
func YamlStructureIsValid(yamlFile string, schema interface{}) error {
	file, err := os.Open(yamlFile)
	if err != nil {
		return fmt.Errorf("could not open the yaml file: %s", err.Error())
	}

	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(schema); err != nil {
		return fmt.Errorf("the yaml file %s did not have a valid structure: %s", yamlFile, err.Error())
	}

	return nil
}
