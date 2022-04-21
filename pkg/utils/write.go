package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/asadbekGo/telegram-message-service/config"
)

var (
	client  = &http.Client{}
	method  = "GET"
	payload = strings.NewReader(``)
	url     = fmt.Sprintf("https://api.telegram.org/file/bot%s/", config.Load().TelegramBot.Token)
)

func Write(fileName, fileMedia, filePath string) error {
	req, err := http.NewRequest(method, url+filePath, payload)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var (
		check  bool
		number string
	)

	for _, c := range filePath {
		if c >= '0' && c <= '9' {
			check = true
			break
		}
	}

	if check {
		re := regexp.MustCompile("[0-9]+")
		number = re.FindAllString(filePath, -1)[0]
	}

	file, err := os.Create(fmt.Sprintf("receiveMedia/%s/%s", fileMedia, number+fileName))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return err
	}
	return nil
}
