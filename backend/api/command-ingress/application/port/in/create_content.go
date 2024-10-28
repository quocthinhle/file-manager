package inputport

import (
	"context"

	"github.com/quocthinhle/file-manager-api/command-ingress/application/domain/entity"
	"github.com/quocthinhle/file-manager-api/command-ingress/application/port/in/command"
)

type CreateContentUseCase interface {
	CreateContent(
		ctx context.Context,
		createContentCommand command.CreateContentCommand,
	) (entity.Content, error)
}
