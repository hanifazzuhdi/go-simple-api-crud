package entity

import "time"

type Customer struct {
	CstId         int
	CstName       string
	CstDob        time.Time
	CstPhonenum   string
	CstEmail      string
	NationalityId int
	Nationality   Nationality
	FamilyList    []FamilyList
}
