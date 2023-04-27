package github

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type ConsumedLicenses struct {
	TotalSeatsConsumed  int `json:"total_seats_consumed"`
	TotalSeatsPurchased int `json:"total_seats_purchased"`
}

func GetConsumedLicenses(enterprise string, token string) *ConsumedLicenses {

	requestURL := fmt.Sprintf("https://api.github.com/enterprises/%s/consumed-licenses", enterprise)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)

	if err != nil {
		fmt.Println("Error creating request")
		os.Exit(1)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Error sending request", err)
		os.Exit(1)
	}

	defer res.Body.Close()

	resBody, _ := ioutil.ReadAll(res.Body)
	response := string(resBody)

	var consumedLicenses ConsumedLicenses
	json.Unmarshal([]byte(response), &consumedLicenses)

	return &consumedLicenses
}

func ReadFile(owner, repo, path, token string) ([]byte, error) {
	// Set up the request to get the file contents
	requestURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s", owner, repo, path)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	// Send the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Parse the response
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var fileData struct {
		Content string `json:"content"`
	}
	if err := json.Unmarshal(resBody, &fileData); err != nil {
		return nil, err
	}

	// Decode the base64-encoded file contents
	fileContents, err := base64.StdEncoding.DecodeString(fileData.Content)

	if err != nil {
		return nil, err
	}

	return fileContents, nil
}

func GetFileSHA(owner, repo, path, token string) (string, error) {
	// Set up the request to get the file metadata
	requestURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s", owner, repo, path)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	// Send the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// Parse the response
	var fileData struct {
		SHA string `json:"sha"`
	}
	if err := json.NewDecoder(res.Body).Decode(&fileData); err != nil {
		return "", err
	}

	return fileData.SHA, nil
}

func UpdateFile(owner, repo, path, token string, content []byte, message string) error {
	// Encode the new file contents as base64
	newContents := base64.StdEncoding.EncodeToString(content)

	// Get the SHA of the existing file
	sha, err := GetFileSHA(owner, repo, path, token)
	if err != nil {
		return err
	}

	// Set up the request to update the file contents
	requestURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s", owner, repo, path)
	requestBody := struct {
		Message string `json:"message"`
		Content string `json:"content"`
		SHA     string `json:"sha"`
	}{
		Message: message,
		Content: newContents,
		SHA:     sha,
	}
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPut, requestURL, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	// Send the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
