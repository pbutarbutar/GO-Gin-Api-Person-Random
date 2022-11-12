package delivery

import (
	"api-random-user/src/domain"
	mocksDomain "api-random-user/src/shared/mocks"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
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
		ctrl            *gomock.Controller
		mockPersonUcase *mocksDomain.MockPersonUseCase
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
		mockPersonUcase = mocksDomain.NewMockPersonUseCase(ctrl)
	})

	Describe("Test Usecase Person", func() {
		Context("Access Repository", func() {
			It("Should error response", func() {
				mockPersonUcase.EXPECT().GetPersons(gomock.Any()).Return(
					[]domain.Person{},
					fmt.Errorf("Error"),
				).AnyTimes()

				r := SetUpRouter()
				NewPersonHandler(r, mockPersonUcase)

				req, err := http.NewRequest("GET", "/api/person", nil)
				Expect(err).To(BeNil())

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)

				result, err := ioutil.ReadAll(w.Body)
				Expect(err).To(BeNil())

				personExpected := map[string]interface{}{
					"error": "Error",
				}
				personData, err := json.Marshal(personExpected)
				Expect(err).To(BeNil())
				Expect(result).To(Equal(personData))
			})

			It("Should success response", func() {
				mockPersonUcase.EXPECT().GetPersons(gomock.Any()).Return(
					personExpected,
					nil,
				).AnyTimes()

				r := SetUpRouter()
				NewPersonHandler(r, mockPersonUcase)

				req, err := http.NewRequest("GET", "/api/person", nil)
				Expect(err).To(BeNil())

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)

				result, err := ioutil.ReadAll(w.Body)
				Expect(err).To(BeNil())

				personData, err := json.Marshal(personExpected)
				Expect(err).To(BeNil())
				Expect(result).To(Equal(personData))
			})

		})
	})
})

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}
