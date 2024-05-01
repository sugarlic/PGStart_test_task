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

	"github.com/test/client/utils"
	"github.com/test/pkg/models"
)

func getCommandList(client *http.Client) ([]*models.Command, error) {
	url := "http://158.160.89.127:8080"

	// запрос
	body, err := utils.SendRequest(client, url)
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
		fmt.Printf("ID: %d | Title: %s | ", elem.ID, elem.Title)
		if len(elem.Exec_res) > 0 {
			fmt.Println("Executed")
		} else {
			fmt.Println("Not executed")
		}
	}

	return res, nil
}

func getCommandById(client *http.Client, id int) (*models.Command, error) {
	url := fmt.Sprintf("http://158.160.89.127:8080/command?id=%d", id)

	// запрос
	body, err := utils.SendRequest(client, url)
	if err != nil {
		return nil, err
	}

	// чтение ответа
	var res *models.Command
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	utils.PrintCommand(res)
	return res, nil
}

func execCommandById(client *http.Client, id int) error {
	url := fmt.Sprintf("http://158.160.89.127:8080/command/exec?id=%d", id)

	// запрос
	respBody, err := utils.SendRequest(client, url)
	if err != nil {
		return err
	}

	// чтение ответа
	log.Println(string(respBody))

	return nil
}

func sendCommand(client *http.Client, filePath string) error {
	// формирование тела запроса
	scriptContent, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	data := models.Command{ID: 1, Title: filepath.Base(filePath),
		Content: string(scriptContent)}

	json_data, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// запрос
	url := "http://158.160.89.127:8080/command/create"

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

func deleteCommand(client *http.Client, id int) error {
	// запрос
	url := fmt.Sprintf("http://158.160.89.127:8080/command/delete?id=%d", id)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("User-Agent", "My Client")

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
	send_command := flag.String("f", "", "Send command to the server")
	exec_command := flag.Int("e", -1, "Exec command on the server")
	del_command := flag.Int("d", -1, "Delete command on the server")

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
		err := sendCommand(client, *send_command)
		if err != nil {
			log.Println(err)
		}
	}
	if *exec_command != -1 {
		err := execCommandById(client, *exec_command)
		if err != nil {
			log.Println(err)
		}
	}
	if *del_command != -1 {
		err := deleteCommand(client, *del_command)
		if err != nil {
			log.Println(err)
		}
	}

}
