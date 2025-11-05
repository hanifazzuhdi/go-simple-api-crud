package handler

type FamilyListResponse struct {
	FlId       int    `json:"fl_id"`
	CstId      int    `json:"cst_id"`
	FlRelation string `json:"fl_relation"`
	FlName     string `json:"fl_name"`
	FlDob      string `json:"fl_dob"`
}

type FamilyListCreateRequest struct {
	CstId      int    `json:"cst_id"`
	FlRelation string `json:"fl_relation"`
	FlName     string `json:"fl_name"`
	FlDob      string `json:"fl_dob"`
}

type FamilyListUpdateRequest struct {
	FlId       int    `json:"fl_id"`
	CstId      int    `json:"cst_id"`
	FlRelation string `json:"fl_relation"`
	FlName     string `json:"fl_name"`
	FlDob      string `json:"fl_dob"`
}
