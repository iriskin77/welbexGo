package repository

import (
	"fmt"

	"github.com/iriskin77/testgo/models"
	"github.com/jmoiron/sqlx"
)

const (
	filesTable = "file"
)

type FileDB struct {
	db *sqlx.DB
}

func NewFileDB(db *sqlx.DB) *FileDB {
	return &FileDB{db: db}
}

func (f *FileDB) UploadFile(file *models.File) int {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, file_path) VALUES ($1, $2) RETURNING id", filesTable)
	row := f.db.QueryRow(query, file.Name, file.File_path)
	if err := row.Scan(&id); err != nil {
		return 0
	}

	return id
}

func (f *FileDB) DownloadFile(id int) (*models.File, error) {

	var fileById models.File

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", filesTable)
	err := f.db.Get(&fileById, query, id)

	return &fileById, err

}
