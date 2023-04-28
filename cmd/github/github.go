package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Issue struct {
	Body string `json:"body"`
}

type ConsumedLicenses struct {
	TotalSeatsConsumed  int `json:"total_seats_consumed"`
	TotalSeatsPurchased int `json:"total_seats_purchased"`
}

func GetConsumedLicenses(enterprise string, token string) (*ConsumedLicenses, error) {

	requestURL := fmt.Sprintf("https://api.github.com/enterprises/%s/consumed-licenses", enterprise)

	res, err := request(http.MethodGet, requestURL, token, nil)

	if err != nil {
		return nil, err
	}

	response := string(res)

	var consumedLicenses ConsumedLicenses
	json.Unmarshal([]byte(response), &consumedLicenses)

	return &consumedLicenses, nil
}

func CreateIssue(owner, repo, token, title string, body []byte, labels []string) error {
	// Set up the request to update the file contents
	requestURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", owner, repo)
	requestBody := struct {
		Title  string   `json:"title"`
		Body   string   `json:"body"`
		Labels []string `json:"labels"`
	}{
		Title:  title,
		Body:   string(body),
		Labels: labels,
	}
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	_, err = request(http.MethodPost, requestURL, token, requestBodyBytes)

	return err
}

func GetLatestIssueWithLabel(owner, repo, label, token string) (*Issue, error) {
	requestURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues?labels=%s", owner, repo, label)

	res, err := request(http.MethodGet, requestURL, token, nil)

	if err != nil {
		return nil, err
	}

	var issues []Issue
	err = json.Unmarshal(res, &issues)
	if err != nil {
		return nil, err
	}

	if len(issues) == 0 {
		return nil, nil
	}

	return &issues[0], nil
}

func CreateLabel(owner, repo, token, name string) error {
	// Set up the request to update the file contents
	requestURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/labels", owner, repo)
	requestBody := struct {
		Name string `json:"name"`
	}{
		Name: name,
	}
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	request(http.MethodPost, requestURL, token, requestBodyBytes)

	return nil
}

func request(method, url, token string, body []byte) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}
