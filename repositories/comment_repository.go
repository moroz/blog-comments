package repositories

import (
	"context"
	"database/sql"

	"github.com/moroz/comments/types"
)

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{db}
}

func (r *CommentRepository) InsertComment(ctx context.Context, params *types.CommentParams) (*types.Comment, error) {
	panic("unimplemented")
}
