package data

import (
	"github.com/Bruheem/Portail_de_Reservation/internal/models"
	"github.com/Bruheem/Portail_de_Reservation/internal/validator"
)

type Library struct {
	ID        uint64
	Name      string
	CreatedBy string
}

func ValidateLibrary(v *validator.Validator, library *models.Library) {
	v.Check(library.Name != "", "name", "must not be empty")
	v.Check(library.CreatedBy != "", "createdby", "there must be an author")
}
