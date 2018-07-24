package models

import (
	"github.com/ilhammhdd/e-vote/utils"
)

type Category struct {
	Id                 uint                  `json:"id"`
	FileId             uint                  `json:"file_id"`
	Name               NullString            `json:"name"`
	CreatedAt          NullTime              `json:"created_at"`
	UpdatedAt          NullTime              `json:"updated_at"`
	File               *File                 `json:"files,omitempty"`
	CategoryNomination *[]CategoryNomination `json:"category_nominations,omitempty"`
}

func (c Category) PrimaryKey() map[string]uint {
	return map[string]uint{"id": c.Id}
}

func (c Category) ForeignKey() map[string]uint {
	return map[string]uint{"file_id": c.FileId}
}

func (c Category) Columns() map[string]interface{} {
	return map[string]interface{}{"name": c.Name, "created_at": c.CreatedAt, "updated_at": c.UpdatedAt}
}

func (c Category) GetAll(stmnt string, params ...interface{}) *[]Category {
	var categories []Category
	var category Category
	if len(params) == 0 {
		rows, err := utils.DB.Query(stmnt)
		handleError(err)
		for rows.Next() {
			rows.Scan(&category.Id, &category.FileId, &category.Name, &category.CreatedAt, &category.UpdatedAt)
			categories = append(categories, category)
		}
	} else {
		stmntOut, err := utils.DB.Prepare(stmnt)
		handleError(err)
		rows, err := stmntOut.Query(params...)
		handleError(err)
		for rows.Next() {
			rows.Scan(&category.Id, &category.FileId, &category.Name, &category.CreatedAt, &category.UpdatedAt)
			categories = append(categories, category)
		}
	}
	return &categories
}
