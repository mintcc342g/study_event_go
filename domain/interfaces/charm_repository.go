package interfaces

import (
	"context"
	"study-event-go/domain/entity"
	"study-event-go/types"
)

// CharmRepository ...
type CharmRepository interface {
	SaveCreator(ctx context.Context, creator *entity.CharmCreator) (*entity.CharmCreator, error)
	Creator(ctx context.Context, creatorID types.CharmCreatorID) (*entity.CharmCreator, error)

	SaveModel(ctx context.Context, model *entity.CharmModel) (*entity.CharmModel, error)
	ModelsByIDs(ctx context.Context, ids []types.CharmModelID) ([]*entity.CharmModel, error)

	SaveCharms(ctx context.Context, charms ...*entity.Charm) ([]*entity.Charm, error)
}
