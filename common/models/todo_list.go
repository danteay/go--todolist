package models

import "github.com/danteay/go-todolist/common/database"

type TodoList struct {
	ID          string
	Name        string
	Description string
	CreatedAt   string
	UpdatedAt   string
}

func (l TodoList) Store(db database.Database) error {
	query := "INSER INTO lists (name, description) " +
		"VALUES ('" + l.Name + "', '" + l.Description + "')"

	_, err := db.Connection().Exec(query)
	if err != nil {
		return err
	}

	return nil
}
