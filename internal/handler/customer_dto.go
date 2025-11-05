package handler

type CustomerResponse struct {
	CstId         int    `json:"cst_id"`
	CstName       string `json:"cst_name"`
	CstDob        string `json:"cst_dob"`
	CstPhonenum   string `json:"cst_phonenum"`
	CstEmail      string `json:"cst_email"`
	NationalityId int    `json:"nationality_id"`

	// relation
	Nationality NationalityResponse  `json:"nationality"`
	FamilyList  []FamilyListResponse `json:"family_list"`
}

type CustomerActionResponse struct {
	CstId         int    `json:"cst_id"`
	CstName       string `json:"cst_name"`
	CstDob        string `json:"cst_dob"`
	CstPhonenum   string `json:"cst_phonenum"`
	CstEmail      string `json:"cst_email"`
	NationalityId int    `json:"nationality_id"`
}

type CustomerCreateRequest struct {
	CstName       string `json:"cst_name"`
	CstDob        string `json:"cst_dob"`
	CstPhonenum   string `json:"cst_phonenum"`
	CstEmail      string `json:"cst_email"`
	NationalityId int    `json:"nationality_id"`
}

type CustomerUpdateRequest struct {
	CstId         int    `json:"cst_id"`
	CstName       string `json:"cst_name"`
	CstDob        string `json:"cst_dob"`
	CstPhonenum   string `json:"cst_phonenum"`
	CstEmail      string `json:"cst_email"`
	NationalityId int    `json:"nationality_id"`
}

type CustomerSyncFamiliesRequest struct {
	FlId       int    `json:"fl_id"`
	FlRelation string `json:"fl_relation"`
	FlName     string `json:"fl_name"`
	FlDob      string `json:"fl_dob"`
}
