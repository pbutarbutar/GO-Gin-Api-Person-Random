package usecase

import (
	"api-random-user/src/domain"
	"context"
)

type PersonUseCase struct {
	PersonRepo domain.PersonRepository
}

func NewPersonUseCase(personRepo domain.PersonRepository) domain.PersonUseCase {
	return PersonUseCase{
		PersonRepo: personRepo,
	}
}

func (p PersonUseCase) GetPersons(ctx context.Context) (result []domain.Person, err error) {
	return p.PersonRepo.GetPersons(ctx)
}
