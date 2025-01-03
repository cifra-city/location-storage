package db

import (
	"database/sql"

	"github.com/cifra-city/location-storage/internal/data/db/sqlcore"
)

type Databaser struct {
	Cities  Cities
	Streets Streets
}

func NewDBConnection(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func NewDatabaser(url string) (*Databaser, error) {
	db, err := NewDBConnection(url)
	if err != nil {
		return nil, err
	}
	queries := sqlcore.New(db)
	return &Databaser{
		Cities:  NewCities(queries),
		Streets: NewStreets(queries),
	}, nil
}
