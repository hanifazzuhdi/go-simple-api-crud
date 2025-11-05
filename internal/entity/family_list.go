package entity

import "time"

type FamilyList struct {
	FlId       int
	FlRelation string
	FlName     string
	FlDob      time.Time
	CstId      int
}
