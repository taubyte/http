package context

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	service "github.com/taubyte/http"
	"github.com/gorilla/mux"
)

func (c *Context) returnData(code int, interfaceData interface{}) error {
	if c.rawResponse == true {
		var err error

		switch data := interfaceData.(type) {
		case []byte:
			c.ctx.ResponseWriter.WriteHeader(code)
			_, err = c.ctx.ResponseWriter.Write(data)
		case string:
			c.ctx.ResponseWriter.WriteHeader(code)
			_, err = c.ctx.ResponseWriter.Write([]byte(data))
		case service.RawData:
			c.ctx.ResponseWriter.Header().Set("Content-Type", data.ContentType)
			c.ctx.ResponseWriter.WriteHeader(code)
			_, err = c.ctx.ResponseWriter.Write(data.Data)
		case service.RawStream:
			c.ctx.ResponseWriter.Header().Set("Content-Type", data.ContentType)
			c.ctx.ResponseWriter.WriteHeader(code)
			rbuf := make([]byte, 1024)
			for {
				var n int
				n, err = data.Stream.Read(rbuf)
				if n == 0 || err != nil {
					if err == io.EOF {
						err = nil
					}
					break
				}

				_, err = c.ctx.ResponseWriter.Write(rbuf[:n])
				if err != nil {
					break
				}
			}
			data.Stream.Close()
		}
		if err != nil {
			return fmt.Errorf("writing raw response failed with: %s", err)
		}
	} else {
		var m string
		m, err := c.formatBody(interfaceData)
		if err != nil {
			c.returnError(http.StatusInternalServerError, err)
			return err
		}
		_, err = c.ctx.ResponseWriter.Write([]byte(m))
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Context) returnError(code int, err error) {
	m, _ := c.formatBody(
		map[string]interface{}{
			"code":  code,
			"error": err.Error(),
		},
	)

	// TODO log error here
	c.ctx.ResponseWriter.Write([]byte(m))
	c.ctx.ResponseWriter.WriteHeader(code)
}

func (c *Context) formatBody(m interface{}) (string, error) {
	out, err := json.Marshal(m)
	if err != nil {
		return "", err
	}

	return string(out), err
}

func (ctx *Context) extractVariables(required []string, optional []string) (map[string]interface{}, error) {
	if len(required)+len(optional) == 0 {
		return map[string]interface{}{}, nil
	}

	var body map[string]interface{}
	if len(ctx.body) > 0 {
		err := json.Unmarshal(ctx.body, &body)
		if err != nil {
			return nil, err
		}
	}

	request := ctx.Request()
	vars := mux.Vars(request)

	xVars := make(map[string]interface{})
	add := func(k string) bool {
		if q := request.URL.Query(); q != nil && q.Has(k) == true {
			xVars[k] = q.Get(k)
			return true
		} else if v, ok := vars[k]; ok == true {
			xVars[k] = v
			return true
		} else if v := request.Header.Get(k); v != "" {
			xVars[k] = v
			return true
		} else if v, ok := body[k]; ok {
			xVars[k] = v
			return true
		}

		return false
	}

	for _, k := range optional {
		add(k)
	}

	for _, k := range required {
		if add(k) == false {
			return nil, fmt.Errorf("Processing `%s`, key `%s` not found!", request.URL, k)
		}
	}

	return xVars, nil
}
