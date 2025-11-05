package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
)

const layoutDate = "2006-01-02"

type ApiResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func responseOk(w http.ResponseWriter, code int, data interface{}) {
	response := ApiResponse{
		Code:   code,
		Status: "success",
		Data:   data,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}
}

func responseError(w http.ResponseWriter, code int, data string) {
	response := ApiResponse{
		Code:   code,
		Status: "error",
		Data:   data,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}
}

func handleAppError(w http.ResponseWriter, err error) {
	if errors.Is(err, pgx.ErrNoRows) {
		responseError(w, http.StatusNotFound, "Resource not found!")
		return
	}

	responseError(w, http.StatusInternalServerError, err.Error())
}
