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
func (m *CommandModel) Insert(title, content string) error {
	stmt := `INSERT INTO commands (title, content, exec_res)
	VALUES($1, $2, $3)`

	_, err := m.DB.Exec(stmt, title, content, "")
	if err != nil {
		return err
	}

	return nil
}

// Get - Метод для возвращения данных заметки по её идентификатору ID.
func (m *CommandModel) Get(id int) (*models.Command, error) {
	stmt := `SELECT id, title, content, exec_res FROM commands
    WHERE id = $1`

	row := m.DB.QueryRow(stmt, id)

	s := &models.Command{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Exec_res)
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
	stmt := `SELECT id, title, content, exec_res FROM commands
	ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var commands []*models.Command

	for rows.Next() {
		s := &models.Command{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Exec_res)
		if err != nil {
			return nil, err
		}
		commands = append(commands, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return commands, nil
}

// Update обновляет вывод команды в таблице
func (m *CommandModel) Update(id int, exec_res string) error {
	stmt := `UPDATE commands
	SET exec_res = $2
	WHERE id = $1;`

	_, err := m.DB.Exec(stmt, id, exec_res)
	if err != nil {
		return err
	}

	return nil
}
