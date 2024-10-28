package inputport

import (
	"context"

	"github.com/quocthinhle/file-manager-api/command-ingress/application/domain/entity"
)

type FetchContentUseCase interface {
	FetchRootContents(ctx context.Context) ([]entity.Content, error)
}
