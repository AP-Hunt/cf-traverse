package testfixtures

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
)

type APIServer struct {
	pathResponses map[string][]byte
	httpServer    *httptest.Server
	host          string
	port          string
}

func NewAPIServer() *APIServer {
	apiServer := APIServer{
		pathResponses: map[string][]byte{},
	}

	httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		query := r.URL.Query().Encode()

		fullPath := path
		if query != "" {
			unescapedQuery, err := url.QueryUnescape(query)
			if err != nil {
				panic(err)
			}
			fullPath = fullPath+"?"+unescapedQuery
		}
		if response, ok := apiServer.pathResponses[fullPath]; ok {
			w.WriteHeader(http.StatusOK)
			w.Write(response)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))

	apiServer.httpServer = httpServer

	apiServer.PathReturns("/v2/info", []byte(V2Info))
	apiServer.PathReturns("/v3/info", []byte(V3Info))

	return &apiServer
}

func (api *APIServer) ListenerAddr() string {
	return api.httpServer.Listener.Addr().String()
}

func (api *APIServer) PathReturns(path string, bytes []byte) {
	// Convert to a URI type and format correctly, to ensure
	// that the query strings are consistently ordered when
	// a request is parsed in the server
	u, err  := url.ParseRequestURI(path)

	if err != nil {
		panic(fmt.Sprintf("String '%s' is not a valid request URI", path))
	}

	p := u.Path
	query := u.Query().Encode()

	fullPath := p
	if query != "" {
		unescapedQuery, err := url.QueryUnescape(query)
		if err != nil {
			panic(err)
		}
		fullPath = fullPath+"?"+unescapedQuery
	}

	api.pathResponses[fullPath] = bytes
}

func (api *APIServer) Stop() {
	api.httpServer.Close()
}
