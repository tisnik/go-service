package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type UserStorage interface {
	ReadUsers() ([]User, error)
	AddUser(name string, surname string) error
	DeleteUser(id string) error
}

type StorageImpl struct {
	connection *sql.DB
}

func NewStorage(dbType, dbName string) (StorageImpl, error) {
	connection, err := sql.Open(dbType, dbName)
	if err != nil {
		return StorageImpl{}, err
	}
	log.Printf("Connected to database %v", connection)
	return StorageImpl{
		connection: connection,
	}, nil
}

func (s StorageImpl) ReadUsers() ([]User, error) {
	rows, err := s.connection.Query("SELECT id, name, surname FROM users ORDER BY id")
	if err != nil {
		return []User{}, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var id int
		var name string
		var surname string

		if err := rows.Scan(&id, &name, &surname); err != nil {
			return users, err
		}
		users = append(users, User{
			ID:      id,
			Name:    name,
			Surname: surname,
		})
	}

	return users, nil
}

func (s StorageImpl) AddUser(name string, surname string) error {
	statement, err := s.connection.Prepare("INSERT INTO users(name, surname) VALUES (?, ?)")
	_, err = statement.Exec(name, surname)
	return err
}

func (s StorageImpl) DeleteUser(id string) error {
	statement, err := s.connection.Prepare("DELETE FROM users WHERE id=?")
	_, err = statement.Exec(id)
	return err
}
