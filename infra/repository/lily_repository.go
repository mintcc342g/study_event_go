package repository

import (
	"context"
	"study-event-go/domain/entity"
	"study-event-go/domain/interfaces"
	"study-event-go/domain/vo"
	"study-event-go/ent"
	entLily "study-event-go/ent/lily"
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

func (l *lilyRepository) Save(ctx context.Context, lily *entity.Lily) (*entity.Lily, error) {

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

func (l *lilyRepository) Lily(ctx context.Context, id types.LilyID) (*entity.Lily, error) {

	entModel, err := l.conn.Lily.
		Query().
		Where(
			entLily.ID(id),
			entLily.DeletedAtIsNil()).
		Only(ctx)
	if err != nil {
		// logger
		return nil, errors.NotFoundf("id[%d]", id)
	}

	return &entity.Lily{
		ID:              entModel.ID,
		CreatedAt:       entModel.CreatedAt,
		UpdatedAt:       entModel.UpdatedAt,
		DeletedAt:       entModel.DeletedAt,
		CauseOfDeletion: entModel.CauseOfDeletion,
		Name: &vo.Name{
			First:  entModel.FirstName,
			Middle: entModel.MiddleName,
			Last:   entModel.LastName,
		},
		Birth:    *entModel.Birth,
		Rank:     entModel.Rank,
		GardenID: entModel.GardenID,
		LegionID: entModel.LegionID,
	}, nil
}

func (l *lilyRepository) LiliesByRank(ctx context.Context, gardenID types.GardenID, rank int) ([]*entity.Lily, error) {
	return nil, nil
}
