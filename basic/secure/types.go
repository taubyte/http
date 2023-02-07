package secure

import (
	basicHttp "github.com/taubyte/http/basic"
)

type Service struct {
	*basicHttp.Service
	err  error
	cert []byte
	key  []byte
}

type GetCertificateHandler func(s *Service, args ...interface{}) ([]byte, []byte, error)
