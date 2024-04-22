package mocks

import "github.com/test/pkg/models"

type MockCommands struct {
	Commands   []*models.Command
	ExpectCall map[string]int
	Err        error // Общая ошибка для всех методов, если нужно
}

func (m *MockCommands) Latest() ([]*models.Command, error) {
	m.ExpectCall["Latest"]++
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Commands, nil
}

func (m *MockCommands) Get(id int) (*models.Command, error) {
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

func (m *MockCommands) Insert(title, content string) error {
	m.ExpectCall["Insert"]++
	if m.Err != nil {
		return m.Err
	}
	// Добавляем команду в "базу данных"
	newCommand := &models.Command{ID: len(m.Commands) + 1, Title: title, Content: content}
	m.Commands = append(m.Commands, newCommand)
	return nil
}

func (m *MockCommands) Update(id int, exec_res string) error {
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

func (m *MockCommands) Delete(id int) error {
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
