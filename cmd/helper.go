package cmd

import (
	"os"
)

//write a function that reads a file and returns its content as a json object
func ReadFile(filename string) ([]byte, error) {
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func WriteFile(filename string, content []byte) error {
	err := os.WriteFile(filename, content, 0644)
	if err != nil {
		return err
	}

	return nil
}