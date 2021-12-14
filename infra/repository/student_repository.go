package repository

import (
	"context"
	"study_event_go/domain/entity"
	"study_event_go/domain/interfaces"
	"study_event_go/types"
)

type lilyORMRepository struct {
}

// NewLilyRepository ...
func NewLilyRepository() interfaces.LilyRepository {
	return &lilyORMRepository{}
}

func (s *lilyORMRepository) TopClassLilies(ctx context.Context, gardenID types.GardenID, memberCnt int) ([]*entity.Lily, error) {
	return nil, nil
}
