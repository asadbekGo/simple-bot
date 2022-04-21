package postgres

import (
	"database/sql"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (m *mediaRepo) CreatePhoto(message *tgbotapi.Message, filePath string) error {
	result, err := m.db.Exec(`
		INSERT INTO media_photo (
			from_id, first_name, username, file_id, 
			file_unique_id, file_size, width, height, path, created_at)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, current_timestamp)
	`,
		&message.From.ID,
		&message.From.FirstName,
		&message.From.UserName,
		&message.Photo[len(message.Photo)-1].FileID,
		&message.Photo[len(message.Photo)-1].FileUniqueID,
		&message.Photo[len(message.Photo)-1].FileSize,
		&message.Photo[len(message.Photo)-1].Width,
		&message.Photo[len(message.Photo)-1].Height,
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

func (m *mediaRepo) GetPhotos(fromID int64) ([]string, error) {
	rows, err := m.db.Query(`SELECT file_id FROM media_photo WHERE from_id=$1`, fromID)
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
