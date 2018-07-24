package models

import (
	"github.com/ilhammhdd/e-vote/utils"
)

type CategoryNomination struct {
	Id           uint        `json:"id"`
	CategoryId   uint        `json:"category_id"`
	NominationId uint        `json:"nomination_id"`
	CreatedAt    NullTime    `json:"created_at"`
	UpdatedAt    NullTime    `json:"updated_at"`
	Category     *Category   `json:"category,omitempty"`
	Nomination   *Nomination `json:"nomination,omitempty"`
}

func (c CategoryNomination) PrimaryKey() map[string]uint {
	return map[string]uint{"id": c.Id}
}

func (c CategoryNomination) ForeignKey() map[string]uint {
	return map[string]uint{"category_id": c.CategoryId, "nomination_id": c.NominationId}
}

func (c CategoryNomination) Columns() map[string]interface{} {
	return map[string]interface{}{"created_at": c.CreatedAt, "updated_at": c.UpdatedAt}
}

func (c CategoryNomination) GetAll(stmnt string, params ...interface{}) *[]CategoryNomination {
	var categoryNominations []CategoryNomination
	var categoryNomination CategoryNomination
	if len(params) == 0 {
		rows, err := utils.DB.Query(stmnt)
		handleError(err)
		for rows.Next() {
			rows.Scan(&categoryNomination.Id, &categoryNomination.CategoryId, &categoryNomination.NominationId, &categoryNomination.CreatedAt, &categoryNomination.UpdatedAt)
			categoryNominations = append(categoryNominations, categoryNomination)
		}
	} else {
		stmntOut, err := utils.DB.Prepare(stmnt)
		handleError(err)
		rows, err := stmntOut.Query(params...)
		handleError(err)
		for rows.Next() {
			rows.Scan(&categoryNomination.Id, &categoryNomination.CategoryId, &categoryNomination.NominationId, &categoryNomination.CreatedAt, &categoryNomination.UpdatedAt)
			categoryNominations = append(categoryNominations, categoryNomination)
		}
	}
	return &categoryNominations
}
