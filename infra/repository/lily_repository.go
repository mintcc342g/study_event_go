package repository

import (
	"context"
	"study-event-go/domain/entity"
	"study-event-go/domain/interfaces"
	"study-event-go/ent"
	"study-event-go/types"

	"github.com/juju/errors"
)

type lilyRepository struct {
	conn *ent.Client
}

// NewLilyRepository ...
func NewLilyRepository(conn *ent.Client) interfaces.LilyRepository {
	return &lilyRepository{
		conn: conn,
	}
}

func (l *lilyRepository) New(ctx context.Context, lily *entity.Lily) (*entity.Lily, error) {

	entModel, err := l.conn.Lily.
		Create().
		SetFirstName(lily.Name.First).
		SetMiddleName(lily.Name.Middle).
		SetLastName(lily.Name.Last).
		SetBirth(lily.Birth).
		SetRank(lily.Rank).
		SetGardenID(lily.GardenID).
		Save(ctx)
	if err != nil {
		// logger
		return nil, errors.New("internal server error")
	}

	lily.ID = entModel.ID
	lily.CreatedAt = entModel.CreatedAt
	lily.UpdatedAt = entModel.UpdatedAt

	return lily, nil
}

func (l *lilyRepository) TopClassLilies(ctx context.Context, gardenID types.GardenID, memberCnt int) ([]*entity.Lily, error) {
	return nil, nil
}
