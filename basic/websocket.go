package basic

import (
	"net/http"

	iface "github.com/taubyte/go-interfaces/services/http"
	service "github.com/taubyte/http"
	auth "github.com/taubyte/http/auth"
	"github.com/taubyte/http/context"
	"github.com/taubyte/http/request"
)

func (s *Service) WebSocket(def *iface.WebSocketDefinition) {
	route := s.Router.HandleFunc(def.Path, func(w http.ResponseWriter, r *http.Request) {
		logger.Debugf("[WS] %s", r.RequestURI)

		ctx, err := context.New(&request.Request{ResponseWriter: w, HttpRequest: r}, &def.Vars)
		if err != nil {
			logger.Error(err)
			return
		}
		err = ctx.HandleAuth(auth.Scope(def.Scope, def.Auth.Validator))
		if err != nil {
			// enforceScope will return error to Client
			logger.Error(err)
			return
		}

		defer func() {
			if err := ctx.HandleCleanup(def.Auth.GC); err != nil {
				logger.Errorf("cleanup failed with: %s", err)
			}
		}()

		conn, err := service.Upgrader.Upgrade(w, r, nil)
		if err != nil {
			logger.Errorf("[WS] %s -> %w", r.RequestURI, err)
			return
		}

		handler := def.NewHandler(ctx, conn)
		if handler == nil {
			return
		}

		go handler.In()
		go handler.Out()
	})

	if len(def.Host) > 0 {
		route.Host(def.Host)
	}
}
