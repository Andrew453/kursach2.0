package mysql

import (
	"database/sql"

	"golangs.org/kursach2.0/pkg/models"
)

type CommentModel struct {
	DB *sql.DB
}

func (m *CommentModel) Insert(username, content string) (int, error) {
	stmt := `INSERT INTO comments (username, content, created)
    VALUES(?, ?, UTC_TIMESTAMP())`
	result, err := m.DB.Exec(stmt, username, content)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *CommentModel) GetAll() ([]*models.Comment, error) {
	stmt := `SELECT id, username, content, created FROM comments`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []*models.Comment
	for rows.Next() {
		s := &models.Comment{}
		err = rows.Scan(&s.ID, &s.UserName, &s.Content, &s.Created)
		if err != nil {
			return nil, err
		}
		comments = append(comments, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}
