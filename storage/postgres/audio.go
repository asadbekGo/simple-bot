package postgres

import (
	"database/sql"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (m *mediaRepo) CreateAudio(message *tgbotapi.Message, filePath string) error {
	result, err := m.db.Exec(`
		INSERT INTO media_audio (
			from_id, first_name, username, file_name, mime_type, 
			file_id, file_unique_id, file_size, duration, path, created_at)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, current_timestamp)
	`,
		&message.From.ID,
		&message.From.FirstName,
		&message.From.UserName,
		&message.Audio.FileName,
		&message.Audio.MimeType,
		&message.Audio.FileID,
		&message.Audio.FileUniqueID,
		&message.Audio.FileSize,
		&message.Audio.Duration,
		&filePath,
	)
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (m *mediaRepo) GetAudios(fromID int64) ([]string, error) {
	rows, err := m.db.Query(`SELECT file_id FROM media_audio WHERE from_id=$1`, fromID)
	if err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	defer rows.Close() // nolint:errcheck

	var fileIDs []string

	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}

		fileIDs = append(fileIDs, id)
	}

	return fileIDs, nil
}
