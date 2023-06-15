package basic

import (
	"net/http"
	"path"

	"github.com/spf13/afero"
	service "github.com/taubyte/go-interfaces/services/http"
	auth "github.com/taubyte/http/auth"
	"github.com/taubyte/http/context"
	"github.com/taubyte/http/request"
)

func fileSystem(def *service.HeadlessAssetsDefinition) afero.Fs {
	if len(def.Directory) > 0 {
		return afero.NewBasePathFs(def.FileSystem, def.Directory)
	}

	return def.FileSystem
}

func setSPAPath(req *http.Request) {
	req.URL.Path = "/"
}

func (s *Service) ServeAssets(def *service.AssetsDefinition) {
	fs := fileSystem(&def.HeadlessAssetsDefinition)
	fileServer := http.FileServer(afero.NewHttpFs(fs))

	route := s.Router.PathPrefix(def.Path).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Debugf("[Asset] %s", r.RequestURI)

		ctx, err := context.New(&request.Request{ResponseWriter: w, HttpRequest: r}, &def.Vars, context.RawResponse())
		if err != nil {
			logger.Error(err)
			return
		}
		if err = ctx.HandleAuth(auth.Scope(def.Scope, def.Auth.Validator)); err != nil {
			logger.Error(err)
			return
		}

		defer func() {
			if err := ctx.HandleCleanup(def.Auth.GC); err != nil {
				logger.Errorf("cleanup failed with: %s", err)
			}
		}()

		// check whether afile exists at the given path
		sts, err := fs.Stat(r.URL.Path)
		if err != nil {
			if !def.SinglePageApplication {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// file does not exist, serve index.html
			setSPAPath(r)
		} else {
			if !sts.IsDir() {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			setSPAPath(r)
		}

		if def.BeforeServe != nil {
			def.BeforeServe(w)
		}

		// otherwise, use http.FileServer to serve the static dir
		fileServer.ServeHTTP(w, r)
	})

	if len(def.Host) > 0 {
		route.Host(def.Host)
	}
}

func (s *Service) LowLevelAssetHandler(def *service.HeadlessAssetsDefinition, w http.ResponseWriter, r *http.Request) error {
	fs := fileSystem(def)
	fileServer := http.FileServer(afero.NewHttpFs(fs))

	w.Header().Del("Content-Type")

	sts, err := fs.Stat(r.URL.Path)
	if err != nil {
		if !def.SinglePageApplication {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return nil
		}

		setSPAPath(r)
	} else {
		if _, err := fs.Stat(path.Join(r.URL.Path, "index.html")); sts.IsDir() && err != nil {
			if !def.SinglePageApplication {
				w.WriteHeader(http.StatusForbidden)
				return nil
			}

			setSPAPath(r)
		}
	}

	if def.BeforeServe != nil {
		def.BeforeServe(w)
	}

	fileServer.ServeHTTP(w, r)
	return nil
}

func (s *Service) AssetHandler(def *service.HeadlessAssetsDefinition, ctx service.Context) (interface{}, error) {
	fs := fileSystem(def)

	fileServer := http.FileServer(afero.NewHttpFs(fs))

	r := ctx.Request()
	w := ctx.Writer()

	w.Header().Del("Content-Type")
	ctx.SetRawResponse(true)

	sts, err := fs.Stat(r.URL.Path)
	if err != nil {
		if !def.SinglePageApplication {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return nil, nil
		}

		r.URL.Path = "/"
	} else {
		if sts.IsDir() {
			if !def.SinglePageApplication {
				w.WriteHeader(http.StatusForbidden)
				return nil, nil
			}
			r.URL.Path = "/"
		}
	}

	if def.BeforeServe != nil {
		def.BeforeServe(w)
	}

	fileServer.ServeHTTP(w, r)
	return nil, nil
}
