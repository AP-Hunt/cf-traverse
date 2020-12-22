package testfixtures

import (
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
	api.pathResponses[path] = bytes
}

func (api *APIServer) Stop() {
	api.httpServer.Close()
}
