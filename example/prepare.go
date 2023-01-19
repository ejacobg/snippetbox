package example

import "database/sql"

type Prepare struct {
	DB         *sql.DB
	InsertStmt *sql.Stmt
}

func NewPrepare(db *sql.DB) (*Prepare, error) {
	insertStmt, err := db.Prepare("INSERT INTO...")
	if err != nil {
		return nil, err
	}

	return &Prepare{DB: db, InsertStmt: insertStmt}, nil
}

func (p *Prepare) Insert( /*args...*/ ) error {
	_, err := p.InsertStmt.Exec( /*args...*/ )
	return err
}

func Usage() {
	db, err := sql.Open("<driver>", "<dsn>")
	if err != nil {
		// errorLog.Fatal(err)
	}
	defer db.Close()

	exampleModel, err := NewPrepare(db)
	if err != nil {
		// errorLog.Fatal(err)
	}

	defer exampleModel.InsertStmt.Close()
}
