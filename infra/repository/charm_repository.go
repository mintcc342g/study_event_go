package repository

import (
	"context"
	"study-event-go/domain/entity"
	"study-event-go/domain/interfaces"
	"study-event-go/ent"
	entCharmCreator "study-event-go/ent/charmcreator"
	entCharmModel "study-event-go/ent/charmmodel"
	"study-event-go/types"

	"github.com/juju/errors"
)

type charmRepository struct {
	conn *ent.Client
}

// NewCharmRepository ...
func NewCharmRepository(conn *ent.Client) interfaces.CharmRepository {
	return &charmRepository{
		conn: conn,
	}
}

func (c *charmRepository) NewCreator(ctx context.Context, creator *entity.CharmCreator) (*entity.CharmCreator, error) {

	entModel, err := c.conn.CharmCreator.
		Create().
		SetName(creator.Name).
		SetType(creator.Type).
		Save(ctx)
	if err != nil {
		// logger
		return nil, errors.New("internal server error")
	}

	creator.ID = entModel.ID
	creator.CreatedAt = entModel.CreatedAt
	creator.UpdatedAt = entModel.UpdatedAt

	return creator, nil
}

func (c *charmRepository) Creator(ctx context.Context, id types.CharmCreatorID) (*entity.CharmCreator, error) {

	entModel, err := c.conn.CharmCreator.
		Query().
		Where(
			entCharmCreator.ID(id),
			entCharmCreator.DeletedAtIsNil()).
		Only(ctx)
	if err != nil {
		// logger
		return nil, errors.NotFoundf("id[%d]", id)
	}

	return &entity.CharmCreator{
		ID:        entModel.ID,
		CreatedAt: entModel.CreatedAt,
		UpdatedAt: entModel.UpdatedAt,
		DeletedAt: entModel.DeletedAt,
		Name:      entModel.Name,
		Type:      entModel.Type,
	}, nil
}

func (c *charmRepository) NewModel(ctx context.Context, model *entity.CharmModel) (*entity.CharmModel, error) {

	entModel, err := c.conn.CharmModel.
		Create().
		SetName(model.Name).
		SetCreatorID(model.CreatorID).
		Save(ctx)
	if err != nil {
		// logger
		return nil, errors.New("internal server error")
	}

	model.ID = entModel.ID
	model.CreatedAt = entModel.CreatedAt
	model.UpdatedAt = entModel.UpdatedAt

	return model, nil
}

func (c *charmRepository) ModelsByIDs(ctx context.Context, ids []types.CharmModelID) ([]*entity.CharmModel, error) {

	entModels, err := c.conn.CharmModel.
		Query().
		Where(
			entCharmModel.IDIn(ids...),
			entCharmModel.DeletedAtIsNil(),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var models []*entity.CharmModel
	for _, v := range entModels {
		models = append(models, &entity.CharmModel{
			ID:        v.ID,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			DeletedAt: v.DeletedAt,
			Name:      v.Name,
			CreatorID: v.CreatorID,
		})
	}

	return models, nil
}

func (c *charmRepository) RegistCharms(ctx context.Context, charms ...*entity.Charm) ([]*entity.Charm, error) {

	bulk := make([]*ent.CharmCreate, len(charms))
	for i, v := range charms {
		bulk[i] = c.conn.Charm.
			Create().
			SetName(v.Name).
			SetModelID(v.ModelID).
			SetOwnerID(v.OwnerID)
	}

	entModels, err := c.conn.Charm.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		// logger
		return nil, errors.New("internal server error")
	}

	for i, v := range entModels {
		charms[i].ID = v.ID
		charms[i].CreatedAt = v.CreatedAt
		charms[i].UpdatedAt = v.UpdatedAt
		charms[i].DeletedAt = v.DeletedAt
		charms[i].Name = v.Name
		charms[i].ModelID = v.ModelID
		charms[i].OwnerID = v.OwnerID
	}

	return charms, nil
}
