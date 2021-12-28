package repository

import (
	"context"
	"study-event-go/domain/entity"
	"study-event-go/domain/interfaces"
	"study-event-go/ent"
	"study-event-go/types"
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
