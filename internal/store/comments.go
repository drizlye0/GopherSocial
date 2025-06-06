package store

import (
	"context"
	"database/sql"
)

type Comment struct {
	ID         int64  `json:"id"`
	UserID     int64  `json:"user_id"`
	PostID     int64  `json:"post_id"`
	Content    string `json:"content"`
	Created_At string `json:"created_at"`
	User       User   `json:"user"`
}

type CommentStore struct {
	db *sql.DB
}

func (s *CommentStore) Create(ctx context.Context, comment *Comment) error {
	query := `
		INSERT INTO comments(user_id, post_id, content)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		comment.UserID,
		comment.PostID,
		comment.Content,
	).Scan(
		&comment.ID,
		&comment.Created_At,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *CommentStore) GetByPostId(ctx context.Context, postID int64) ([]Comment, error) {
	query := `
		SELECT c.id, c.user_id, c.post_id, c.content, c.created_at, users.username, users.id
		FROM comments c
		JOIN users ON users.id = c.user_id
		WHERE c.post_id = $1
		ORDER BY c.created_at DESC;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var comments []Comment

	for rows.Next() {
		var c Comment
		c.User = User{}

		err := rows.Scan(&c.ID, &c.UserID, &c.PostID, &c.Content, &c.Created_At, &c.User.Username, &c.User.ID)
		if err != nil {
			return nil, err
		}

		comments = append(comments, c)
	}

	return comments, nil
}
