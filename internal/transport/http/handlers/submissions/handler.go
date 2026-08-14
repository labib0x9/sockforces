package submissions

import (
	"github.com/go-playground/validator/v10"
	"github.com/labib0x9/sockforces/internal/app/submissions"
	"github.com/labib0x9/sockforces/internal/transport/http/middlewares"
)

type Handler struct {
	srv         submissions.Service
	validator   *validator.Validate
	middlewares *middlewares.Middlewares
}

func NewHandler(srv submissions.Service, validator *validator.Validate, middlewares *middlewares.Middlewares) *Handler {
	return &Handler{
		srv:         srv,
		validator:   validator,
		middlewares: middlewares,
	}
}
