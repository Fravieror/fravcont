package infrastructure

import (
	"api-test/movement/application"
	"api-test/movement/domain"
	"encoding/json"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	// code := chi.URLParam(r, "code")
	// redirect, err := h.redirectService.Find(code)
	// if err != nil {
	// 	if errors.Cause(err) == ErrRedirectNotFound {
	// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusInternalServerError)
	// 		return
	// 	}
	// }
	// contentType := r.Header.Get("Content-Type")
	// responseBody, err := h.serializer(contentType).Encode(redirect)
	// if err != nil {
	// 	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	// 	return
	// }
	// setupResponse(w, contentType, responseBody, http.StatusOK)
}
func Post(w http.ResponseWriter, r *http.Request) {
	var movement domain.Movement
	bussines := application.NewBussines(&MySQLRepository{})
	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&movement)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bussines.Save(&movement)
	w.Write([]byte(`{"message": "hello world"}`))
}
func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "Pong"}`))
}
