package handler

import (
	"encoding/json"
	"net/http"
	"simple-api/internal/entity"
	"simple-api/internal/service"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type FamilyListHandler struct {
	familyListService service.FamilyListService
}

func NewFamilyListHandler(familyListService service.FamilyListService) *FamilyListHandler {
	return &FamilyListHandler{
		familyListService: familyListService,
	}
}

func (handler FamilyListHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	familyLists, err := handler.familyListService.GetAll(r.Context())
	if err != nil {
		handleAppError(w, err)
		return
	}

	var data []FamilyListResponse
	for _, familyList := range familyLists {
		familyListResponseDto := FamilyListResponse{
			FlId:       familyList.FlId,
			CstId:      familyList.CstId,
			FlRelation: familyList.FlRelation,
			FlName:     familyList.FlName,
			FlDob:      familyList.FlDob.Format(layoutDate),
		}

		data = append(data, familyListResponseDto)
	}

	responseOk(w, 200, data)
}

func (handler FamilyListHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	familyList, err := handler.familyListService.GetById(r.Context(), id)
	if err != nil {
		handleAppError(w, err)
		return
	}

	familyListResponseDto := FamilyListResponse{
		FlId:       familyList.FlId,
		CstId:      familyList.CstId,
		FlRelation: familyList.FlRelation,
		FlName:     familyList.FlName,
		FlDob:      familyList.FlDob.Format(layoutDate),
	}

	responseOk(w, 200, familyListResponseDto)
}

func (handler FamilyListHandler) Create(w http.ResponseWriter, r *http.Request) {
	requestBody := FamilyListCreateRequest{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		handleAppError(w, err)
		return
	}

	parseFlDob, err := time.Parse(layoutDate, requestBody.FlDob)
	if err != nil {
		handleAppError(w, err)
		return
	}

	familyListEntity := entity.FamilyList{
		CstId:      requestBody.CstId,
		FlRelation: requestBody.FlRelation,
		FlName:     requestBody.FlName,
		FlDob:      parseFlDob,
	}

	familyList, err := handler.familyListService.Create(r.Context(), familyListEntity)
	if err != nil {
		handleAppError(w, err)
		return
	}

	familyListResponseDto := FamilyListResponse{
		FlId:       familyList.FlId,
		CstId:      familyList.CstId,
		FlRelation: familyList.FlRelation,
		FlName:     familyList.FlName,
		FlDob:      familyList.FlDob.Format(layoutDate),
	}

	responseOk(w, 201, familyListResponseDto)
}

func (handler FamilyListHandler) Update(w http.ResponseWriter, r *http.Request) {
	requestBody := FamilyListUpdateRequest{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		handleAppError(w, err)
		return
	}

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	parseFlDob, err := time.Parse(layoutDate, requestBody.FlDob)
	if err != nil {
		handleAppError(w, err)
		return
	}

	familyListEntity := entity.FamilyList{
		FlId:       id,
		CstId:      requestBody.CstId,
		FlRelation: requestBody.FlRelation,
		FlName:     requestBody.FlName,
		FlDob:      parseFlDob,
	}

	familyList, err := handler.familyListService.Update(r.Context(), familyListEntity)
	if err != nil {
		handleAppError(w, err)
		return
	}

	familyListResponseDto := FamilyListResponse{
		FlId:       familyList.FlId,
		CstId:      familyList.CstId,
		FlRelation: familyList.FlRelation,
		FlName:     familyList.FlName,
		FlDob:      familyList.FlDob.Format(layoutDate),
	}

	responseOk(w, 200, familyListResponseDto)
}

func (handler FamilyListHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	err := handler.familyListService.Delete(r.Context(), id)
	if err != nil {
		handleAppError(w, err)
	} else {
		responseOk(w, 200, nil)
	}
}
