package storage

import (
	"app/models"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Like struct {
	db *sql.DB
}

func NewLike(db *sql.DB) *Like {
	return &Like{db: db}
}

func (l *Like) Create(n *models.Like) error {
	_, err := l.db.Exec(
		`
			INSERT INTO
				"like"(
					id, user_id, post_id
				) VALUES (
					$1, $2, $3
				)
			`, uuid.NewString(), n.UserId, n.PostId)
	if err != nil {
		return fmt.Errorf("Like create function error: " + err.Error())
	}
	return nil
}

func (l *Like) Update(n *models.Like, id *string) error {

	_, err := l.db.Exec(
		`
			UPDATE
				"like"
			SET 
				post=$2
			WHERE
				id=$1	
			`, id)
	if err != nil {
		return fmt.Errorf("Like update function error: " + err.Error())
	}

	return nil
}

func (l *Like) DeleteLike(id *string) error {

	res, err := l.db.Exec(
		`
			DELETE FROM
				"like"
			WHERE 
				id=$1
			`, id)
	if err != nil {
		return fmt.Errorf("Like delete function error: " + err.Error())
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (l *Like) GetList(req *models.Like) (*models.List, error) {
	var (
		likeList models.List
		query    = `
		SELECT
		id,
		user_id,
		post_id,
		created_at
		FROM "like"`
		filter = " WHERE 1=1 "
		args   []interface{}
	)
	
	if req.UserId != "" {
		args = append(args, req.UserId)
		filter += " AND user_id=$" + fmt.Sprint(len(args))
	}
	query += filter
	rows, err := l.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error querying likes: %v", err)
	}
	defer func() {
		_ = rows.Close()
		}()
		
		for rows.Next() {
			var like models.Like
			err = rows.Scan(
			&like.Id,
			&like.UserId,
			&like.PostId,
			&like.CreatedAt,
		)
		if err != nil {
			fmt.Println("Scan error:", err)
			return nil, err
		}
		likeList.Users = append(likeList.Users, &like)
	}
	likeList.Count = len(likeList.Users)
	return &likeList, nil
}
