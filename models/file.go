package models

import (
	"github.com/ilhammhdd/e-vote/utils"
)

type File struct {
	Id         uint        `json:"id"`
	OriName    NullString  `json:"ori_name"`
	HashName   NullString  `json:"hash_name"`
	Path       NullString  `json:"path"`
	MimeType   NullString  `json:"mime_type"`
	Size       NullFloat64 `json:"size"`
	CreatedAt  NullTime    `json:"created_at"`
	UpdatedAt  NullTime    `json:"updated_at"`
	Categories *[]Category `json:"category,omitempty"`
}

func (f File) PrimaryKey() map[string]uint {
	return map[string]uint{"id": f.Id}
}

func (f File) ForeignKey() map[string]uint {
	return nil
}

func (f File) Columns() map[string]interface{} {
	return map[string]interface{}{"ori_name": f.OriName, "hash_name": f.HashName, "path": f.Path, "mime_type": f.MimeType, "size": f.Size, "created_at": f.CreatedAt, "updated_at": f.UpdatedAt}
}

func (f File) GetAll(stmnt string, params ...interface{}) *[]File {
	var files []File
	var file File
	if len(params) == 0 {
		rows, err := utils.DB.Query(stmnt)
		handleError(err)
		for rows.Next() {
			rows.Scan(&file.Id, &file.OriName, &file.HashName, &file.Path, &file.MimeType, &file.Size, &file.CreatedAt, &file.UpdatedAt)
			files = append(files, file)
		}
	} else {
		stmntOut, err := utils.DB.Prepare(stmnt)
		handleError(err)
		rows, err := stmntOut.Query(params...)
		handleError(err)
		for rows.Next() {
			rows.Scan(&file.Id, &file.OriName, &file.HashName, &file.Path, &file.MimeType, &file.Size, &file.CreatedAt, &file.UpdatedAt)
			files = append(files, file)
		}
	}
	return &files
}
