package storage

import "fmt"

type Storer interface {
	RetrieveMessage() string
}

type Storage struct{}

func NewStorage() (*Storage, func()) {
	return &Storage{}, func() {
		fmt.Println("Storage cleanup function called.")
	}
}

func (s *Storage) RetrieveMessage() string {
	return "Hello, Wire!"
}
