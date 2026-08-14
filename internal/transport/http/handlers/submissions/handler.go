package submissions

import (
	"github.com/go-playground/validator/v10"
	"github.com/labib0x9/sockforces/internal/app/submissions"
)

type Handler struct {
	srv       submissions.Service
	validator *validator.Validate
}

func NewHandler(srv submissions.Service, validator *validator.Validate) *Handler {
	return &Handler{
		srv:       srv,
		validator: validator,
	}
}
