package auth

import (
	"bytes"

	service "github.com/taubyte/go-interfaces/services/http"
)

func AnonymousHandler(ctx service.Context) (interface{}, error) {
	return nil, nil
}

func Scope(scope []string, authHandler service.Handler) service.Handler {
	return func(ctx service.Context) (interface{}, error) {
		auth := []byte(ctx.Request().Header.Get("Authorization"))

		len_auth := len(auth)
		if len_auth > 0 {
			for _, tkn := range AllowedTokenTypes {
				if len_auth > tkn.length+1 && bytes.HasPrefix(auth, tkn.value) {
					ctx.Variables()["Authorization"] = Authorization{
						Type:  tkn.name,
						Token: string(auth[tkn.length+1:]),
						Scope: scope,
					}
				}
			}
		}

		return nil, ctx.HandleAuth(authHandler)
	}
}

func GetAuthorization(c service.Context) *(Authorization) {
	a, ok := c.Variables()["Authorization"]
	if !ok {
		return nil
	}

	v, ok := a.(Authorization)
	if !ok {
		return nil
	}

	return &(v)
}
