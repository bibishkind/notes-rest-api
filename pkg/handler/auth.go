package handler

import "net/http"

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {

}
