package domain

//go:generate mockgen -package=mock_persons -source=src/domain/person.go -destination=./src/shared/mocks/mock_person.go

import (
	"context"
)

type Person struct {
	Gender   string `json:"gender"`
	Fullname string `json:"fullname"`
	Address  string `json:"address"`
	Picture  string `json:"picture"`
}

type PersonUseCase interface {
	GetPersons(ctx context.Context) (result []Person, err error)
}

type PersonRepository interface {
	GetPersons(ctx context.Context) (result []Person, err error)
}
