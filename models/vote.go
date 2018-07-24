package models

import (
	"github.com/ilhammhdd/e-vote/utils"
)

type Vote struct {
	Id                   uint     `json:"id"`
	VoterId              uint     `json:"voter_id"`
	CategoryNominationId uint     `json:"category_nomination_id"`
	CreatedAt            NullTime `json:"created_at"`
	UpdatedAt            NullTime `json:"updated_at"`
}

func (v Vote) PrimaryKey() map[string]uint {
	return map[string]uint{"id": v.Id}
}

func (v Vote) ForeignKey() map[string]uint {
	return map[string]uint{"voter_id": v.VoterId, "category_nomination_id": v.CategoryNominationId}
}

func (v Vote) Columns() map[string]interface{} {
	return map[string]interface{}{"created_at": v.CreatedAt, "updated_at": v.UpdatedAt}
}

func (v Vote) GetAll(stmnt string, params ...interface{}) *[]Vote {
	var votes []Vote
	var vote Vote
	if len(params) == 0 {
		rows, err := utils.DB.Query(stmnt)
		handleError(err)
		for rows.Next() {
			rows.Scan(&vote.Id, &vote.VoterId, &vote.CategoryNominationId, &vote.CreatedAt, &vote.UpdatedAt)
			votes = append(votes, vote)
		}
	} else {
		stmntOut, err := utils.DB.Prepare(stmnt)
		handleError(err)
		rows, err := stmntOut.Query(params...)
		handleError(err)
		for rows.Next() {
			rows.Scan(&vote.Id, &vote.VoterId, &vote.CategoryNominationId, &vote.CreatedAt, &vote.UpdatedAt)
			votes = append(votes, vote)
		}
	}
	return &votes
}
