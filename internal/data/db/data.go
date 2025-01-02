package db

import (
	"database/sql"

	"github.com/cifra-city/location-storage/internal/data/db/dbcore"
)

type Databaser struct {
	Countries Countries
	Cities    Cities
	Districts Districts
	Streets   Streets
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
	queries := dbcore.New(db)
	return &Databaser{
		Countries: NewCountries(queries),
		Cities:    NewCities(queries),
		Districts: NewDistricts(queries),
		Streets:   NewStreets(queries),
	}, nil
}
