package middlewares

import "github.com/labib0x9/sockforces/config"

type Middlewares struct {
	Cnf *config.Config
}

func NewMiddlewares(cnf *config.Config) *Middlewares {
	return &Middlewares{
		Cnf: cnf,
	}
}
