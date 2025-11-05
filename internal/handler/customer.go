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

type CustomerHandler struct {
	customerService service.CustomerService
}

func NewCustomerHandler(customerService service.CustomerService) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
	}
}

func (handler CustomerHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	customers, err := handler.customerService.GetAll(r.Context())
	if err != nil {
		handleAppError(w, err)
		return
	}

	var data []CustomerResponse
	for _, customer := range customers {
		customerResponseDto := CustomerResponse{
			CstId:         customer.CstId,
			CstName:       customer.CstName,
			CstDob:        customer.CstDob.Format(layoutDate),
			CstPhonenum:   customer.CstPhonenum,
			CstEmail:      customer.CstEmail,
			NationalityId: customer.NationalityId,
			Nationality: NationalityResponse{
				NationalityId:   customer.NationalityId,
				NationalityName: customer.Nationality.NationalityName,
				NationalityCode: customer.Nationality.NationalityCode,
			},
		}
		data = append(data, customerResponseDto)
	}

	responseOk(w, 200, data)
}

func (handler CustomerHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	customer, err := handler.customerService.GetById(r.Context(), id)
	if err != nil {
		handleAppError(w, err)
		return
	}

	var familyDTOs []FamilyListResponse
	for _, familyList := range customer.FamilyList {
		familyDTOs = append(familyDTOs, FamilyListResponse{
			FlId:       familyList.FlId,
			CstId:      familyList.CstId,
			FlRelation: familyList.FlRelation,
			FlName:     familyList.FlName,
			FlDob:      familyList.FlDob.Format(layoutDate),
		})
	}

	customerResponseDto := CustomerResponse{
		CstId:         customer.CstId,
		CstName:       customer.CstName,
		CstDob:        customer.CstDob.Format(layoutDate),
		CstPhonenum:   customer.CstPhonenum,
		CstEmail:      customer.CstEmail,
		NationalityId: customer.NationalityId,
		Nationality: NationalityResponse{
			NationalityId:   customer.NationalityId,
			NationalityName: customer.Nationality.NationalityName,
			NationalityCode: customer.Nationality.NationalityCode,
		},
		FamilyList: familyDTOs,
	}

	responseOk(w, 200, customerResponseDto)
}

func (handler CustomerHandler) Create(w http.ResponseWriter, r *http.Request) {
	requestBody := CustomerCreateRequest{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		handleAppError(w, err)
		return
	}

	parseCstDob, err := time.Parse(layoutDate, requestBody.CstDob)
	if err != nil {
		handleAppError(w, err)
		return
	}

	customerEntity := entity.Customer{
		CstName:       requestBody.CstName,
		CstDob:        parseCstDob,
		CstPhonenum:   requestBody.CstPhonenum,
		CstEmail:      requestBody.CstEmail,
		NationalityId: requestBody.NationalityId,
	}
	customer, err := handler.customerService.Create(r.Context(), customerEntity)
	if err != nil {
		handleAppError(w, err)
		return
	}

	customerResponseDto := CustomerActionResponse{
		CstId:         customer.CstId,
		CstName:       customer.CstName,
		CstDob:        customer.CstDob.Format(layoutDate),
		CstPhonenum:   customer.CstPhonenum,
		CstEmail:      customer.CstEmail,
		NationalityId: customer.NationalityId,
	}

	responseOk(w, 201, customerResponseDto)
}

func (handler CustomerHandler) Update(w http.ResponseWriter, r *http.Request) {
	requestBody := CustomerUpdateRequest{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		handleAppError(w, err)
		return
	}

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	parseCstDob, err := time.Parse(layoutDate, requestBody.CstDob)
	if err != nil {
		handleAppError(w, err)
		return
	}

	customerEntity := entity.Customer{
		CstId:         id,
		CstName:       requestBody.CstName,
		CstDob:        parseCstDob,
		CstPhonenum:   requestBody.CstPhonenum,
		CstEmail:      requestBody.CstEmail,
		NationalityId: requestBody.NationalityId,
	}
	customer, err := handler.customerService.Update(r.Context(), customerEntity)
	if err != nil {
		handleAppError(w, err)
		return
	}

	customerResponseDto := CustomerActionResponse{
		CstId:         customer.CstId,
		CstName:       customer.CstName,
		CstDob:        customer.CstDob.Format(layoutDate),
		CstPhonenum:   customer.CstPhonenum,
		CstEmail:      customer.CstEmail,
		NationalityId: customer.NationalityId,
	}

	responseOk(w, 200, customerResponseDto)
}

func (handler CustomerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	err := handler.customerService.Delete(r.Context(), id)
	if err != nil {
		handleAppError(w, err)
	} else {
		responseOk(w, 200, nil)
	}
}

func (handler CustomerHandler) SyncCustomerFamilies(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId, _ := strconv.Atoi(vars["id"])

	var requestBody []CustomerSyncFamiliesRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		handleAppError(w, err)
		return
	}

	var entityFamilyLists []entity.FamilyList
	for _, family := range requestBody {
		parseFlDob, errParse := time.Parse(layoutDate, family.FlDob)
		if errParse != nil {
			handleAppError(w, errParse)
			return
		}

		entityFamilyLists = append(entityFamilyLists, entity.FamilyList{
			FlId:       family.FlId,
			FlRelation: family.FlRelation,
			FlName:     family.FlName,
			FlDob:      parseFlDob,
		})
	}

	_, err = handler.customerService.SyncCustomerFamilies(r.Context(), customerId, entityFamilyLists)
	if err != nil {
		handleAppError(w, err)
		return
	}

	responseOk(w, 200, nil)
}
