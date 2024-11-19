package sqlite

import (
	"database/sql"
	"resume-web-app/models"
)

type ResumeModel struct {
	DB *sql.DB
}

func (model *ResumeModel) Insert(resume models.Resume) error {
	statement := `INSERT INTO resume(name,cellphone,email,webAddress,experience) VALUES(?,?,?,?,?);`
	_, err := model.DB.Exec(statement, resume.Name, resume.Cellphone, resume.Email, resume.WebAddress, resume.Experience)

	return err
}

func (model *ResumeModel) SelectOne(id string) (*models.Resume, error) {
	statement := `SELECT id, name, email, cellphone, webAddress, experience FROM resume WHERE id = ?`
	row := model.DB.QueryRow(statement, id)

	resume := models.Resume{}

	err := row.Scan(&resume.ID, &resume.Name, &resume.Cellphone, &resume.Email, &resume.WebAddress, &resume.Experience)
	if err != nil {
		return nil, err
	}

	return &resume, nil

}

func (model *ResumeModel) All() ([]models.Resume, error) {
	statement := `SELECT id,name,email FROM resume ORDER BY id DESC`
	rows, err := model.DB.Query(statement)

	if err != nil {
		return nil, err
	}

	resumes := []models.Resume{}

	for rows.Next() {
		resume := models.Resume{}
		err := rows.Scan(&resume.ID, &resume.Name, &resume.Email)
		if err != nil {
			return nil, err
		}
		resumes = append(resumes, resume)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return resumes, nil
}
