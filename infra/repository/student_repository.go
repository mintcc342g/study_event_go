package repository

import (
	"context"
	"study_event_go/domain/entity"
	"study_event_go/domain/interfaces"
	"study_event_go/ent"
	"study_event_go/types"
)

type lilyORMRepository struct {
	conn *ent.Client
}

// NewLilyRepository ...
func NewLilyRepository(conn *ent.Client) interfaces.LilyRepository {
	return &lilyORMRepository{
		conn: conn,
	}
}

func (s *lilyORMRepository) TopClassLilies(ctx context.Context, gardenID types.GardenID, memberCnt int) ([]*entity.Lily, error) {
	return nil, nil
}
