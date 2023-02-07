package context

import (
	"net/http"

	service "github.com/taubyte/http"
	"github.com/taubyte/http/request"
)

func New(ctx *request.Request, vars *service.Variables, options ...Option) (service.Context, error) {
	var err error

	c := &Context{
		ctx: ctx,
	}

	for _, opt := range options {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}

	c.body = ctx.Body()

	c.variables, err = c.extractVariables(vars.Required, vars.Optional)
	if err != nil {
		c.returnError(http.StatusNotAcceptable, err)
		return nil, err
	}

	if c.rawResponse == false {
		c.ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	}

	return c, nil
}
