package clienthttp

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/asadbekGo/telegram-message-service/config"
)

type ClientAPI struct {
	HttpClient http.Client
}

var (
	ErrEmptyOrInvalidBotToken = errors.New("invalid bot token")
)

// New returns new instance of ClientApi for telegram bot
func New(args *ClientAPI) (*ClientAPI, error) {
	err := args.validate()
	if err != nil {
		return &ClientAPI{}, err
	}

	subscribeAPI := ClientAPI{
		HttpClient: args.HttpClient,
	}

	return &subscribeAPI, nil
}

func (s *ClientAPI) validate() error {
	if config.Load().TelegramBot.Token == "" {
		return ErrEmptyOrInvalidBotToken
	}
	return nil
}

func (s *ClientAPI) Do(
	ctx context.Context,
	method string,
	url string,
	requestBody []byte,
	response interface{},
) error {
	request, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	res, err := s.HttpClient.Do(request)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	errResponse, err := hasError(body)
	if err != nil {
		return err
	}

	if errResponse != nil {
		return errors.New(fmt.Sprintf("response ok: %v, error_code:%d, description:%s",
			errResponse.Ok, errResponse.ErrorCode, errResponse.Description))
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	return nil
}

func hasError(r []byte) (*tgbotapi.APIResponse, error) {
	var errResponse *tgbotapi.APIResponse
	err := json.Unmarshal(r, &errResponse)
	if err != nil {
		return nil, err
	}
	if errResponse.ErrorCode != 0 {
		return errResponse, nil
	}
	return nil, nil
}
