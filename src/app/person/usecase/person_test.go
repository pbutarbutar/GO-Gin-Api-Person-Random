package usecase

import (
	"api-random-user/src/domain"
	mocksDomain "api-random-user/src/shared/mocks"
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUseCase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Person Use Case")
}

var _ = Describe("Person Use Case", func() {
	var (
		ucPerson       domain.PersonUseCase
		ctrl           *gomock.Controller
		mockPersonRepo *mocksDomain.MockPersonRepository
	)
	personExpected := []domain.Person{
		{
			Gender:   "Male",
			Fullname: "Parul Butar",
			Address:  "Depok Jabar",
			Picture:  "p.jpg",
		},
	}

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockPersonRepo = mocksDomain.NewMockPersonRepository(ctrl)
		ucPerson = NewPersonUseCase(mockPersonRepo)
	})

	Describe("Test Usecase Person", func() {
		Context("Access Repository", func() {
			It("should erro", func() {
				mockPersonRepo.EXPECT().GetPersons(gomock.Any()).Return(
					[]domain.Person{},
					fmt.Errorf("Error"),
				)
				_, err := ucPerson.GetPersons(context.TODO())
				Expect(err).NotTo(BeNil())
			})
			It("should contain data", func() {
				mockPersonRepo.EXPECT().GetPersons(gomock.Any()).Return(
					personExpected,
					nil,
				)
				person, err := ucPerson.GetPersons(context.TODO())
				Expect(err).To(BeNil())
				Expect(person).To(Equal(personExpected))
			})
		})
	})
})
