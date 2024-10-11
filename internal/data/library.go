package data

import validator "github.com/Bruheem/Portail_de_Reservation/internal"

type Library struct {
	ID        uint64
	Name      string
	CreatedBy string
}

func ValidateLibrary(v *validator.Validator, library *Library) {
	v.Check(library.Name != "", "name", "must not be empty")
	v.Check(library.CreatedBy != "", "createdby", "there must be an author")
}
