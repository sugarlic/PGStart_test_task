package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/test/pkg/models"
)

func getCommandList(client *http.Client) ([]*models.Command, error) {
	// запрос
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080", nil)
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

	// чтение ответа
	var res []*models.Command
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	for _, elem := range res {
		fmt.Println("ID: ", elem.ID, "Title: ", elem.Title, "Content: ", elem.Content)
	}

	return res, nil
}

func getCommandById(client *http.Client, id int) (*models.Command, error) {
	url := fmt.Sprintf("http://127.0.0.1:8080/command?id=%d", id) // http://127.0.0.1:8080?id=$

	// запрос
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

	// чтение ответа
	var res *models.Command
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	log.Println("ID: ", res.ID, "Title: ", res.Title, "Content: ", res.Content)

	return res, nil
}

func SendCommand(client *http.Client, filePath string) error {
	// формирование тела запроса
	scriptContent, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	data := models.Command{ID: 1, Title: filepath.Base(filePath), Content: string(scriptContent)}

	json_data, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// запрос
	url := "http://127.0.0.1:8080/command/create"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json_data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Println(string(responseBody))
	return nil
}

func main() {

	get_commands := flag.Bool("c", false, "Get list of commands")
	command_id := flag.Int("g", -1, "Get command by it's id")
	send_command := flag.String("f", "", "Send command on the server")

	flag.Parse()

	client := &http.Client{}

	if *get_commands {
		_, err := getCommandList(client)
		if err != nil {
			log.Println(err)
		}
	}
	if *command_id >= 0 {
		_, err := getCommandById(client, *command_id)
		if err != nil {
			log.Println(err)
		}
	}
	if *send_command != "" {
		err := SendCommand(client, *send_command)
		if err != nil {
			log.Println(err)
		}
	}

}
