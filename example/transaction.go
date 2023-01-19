package example

import "database/sql"

type Transaction struct {
	DB *sql.DB
}

func (t *Transaction) ExampleTransaction() error {
	tx, err := t.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO...")
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE...")
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}
