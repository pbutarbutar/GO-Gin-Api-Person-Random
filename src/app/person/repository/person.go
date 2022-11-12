package repository

import (
	"api-random-user/config"
	"api-random-user/src/domain"
	"api-random-user/src/model/entity"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type PersonRepo struct {
	ApiCfg *config.ApiPersonConfig
}

func NewPersonUseCase(apiPersonConfig *config.ApiPersonConfig) domain.PersonRepository {
	return PersonRepo{
		ApiCfg: apiPersonConfig,
	}
}

func (p PersonRepo) GetPersons(ctx context.Context) (result []domain.Person, err error) {
	var randomuser entity.Randomuser
	res, err := http.Get(p.ApiCfg.Url)
	if err != nil {
		return
	}

	if res.StatusCode == 502 {
		err = fmt.Errorf("Error is status 502")
		return
	}

	if res.StatusCode == 400 {
		return
	}

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&randomuser)
	person := domain.Person{}
	for _, rUser := range randomuser.Results {
		fullName := rUser.Name.First + " " + rUser.Name.Last
		address := rUser.Location.Street.Name + " " + rUser.Location.City
		person = domain.Person{
			Gender:   rUser.Gender,
			Fullname: fullName,
			Address:  address,
			Picture:  rUser.Picture.Large,
		}
		result = append(result, person)
	}
	return
}
