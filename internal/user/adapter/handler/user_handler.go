package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	. "hexrestapi/internal/user/domain"
	. "hexrestapi/internal/user/service"
)

const (
	InternalServerError = "Internal Server Error"
)

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

type UserHandler struct {
	service UserService
}

// Search implements port.UserTransport
func (*UserHandler) Search(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func (h *UserHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}

	user, err := h.service.Load(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusOK, user)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user User
	er1 := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}

	res, er2 := h.service.Create(r.Context(), &user)
	if er2 != nil {
		http.Error(w, er1.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusCreated, res)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var user User
	er1 := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}

	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	if len(user.ID) == 0 {
		user.ID = id
	} else if id != user.ID {
		http.Error(w, "Id not match", http.StatusBadRequest)
		return
	}

	res, er2 := h.service.Update(r.Context(), &user)
	if er2 != nil {
		http.Error(w, er2.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusOK, res)
}

func (h *UserHandler) Patch(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}

	var user User
	er1 := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}
	if len(user.ID) == 0 {
		user.ID = id
	} else if id != user.ID {
		http.Error(w, "Id not match", http.StatusBadRequest)
		return
	}

	// res, er4 := h.service.Patch(r.Context(), &user)
	// if er4 != nil {
	// 	http.Error(w, er4.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// JSON(w, http.StatusOK, res)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	res, err := h.service.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusOK, res)
}

func JSON(w http.ResponseWriter, code int, res interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(res)
}
