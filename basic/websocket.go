package basic

import (
	"net/http"

	service "github.com/taubyte/http"
	auth "github.com/taubyte/http/auth"
	"github.com/taubyte/http/context"
	"github.com/taubyte/http/request"
)

func (s *Service) WebSocket(def *service.WebSocketDefinition) {
	route := s.Router.HandleFunc(def.Path, func(w http.ResponseWriter, r *http.Request) {
		logger.Debugf("[WS] %s", r.RequestURI)

		_ctx, err := context.New(&request.Request{ResponseWriter: w, HttpRequest: r}, &def.Vars)
		if err != nil {
			// New Context will return error to Client
			logger.Error(err)
			return
		}
		err = _ctx.HandleAuth(auth.Scope(def.Scope, def.Auth.Validator))
		if err != nil {
			// enforceScope will return error to Client
			logger.Error(err)
			return
		}

		defer func() {
			cleanupErr := _ctx.HandleCleanup(def.Auth.GC)
			if err != nil {
				logger.Errorf("cleanup failed with: %s", cleanupErr)
			}
		}()

		conn, err := service.Upgrader.Upgrade(w, r, nil)
		if err != nil {
			logger.Errorf("[WS] %s -> %w", r.RequestURI, err)
			return
		}

		handler := def.NewHandler(_ctx, conn)
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
