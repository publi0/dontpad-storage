package user

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type DataAPI interface {
	SaveFile(file FileData) error
	GetFiles() ([]FileData, error)
	GetFile(id string) (FileData, error)
	DeleteFile(id string) (int64, error)
}

type Data struct {
	db *sql.DB
}

func NewData(db *sql.DB) *Data {
	return &Data{db: db}
}

func (d *Data) SaveFile(file FileData) error {
	_, err := d.db.Exec("INSERT INTO files (id, name, md5, created_at) VALUES (?, ?, ?, ?)",
		file.ID, file.Name, file.MD5, file.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (d *Data) GetFiles() ([]FileData, error) {
	rows, err := d.db.Query("SELECT * FROM files")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []FileData
	for rows.Next() {
		var file FileData
		err := rows.Scan(&file.ID, &file.Name, &file.MD5, &file.CreatedAt)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	return files, nil
}

func (d *Data) GetFile(id string) (FileData, error) {
	var file FileData
	err := d.db.QueryRow("SELECT * FROM files WHERE id = ?", id).
		Scan(&file.ID, &file.Name, &file.MD5, &file.CreatedAt)
	if err != nil {
		return FileData{}, err
	}
	return file, nil
}

func (d *Data) DeleteFile(id string) (int64, error) {
	r, err := d.db.Exec("DELETE FROM files WHERE id = ?", id)
	if err != nil {
		return 0, err
	}
	affected, err := r.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affected, nil
}
