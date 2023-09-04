package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jolienai/unicorn-factory/internal/unicorn"
	"github.com/jolienai/unicorn-factory/model"
	"github.com/jolienai/unicorn-factory/util"
)

type UnicornHandler struct {
	Repository *unicorn.UnicornRepository
	Factory    *unicorn.UnicornFactory
}

func NewUnicornHandler(repository *unicorn.UnicornRepository, factory *unicorn.UnicornFactory) *UnicornHandler {
	return &UnicornHandler{
		Repository: repository,
		Factory:    factory,
	}
}

func (h *UnicornHandler) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	responseType := r.Header.Get("Response-Type")

	requestId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	unicorns, err := h.Repository.Get(requestId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if len(unicorns) > 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

	if responseType == "minimal" {
		json.NewEncoder(w).Encode(len(unicorns))
	} else {
		json.NewEncoder(w).Encode(unicorns)
	}
}

func (h *UnicornHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request model.UnicornToProduceRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := util.ValidateStruct(request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	amount := request.Amount
	id, err := h.Repository.Create(amount)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ch := make(chan unicorn.UniCorn, amount)
	go h.Factory.Produce(amount, id, ch)
	go h.Repository.Update(amount, id, ch)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(id)
}
