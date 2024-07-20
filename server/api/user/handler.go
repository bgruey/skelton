package user

import (
	"api-server/api/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type UserHandler struct {
	sync.Mutex
	users map[string]User
}

func NewUserHandler(mux *http.ServeMux) *UserHandler {
	ret := new(UserHandler)
	ret.users = make(map[string]User, 0)

	mux.Handle("/users", ret)

	return ret
}

func (uh *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.EnableCors(&w)

	fmt.Printf("Raw path: %s\n", r.URL.RawPath)

	switch r.Method {
	case "GET":
		uh.get(w, r)
	case "POST":
		uh.post(w, r)
	// case "PUT", "PATCH":
	// 	ph.put(w, r)
	// case "DELETE":
	// 	ph.delete(w, r)
	case "OPTIONS":
		uh.options(w, r)
	default:
		api.RespondWithError(w, http.StatusMethodNotAllowed, "invalid method")
	}
}

func (uh *UserHandler) get(w http.ResponseWriter, r *http.Request) {
	defer uh.Unlock()
	uh.Lock()
	id, err := api.IdFromUrl(r)
	if err != nil {
		api.RespondWithJSON(w, http.StatusOK, nil)
		return
	}
	if id >= len(uh.users) || id < 0 {
		api.RespondWithError(w, http.StatusNotFound, "not found")
		return
	}
	api.RespondWithJSON(w, http.StatusOK, nil)
}

func (uh *UserHandler) options(w http.ResponseWriter, r *http.Request) {
	api.RespondWithJSON(w, 200, r.Body.Close())
}
func (uh *UserHandler) post(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		api.RespondWithError(w, http.StatusUnsupportedMediaType, "content type 'application/json' required")
		return
	}

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		api.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer uh.Unlock()
	uh.Lock()

	uh.users[user.Email] = user

	api.RespondWithJSON(w, http.StatusCreated, user)
}
