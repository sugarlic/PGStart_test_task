package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/test/pkg/models"
	_ "github.com/test/pkg/models"
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
		fmt.Println(elem.ID, " ", elem.Title, " ", elem.Content)
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

	fmt.Println(res.ID, " ", res.Title, " ", res.Content)

	return res, nil
}

func main() {

	get_commands := flag.Bool("c", false, "Get list of commands")
	command_id := flag.Int("g", -1, "Get command by it's id")
	// set_command := flag.String("f", "", "Set command on the server")

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

}
