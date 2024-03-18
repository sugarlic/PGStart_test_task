package postgre

import (
	"database/sql"
	"errors"

	"github.com/test/pkg/models"
)

type CommandModel struct {
	DB *sql.DB
}

// Insert - Метод для создания новой заметки в базе дынных.
func (m *CommandModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO commands (title, content)
	VALUES(?, ?)`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Get - Метод для возвращения данных заметки по её идентификатору ID.
func (m *CommandModel) Get(id int) (*models.Command, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
    WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id)

	s := &models.Command{}

	err := row.Scan(&s.ID, &s.Title, &s.Content)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

// Latest - Метод возвращает последние 10 заметок.
func (m *CommandModel) Latest() ([]*models.Command, error) {
	stmt := `SELECT id, title, content FROM commands
	ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var snippets []*models.Command

	for rows.Next() {
		s := &models.Command{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
