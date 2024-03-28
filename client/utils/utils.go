package utils

import (
	"fmt"
	"io"
	"net/http"

	"github.com/test/pkg/models"
)

func SendRequest(client *http.Client, url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "My Client")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 404 {
		return nil, models.ErrNoRecord
	}

	return body, nil
}

func PrintCommand(s *models.Command) {
	fmt.Printf("ID: %d | Title: %s\n\nContent:\n%s\n\nExecution res:\n%s", s.ID, s.Title, s.Content, s.Exec_res)
}
