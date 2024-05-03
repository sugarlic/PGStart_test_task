package postgre

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/pgxpool"
	"github.com/test/pkg/models"
)

type CommandService interface {
	Latest(ctx context.Context) ([]*models.Command, error)
	Get(ctx context.Context, id int) (*models.Command, error)
	Insert(ctx context.Context, title, content string) error
	Update(ctx context.Context, id int, exec_res string) error
	Delete(ctx context.Context, id int) error
}

type CommandModel struct {
	DB *pgxpool.Pool
}

// Insert - Метод для создания новой заметки в базе дынных.
func (m *CommandModel) Insert(ctx context.Context, title, content string) error {
	stmt := `INSERT INTO commands (title, content, exec_res)
	VALUES($1, $2, $3)`

	_, err := m.DB.Exec(ctx, stmt, title, content, "")
	if err != nil {
		return err
	}

	return nil
}

// Get - Метод для возвращения данных заметки по её идентификатору ID.
func (m *CommandModel) Get(ctx context.Context, id int) (*models.Command, error) {
	stmt := `SELECT id, title, content, exec_res FROM commands
    WHERE id = $1`

	row := m.DB.QueryRow(ctx, stmt, id)

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
func (m *CommandModel) Latest(ctx context.Context) ([]*models.Command, error) {
	stmt := `SELECT id, title, content, exec_res FROM commands
	ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(ctx, stmt)
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
func (m *CommandModel) Update(ctx context.Context, id int, exec_res string) error {
	stmt := `UPDATE commands
	SET exec_res = $2
	WHERE id = $1;`

	_, err := m.DB.Exec(ctx, stmt, id, exec_res)
	if err != nil {
		return err
	}

	return nil
}

// Delete удаляет команду из таблицы
func (m *CommandModel) Delete(ctx context.Context, id int) error {
	stmt := `DELETE FROM commands
	WHERE id = $1;`

	_, err := m.DB.Exec(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}
