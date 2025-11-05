package handler

import (
	"encoding/json"
	"net/http"
	"simple-api/internal/entity"
	"simple-api/internal/service"
	"strconv"

	"github.com/gorilla/mux"
)

type NationalityHandler struct {
	nationalityService service.NationalityService
}

func NewNationalityHandler(nationalityService service.NationalityService) NationalityHandler {
	return NationalityHandler{
		nationalityService: nationalityService,
	}
}

func (handler NationalityHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	nationalities, err := handler.nationalityService.GetAll(r.Context())
	if err != nil {
		handleAppError(w, err)
		return
	}

	var data []NationalityResponse
	for _, nationality := range nationalities {
		nationalityResponse := NationalityResponse{
			NationalityId:   nationality.NationalityId,
			NationalityName: nationality.NationalityName,
			NationalityCode: nationality.NationalityCode,
		}

		data = append(data, nationalityResponse)
	}

	responseOk(w, 200, data)
}

func (handler NationalityHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	nationality, err := handler.nationalityService.GetById(r.Context(), id)
	if err != nil {
		handleAppError(w, err)
		return
	}

	nationalityResponse := NationalityResponse{
		NationalityId:   nationality.NationalityId,
		NationalityName: nationality.NationalityName,
		NationalityCode: nationality.NationalityCode,
	}

	responseOk(w, 200, nationalityResponse)
}

func (handler NationalityHandler) Create(w http.ResponseWriter, r *http.Request) {
	requestBody := NationalityCreateRequest{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		handleAppError(w, err)
		return
	}

	nationalityEntity := entity.Nationality{
		NationalityName: requestBody.NationalityName,
		NationalityCode: requestBody.NationalityCode,
	}

	nationality, err := handler.nationalityService.Create(r.Context(), nationalityEntity)
	if err != nil {
		handleAppError(w, err)
		return
	}

	nationalityResponse := NationalityResponse{
		NationalityId:   nationality.NationalityId,
		NationalityName: nationality.NationalityName,
		NationalityCode: nationality.NationalityCode,
	}

	responseOk(w, 201, nationalityResponse)
}

func (handler NationalityHandler) Update(w http.ResponseWriter, r *http.Request) {
	requestBody := NationalityUpdateRequest{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		handleAppError(w, err)
		return
	}

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	nationalityEntity := entity.Nationality{
		NationalityId:   id,
		NationalityName: requestBody.NationalityName,
		NationalityCode: requestBody.NationalityCode,
	}

	nationality, err := handler.nationalityService.Update(r.Context(), nationalityEntity)
	if err != nil {
		handleAppError(w, err)
		return
	}

	nationalityResponse := NationalityResponse{
		NationalityId:   nationality.NationalityId,
		NationalityName: nationality.NationalityName,
		NationalityCode: nationality.NationalityCode,
	}

	responseOk(w, 200, nationalityResponse)
}

func (handler NationalityHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	err := handler.nationalityService.Delete(r.Context(), id)
	if err != nil {
		handleAppError(w, err)
	} else {
		responseOk(w, 200, nil)
	}
}
