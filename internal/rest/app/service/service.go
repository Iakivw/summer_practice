package service

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"pr1/internal/rest/model"
	"pr1/internal/rest/repository"
	"strconv"
)

type Service struct {
	bicycle repository.BicycleRepository
}

func NewService(bicycle repository.BicycleRepository) *Service {
	return &Service{
		bicycle: bicycle,
	}
}

type CreateRequest struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int64  `json:"price"`
}

func (s *Service) Create(w http.ResponseWriter, r *http.Request) {
	req := new(CreateRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}
	r.Body.Close()

	if err := s.bicycle.Create(r.Context(), model.Bicycle{
		Brand: req.Brand,
		Model: req.Model,
		Price: req.Price,
	}); err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusCreated, nil)
}

type GetResponse struct {
	ID    int64  `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int64  `json:"price"`
}

func (s *Service) Get(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	bicycle, err := s.bicycle.Read(r.Context(), int64(id))
	switch {
	case err == nil:
		response(w, http.StatusOK, bicycle)
	case errors.Is(err, sql.ErrNoRows):
		responseError(w, http.StatusNotFound, err)
	default:
		responseError(w, http.StatusInternalServerError, err)
	}
}

type GetAllResponse struct {
	Results []GetResponse `json:"results"`
}

func (s *Service) GetAll(w http.ResponseWriter, r *http.Request) {
	bicycles, err := s.bicycle.List(r.Context())
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}
	result := make([]GetResponse, len(bicycles))
	for i, bicycle := range bicycles {
		result[i] = GetResponse{
			ID:    bicycle.ID,
			Brand: bicycle.Brand,
			Model: bicycle.Model,
			Price: bicycle.Price,
		}
	}
	response(w, http.StatusOK, GetAllResponse{
		Results: result,
	})
}

type UpdateRequest struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int64  `json:"price"`
}

func (s *Service) Update(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}
	req := new(UpdateRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}
	r.Body.Close()

	if err := s.bicycle.Update(r.Context(), model.Bicycle{
		ID:    int64(id),
		Brand: req.Brand,
		Model: req.Model,
		Price: req.Price,
	}); err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusNoContent, nil)
}

type DeleteRequest struct {
	Id int64 `json:"id"`
}

func (s *Service) Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.bicycle.Delete(r.Context(), int64(id)); err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}
	response(w, http.StatusNoContent, nil)
}

func response(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Println(err)
		}
	}
}

func responseError(w http.ResponseWriter, code int, err error) {
	response(w, code, map[string]string{"error :": err.Error()})
}
