package database

import "database/sql"

type Database interface {
	Connection() *sql.DB
}

type Postgres struct {
	conn *sql.DB
}

func NewPostgres() *Postgres {
	////

	return nil
}

func (p Postgres) Connection() *sql.DB {
	return p.conn
}