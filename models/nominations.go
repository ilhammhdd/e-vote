package models

import (
	"github.com/ilhammhdd/e-vote/utils"
)

type Nomination struct {
	Id          uint       `json:"id"`
	FileId      uint       `json:"file_id"`
	FullName    NullString `json:"full_name"`
	Dob         NullTime   `json:"dob"`
	Poo         NullString `json:"poo"`
	Description NullString `json:"description"`
	CreatedAt   NullTime   `json:"created_at"`
	UpdatedAt   NullTime   `json:"updated_at"`
}

func (n Nomination) PrimaryKey() map[string]uint {
	return map[string]uint{"id": n.Id}
}

func (n Nomination) ForeignKey() map[string]uint {
	return map[string]uint{"file_id": n.FileId}
}

func (n Nomination) Columns() map[string]interface{} {
	return map[string]interface{}{"full_name": n.FullName, "dob": n.Dob, "poo": n.Poo, "description": n.Description, "created_at": n.CreatedAt, "updated_at": n.UpdatedAt}
}

func (n Nomination) GetAll(stmnt string, params ...interface{}) *[]Nomination {
	var nominations []Nomination
	var nomination Nomination
	if len(params) == 0 {
		rows, err := utils.DB.Query(stmnt)
		handleError(err)
		for rows.Next() {
			rows.Scan(&nomination.Id, &nomination.FileId, &nomination.FullName, &nomination.Dob, &nomination.Poo, &nomination.Description, &nomination.CreatedAt, &nomination.UpdatedAt)
			nominations = append(nominations, nomination)
		}
	} else {
		stmntOut, err := utils.DB.Prepare(stmnt)
		handleError(err)
		rows, err := stmntOut.Query(params...)
		handleError(err)
		for rows.Next() {
			rows.Scan(&nomination.Id, &nomination.FileId, &nomination.FullName, &nomination.Dob, &nomination.Poo, &nomination.Description, &nomination.CreatedAt, &nomination.UpdatedAt)
			nominations = append(nominations, nomination)
		}
	}
	return &nominations
}
