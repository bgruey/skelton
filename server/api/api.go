package api

import (
	"api-server/api/user"
	apiUtils "api-server/api/utils"
	"fmt"
	"log"
	"net/http"
)

type APIHandler struct {
	ServerMux *http.ServeMux
	user      *user.UserHandler
	port      string
}

func NewAPIHandler(port int) *APIHandler {
	ret := new(APIHandler)
	ret.ServerMux = http.NewServeMux()
	ret.port = fmt.Sprintf(":%d", port)

	// Pass server mux to register all paths for sub-handler
	ret.user = user.NewUserHandler(ret.ServerMux)

	ret.ServerMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		apiUtils.EnableCors(&w)
		fmt.Fprintf(w, "Hello world!\n")
	})
	return ret
}

func (a *APIHandler) Run() {
	fmt.Printf("Starting server, listening on port %s\n", a.port)
	log.Fatal(http.ListenAndServe(a.port, a.ServerMux))
}
