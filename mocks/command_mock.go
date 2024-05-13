package mocks

import (
	"context"

	"github.com/test/pkg/models"
)

type MockCommands struct {
	Commands   []*models.Command
	ExpectCall map[string]int
	Err        error
}

func (m *MockCommands) Latest(ctx context.Context) ([]*models.Command, error) {
	m.ExpectCall["Latest"]++
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Commands, nil
}

func (m *MockCommands) Get(ctx context.Context, id int) (*models.Command, error) {
	m.ExpectCall["Get"]++
	if m.Err != nil {
		return nil, m.Err
	}
	for _, cmd := range m.Commands {
		if cmd.ID == id {
			return cmd, nil
		}
	}
	return nil, models.ErrNoRecord
}

func (m *MockCommands) Insert(ctx context.Context, title, content string) error {
	m.ExpectCall["Insert"]++
	if m.Err != nil {
		return m.Err
	}
	// Добавляем команду в "базу данных"
	newCommand := &models.Command{ID: len(m.Commands) + 1, Title: title, Content: content}
	m.Commands = append(m.Commands, newCommand)
	return nil
}

func (m *MockCommands) Update(ctx context.Context, id int, exec_res string) error {
	m.ExpectCall["Update"]++
	if m.Err != nil {
		return m.Err
	}
	for _, cmd := range m.Commands {
		if cmd.ID == id {
			cmd.Exec_res = exec_res
			return nil
		}
	}
	return models.ErrNoRecord
}

func (m *MockCommands) Delete(ctx context.Context, id int) error {
	m.ExpectCall["Delete"]++
	if m.Err != nil {
		return m.Err
	}
	for i, cmd := range m.Commands {
		if cmd.ID == id {
			m.Commands = append(m.Commands[:i], m.Commands[i+1:]...)
			return nil
		}
	}
	return models.ErrNoRecord
}
