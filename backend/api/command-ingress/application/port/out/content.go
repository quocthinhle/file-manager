package outputport

import (
	"context"
	"github.com/google/uuid"
	"github.com/quocthinhle/file-manager-api/command-ingress/application/domain/entity"
)

type FetchContentOutputPort interface {
	// FetchRootDirectoryContent fetch the root directory content
	FetchRootDirectoryContent(ctx context.Context, ownerID uuid.UUID) ([]entity.Content, error)
}

type CreateContentOutputPort interface {
	// CreateDirectory create a directory with the given name and parentID
	CreateDirectory(ctx context.Context, name string, parentID string) (entity.Content, error)
}
