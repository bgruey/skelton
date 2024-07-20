package product

import (
	"api-server/api/utils"
	"encoding/json"
	"io"
	"net/http"
	"sync"
)

type ProductHandler struct {
	sync.Mutex
	products Products
}

func NewProductHandler(mux *http.ServeMux) *ProductHandler {
	ret := new(ProductHandler)
	ret.products = Products{
		Product{"Shoes", 25.00},
		Product{"Webcam", 50.00},
		Product{"Mic", 20.00},
	}
	mux.Handle("/products", ret)
	mux.Handle("/products/", ret)
	return ret
}

func (ph *ProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.EnableCors(&w)

	switch r.Method {
	case "GET":
		ph.get(w, r)
	case "POST":
		ph.post(w, r)
	case "PUT", "PATCH":
		ph.put(w, r)
	case "DELETE":
		ph.delete(w, r)
	default:
		api.RespondWithError(w, http.StatusMethodNotAllowed, "invalid method")
	}
}

func (ph *ProductHandler) get(w http.ResponseWriter, r *http.Request) {
	defer ph.Unlock()
	ph.Lock()
	id, err := api.IdFromUrl(r)
	if err != nil {
		api.RespondWithJSON(w, http.StatusOK, ph.products)
		return
	}
	if id >= len(ph.products) || id < 0 {
		api.RespondWithError(w, http.StatusNotFound, "not found")
		return
	}
	api.RespondWithJSON(w, http.StatusOK, ph.products[id])
}

func (ph *ProductHandler) post(w http.ResponseWriter, r *http.Request) {
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
	var product Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		api.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer ph.Unlock()
	ph.Lock()
	ph.products = append(ph.products, product)
	api.RespondWithJSON(w, http.StatusCreated, product)
}

func (ph *ProductHandler) put(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	id, err := api.IdFromUrl(r)
	if err != nil {
		api.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
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
	var product Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		api.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer ph.Unlock()
	ph.Lock()
	if id >= len(ph.products) || id < 0 {
		api.RespondWithError(w, http.StatusNotFound, "not found")
		return
	}
	if product.Name != "" {
		ph.products[id].Name = product.Name
	}
	if product.Price != 0.0 {
		ph.products[id].Price = product.Price
	}
	api.RespondWithJSON(w, http.StatusOK, ph.products[id])
}

func (ph *ProductHandler) delete(w http.ResponseWriter, r *http.Request) {
	id, err := api.IdFromUrl(r)
	if err != nil {
		api.RespondWithError(w, http.StatusNotFound, "not found")
		return
	}
	defer ph.Unlock()
	ph.Lock()
	if id >= len(ph.products) || id < 0 {
		api.RespondWithError(w, http.StatusNotFound, "not found")
		return
	}
	if id < len(ph.products)-1 {
		ph.products[len(ph.products)-1], ph.products[id] = ph.products[id], ph.products[len(ph.products)-1]
	}
	ph.products = ph.products[:len(ph.products)-1]
	api.RespondWithJSON(w, http.StatusNoContent, "")
}
