package inputport

import (
	"context"
	"github.com/google/uuid"

	"github.com/quocthinhle/file-manager-api/command-ingress/application/domain/entity"
)

type FetchContentUseCase interface {
	FetchRootContents(ctx context.Context) ([]entity.Content, error)
	FetchContent(ctx context.Context, id uuid.UUID) (entity.Content, error)
}
