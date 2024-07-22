package user

import (
	"api-server/api/utils"
	"api-server/model"
	"api-server/pkg/sendgrid"
	"api-server/repo"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type UserHandler struct {
	sync.Mutex
	emailClient *sgclient.SendgridClient
	DB          *repo.PostgresClient
}

func NewUserHandler(mux *http.ServeMux) *UserHandler {
	ret := new(UserHandler)

	ret.emailClient = sgclient.NewSendgridClient()
	ret.DB = repo.New()

	mux.Handle("/users", ret)

	user := ret.DB.GetUserByEmail("12p49jpmasf")
	fmt.Printf("%t: %+v\n", user.ID == 0, user)
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
	uuid, err := api.UuidFromUrl("/users/", r)
	if err != nil {
		api.RespondWithJSON(w, http.StatusOK, nil)
		return
	}

	user := uh.DB.GetUserByUuid(uuid)

	api.RespondWithJSON(w, http.StatusOK, user)
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

	defer uh.Unlock()
	uh.Lock()

	var user model.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		api.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	user.Update()
	err = uh.DB.SaveUser(&user)
	if err != nil {
		api.RespondWithJSON(w, http.StatusConflict, err)
	}

	// uh.sendWelcomeEmail(user)
	api.RespondWithJSON(w, http.StatusCreated, user)
}

func (uh *UserHandler) sendWelcomeEmail(user model.User) error {
	email := sgclient.Email{
		ToName:    user.Name,
		ToEmail:   user.Email,
		FromName:  "Admin",
		FromEmail: "admin@arcadetraxx.com",
		Subject:   "Welcome!",
		BodyText:  "Body text.",
		BodyHTML:  "<b>HTML text.</b>",
	}

	return uh.emailClient.SendEmail(email)
}
