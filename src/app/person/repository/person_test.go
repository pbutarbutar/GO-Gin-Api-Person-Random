package repository

import (
	"api-random-user/config"
	"api-random-user/src/domain"
	"api-random-user/src/model/entity"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUseCase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Person Repository")
}

var _ = Describe("Do test person repository", func() {
	Describe("Test Person", func() {
		Context("Access Repository", func() {
			It("http get should error", func() {
				repoPerson := NewPersonUseCase(&config.ApiPersonConfig{
					Url: "",
				})
				_, err := repoPerson.GetPersons(context.TODO())
				Expect(err).NotTo(BeNil())
			})
			It("should not success status is bad request", func() {
				personApiResult := entity.Randomuser{
					Results: []entity.ResultsData{},
				}
				var personExpected []domain.Person

				personData, err := json.Marshal(personApiResult)
				Expect(err).To(BeNil())
				svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte(personData))
				}))
				defer svr.Close()
				repoPerson := NewPersonUseCase(&config.ApiPersonConfig{
					Url: svr.URL,
				})
				personsResult, err := repoPerson.GetPersons(context.TODO())
				Expect(err).To(BeNil())
				Expect(personsResult).To(Equal(personExpected))
			})
			It("should error status is Bad Gateway", func() {
				personApiResult := entity.Randomuser{
					Results: []entity.ResultsData{},
				}
				personData, err := json.Marshal(personApiResult)
				Expect(err).To(BeNil())
				svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusBadGateway)
					w.Write([]byte(personData))
				}))
				defer svr.Close()
				repoPerson := NewPersonUseCase(&config.ApiPersonConfig{
					Url: svr.URL,
				})
				_, err = repoPerson.GetPersons(context.TODO())
				Expect(err).NotTo(BeNil())
			})

			It("should Success", func() {
				personApiResult := entity.Randomuser{
					Results: []entity.ResultsData{
						{
							Gender: "Male",
							Name: entity.NameDetail{
								First: "Parul",
								Last:  "Butar",
							},
							Location: entity.LocationDetail{
								Street: entity.StreetDetail{
									Name: "Depok",
								},
								City: "Jabar",
							},
							Picture: entity.PictureDetail{
								Large: "p.jpg",
							},
						},
					},
				}
				personExpected := []domain.Person{
					{
						Gender:   "Male",
						Fullname: "Parul Butar",
						Address:  "Depok Jabar",
						Picture:  "p.jpg",
					},
				}
				personData, err := json.Marshal(personApiResult)
				Expect(err).To(BeNil())
				svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(personData))
				}))
				defer svr.Close()
				repoPerson := NewPersonUseCase(&config.ApiPersonConfig{
					Url: svr.URL,
				})
				personsResult, err := repoPerson.GetPersons(context.TODO())
				Expect(err).To(BeNil())
				Expect(personsResult).To(Equal(personExpected))
			})
		})
	})
})
