package basic

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	service "github.com/taubyte/go-interfaces/services/http"
	auth "github.com/taubyte/http/auth"
	"github.com/taubyte/http/context"
	"github.com/taubyte/http/request"
)

func (s *Service) handleRequest(ctx *request.Request, vars *service.Variables, scope []string, authHandler service.Handler, handler service.Handler, cleanupHandler service.Handler, options ...context.Option) {
	_ctx, err := context.New(ctx, vars, options...)
	if err != nil {
		// New Context will return error to Client
		logger.Error(err)
		return
	}

	err = _ctx.HandleAuth(auth.Scope(scope, authHandler))
	if err != nil {
		// enforceScope will return error to Client
		logger.Error(err)
		return
	}

	defer func() {
		cleanupErr := _ctx.HandleCleanup(cleanupHandler)
		if err != nil {
			logger.Errorf("cleanup failed with: %s", cleanupErr)
		}
	}()

	err = _ctx.HandleWith(handler)
	if err != nil {
		logger.Error(fmt.Errorf("Calling %s failed with %v", ctx.HttpRequest.URL, err))
	}

	logger.Debugf("%s | %v", string(ctx.HttpRequest.RequestURI), _ctx.Variables())
}

func (s *Service) buildRouteFromDef(def *service.RouteDefinition) *mux.Route {
	route := s.Router.HandleFunc(def.Path, func(w http.ResponseWriter, h *http.Request) {
		logger.Debugf("[GET] %s", h.RequestURI)
		options := make([]context.Option, 0)
		if def.RawResponse == true {
			options = append(options, context.RawResponse())
		}
		s.handleRequest(&request.Request{ResponseWriter: w, HttpRequest: h}, &def.Vars, def.Scope, def.Auth.Validator, def.Handler, def.Auth.GC, options...)
	})

	if len(def.Host) > 0 {
		route.Host(def.Host)
	}

	return route
}
