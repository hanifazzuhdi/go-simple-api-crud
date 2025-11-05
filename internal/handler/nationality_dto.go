package handler

type NationalityResponse struct {
	NationalityId   int    `json:"nationality_id"`
	NationalityName string `json:"nationality_name"`
	NationalityCode string `json:"nationality_code"`
}

type NationalityCreateRequest struct {
	NationalityName string `json:"nationality_name"`
	NationalityCode string `json:"nationality_code"`
}

type NationalityUpdateRequest struct {
	NationalityId   int    `json:"nationality_id"`
	NationalityName string `json:"nationality_name"`
	NationalityCode string `json:"nationality_code"`
}
