package repo

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// NewMediaStorageI ...
type MediaStorageI interface {
	CreateAudio(message *tgbotapi.Message, filePath string) error
	CreatePhoto(message *tgbotapi.Message, filePath string) error
	CreateVideo(message *tgbotapi.Message, filePath string) error
	CreateDocument(message *tgbotapi.Message, filePath string) error
	GetAudios(fromID int64) ([]string, error)
	GetPhotos(fromID int64) ([]string, error)
	GetVideos(fromID int64) ([]string, error)
	GetDocuments(fromID int64) ([]string, error)
}
