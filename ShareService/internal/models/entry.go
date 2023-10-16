package models

import "LinkService/internal/identifiers"

type Entry struct {
	ID       string `json: "id"`
	Password string `json: "password"`
}

func NewEntry(password string) Entry {
	return Entry{
		ID:       identifiers.GenerateUUID(),
		Password: password,
	}
}
