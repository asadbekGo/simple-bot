package models

import (
	"io"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SendFile struct {
	Ok bool `json:"ok"`
}

type FileResponse struct {
	Ok     bool          `json:"ok"`
	Result tgbotapi.File `json:"result"`
}

type NewFile struct {
	PathName string
}

func (p NewFile) NeedsUpload() bool {
	return true
}

func (p NewFile) UploadData() (string, io.Reader, error) {
	file, err := os.Open(p.PathName)
	if err != nil {
		return "", nil, err
	}
	return "", file, nil
}

func (p NewFile) SendData() string {
	return ""
}
