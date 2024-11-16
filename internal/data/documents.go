package data

import (
	"time"

	"github.com/Bruheem/Portail_de_Reservation/internal/models"
	"github.com/Bruheem/Portail_de_Reservation/internal/validator"
)

type Document struct {
	ID             uint64
	Title          string
	Author         string
	YearPublished  time.Time
	ISBN           string
	LibraryID      uint64
	DocumentTypeID uint
}

func ValidateDocument(v *validator.Validator, doc *models.Document) {
	v.Check(doc.Title != "", "title", "must be provided")
	v.Check(doc.Author != "", "author", "must be provided")
	v.Check(doc.YearPublished >= 1000, "yearPublished", "must be more than 1000")
	v.Check(doc.YearPublished <= time.Now().Year(), "yearPublished", "must be less than or equal to the current year")
	v.Check(doc.ISBN != "", "ISBN", "must be provided")
	v.Check(doc.LibraryID != 0, "libraryID", "must be provided")
	v.Check(doc.DocumentTypeID != 0, "documentTypeID", "must be provided")
}
