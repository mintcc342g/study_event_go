package repository

import (
	"context"
	"study_event_go/domain/entity"
	"study_event_go/domain/interfaces"
)

type lilyORMRepository struct {
}

// NewLilyRepository ...
func NewLilyRepository() interfaces.LilyRepository {
	return &lilyORMRepository{}
}

func (s *lilyORMRepository) TopClassLilies(ctx context.Context, gardenID uint64, memberCnt int) ([]*entity.Lily, error) {
	return nil, nil
}
