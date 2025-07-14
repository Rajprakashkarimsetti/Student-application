package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rajprakash/student/metrics"
	"github.com/rajprakash/student/models"
	"github.com/rajprakash/student/service"
	"net/http"
	"strconv"
)

type Handler struct {
	handler service.Student
}

func NewHandler(student service.Student) *Handler {
	return &Handler{student}
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	metrics.RequestCount.Inc()
	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err := h.handler.Post(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "student created"})

}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	metrics.RequestCount.Inc()
	students, err := h.handler.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(students)
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	metrics.RequestCount.Inc()

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	student, err := h.handler.GetByID(int(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	metrics.RequestCount.Inc()
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var student models.Student

	student.ID = int(id)
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	res, err := h.handler.Put(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	metrics.RequestCount.Inc()
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.handler.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "student deleted"})
}
