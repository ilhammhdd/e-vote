package models

import (
	"github.com/ilhammhdd/e-vote/utils"
)

type Voter struct {
	Id        uint     `json:"id"`
	RegisNum  string   `json:"regis_num"`
	FullName  string   `json:"full_name"`
	CreatedAt NullTime `json:"created_at"`
	UpdatedAt NullTime `json:"updated_at"`
}

func (v Voter) PrimaryKey() map[string]uint {
	return map[string]uint{"id": v.Id}
}

func (v Voter) ForeignKey() map[string]uint {
	return nil
}

func (v Voter) Columns() map[string]interface{} {
	return map[string]interface{}{"regis_num": v.RegisNum, "full_name": v.FullName, "created_at": v.CreatedAt, "updated_at": v.UpdatedAt}
}

func (v Voter) GetAll(stmnt string, params ...interface{}) *[]Voter {
	var voters []Voter
	var voter Voter
	if len(params) == 0 {
		rows, err := utils.DB.Query(stmnt)
		handleError(err)
		for rows.Next() {
			rows.Scan(&voter.Id, &voter.RegisNum, &voter.FullName, &voter.CreatedAt, &voter.UpdatedAt)
			voters = append(voters, voter)
		}
	} else {
		stmntOut, err := utils.DB.Prepare(stmnt)
		handleError(err)
		rows, err := stmntOut.Query(params...)
		handleError(err)
		for rows.Next() {
			rows.Scan(&voter.Id, &voter.RegisNum, &voter.FullName, &voter.CreatedAt, &voter.UpdatedAt)
			voters = append(voters, voter)
		}
	}
	return &voters
}
