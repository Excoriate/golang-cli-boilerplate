package parser

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// YamlToStructFromFile converts a yaml file into a struct.
func YamlToStructFromFile(yamlFile string, schema interface{}) error {
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

// YamlToStructWithContent converts a yaml file into a struct.
func YamlToStructWithContent(yamlContent string, schema interface{}) error {
	decoder := yaml.NewDecoder(strings.NewReader(yamlContent))
	if err := decoder.Decode(schema); err != nil {
		return fmt.Errorf("the yaml file did not have a valid structure: %s", err.Error())
	}

	return nil
}

// ConvertTemplateIntoYAML converts a template into a yaml file.
func ConvertTemplateIntoYAML(tmpl bytes.Buffer) (interface{}, error) {
	if tmpl.Len() == 0 {
		return nil, errors.New("the template 'buffer' cannot be empty")
	}

	var result interface{}
	if err := yaml.Unmarshal(tmpl.Bytes(), &result); err != nil {
		return nil, fmt.Errorf("could not unmarshal the template: %s", err)
	}

	return result, nil
}

func ConvertStructToYAML(v interface{}) ([]byte, error) {
	if v == nil {
		return nil, fmt.Errorf("failed to marshal struct to yaml: the struct is nil")
	}

	yml, err := yaml.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal struct to yaml: %s", err.Error())
	}

	return yml, nil
}

func ConvertStuctToJSON(v interface{}) ([]byte, error) {
	if v == nil {
		return nil, fmt.Errorf("failed to marshal struct to json: the struct is nil")
	}

	content, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal struct to json: %s", err.Error())
	}

	return content, nil
}
