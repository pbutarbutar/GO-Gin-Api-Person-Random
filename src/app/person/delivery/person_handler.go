package delivery

import (
	"api-random-user/src/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PersonHandler struct {
	PersonUcase domain.PersonUseCase
}

func NewPersonHandler(r *gin.Engine, uCaseP domain.PersonUseCase) {
	handler := PersonHandler{
		PersonUcase: uCaseP,
	}
	r.GET("/api/person", handler.FetchPerson)

}

func (p PersonHandler) FetchPerson(c *gin.Context) {
	person, err := p.PersonUcase.GetPersons(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, person)
}
