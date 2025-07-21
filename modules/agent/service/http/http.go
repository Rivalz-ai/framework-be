package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/Rivalz-ai/framework-be/framework/log"
)

var (
	bearerToken = "AAAAAAAAAAAAAAAAAAAAAALMxgEAAAAA6qSLveipvgJzoN44opQVvk%2FHZZg%3DJx57H2476ver4KF0W6OadvEcOFazfVKbYSuHw4Z4O38l6ziB3d"
)

// SendRequest sends an HTTP request with the given parameters and returns the status code and response body
func SendRequest(url string, method string, headers map[string]string, body interface{}) (int, map[string]interface{}, error) {
	// Convert body to JSON if not nil
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return 0, nil, err
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	// Create new request
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return 0, nil, err
	}

	// Add headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Set default Content-Type if body is present and Content-Type is not set
	if body != nil && req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	// Read response body
	var result map[string]interface{}
	if resp.Body != nil {
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return resp.StatusCode, nil, err
		}
	}

	return resp.StatusCode, result, nil
}

type TweetData struct {
	Data struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	} `json:"data"`
}

func extractTweetID(url string) (string, error) {
	re := regexp.MustCompile(`x\.com\/\w+\/status\/(\d+)`)
	matches := re.FindStringSubmatch(url)

	if len(matches) < 2 {
		return "", fmt.Errorf("tweet ID not found in URL")
	}

	return matches[1], nil
}
func GetTweetContent(tweetURL string) (string, error) {
	tweetID, err := extractTweetID(tweetURL)
	if err != nil {
		log.Error("Error extracting tweet ID: " + err.Error())
		return "", err
	}

	url := fmt.Sprintf("https://api.twitter.com/2/tweets/%s?tweet.fields=text", tweetID)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+bearerToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("Error getting tweet content: " + err.Error())
		return "", err
	}
	defer resp.Body.Close()

	var result TweetData
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Error("Error decoding tweet content: " + err.Error())
		return "", err
	}

	return result.Data.Text, nil
}
