package context

import "github.com/taubyte/http/request"

type Context struct {
	req         *request.Request
	variables   map[string]interface{}
	body        []byte
	rawResponse bool
}

type Option func(*Context) error
