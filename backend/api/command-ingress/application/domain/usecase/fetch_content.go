package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/quocthinhle/file-manager-api/command-ingress/application/domain/entity"
	outputport "github.com/quocthinhle/file-manager-api/command-ingress/application/port/out"
)

type FetchContentUseCase struct {
	fetchContentOutputPort outputport.FetchContentOutputPort
}

func NewFetchContentUseCase(fetchContentOutputPort outputport.FetchContentOutputPort) *FetchContentUseCase {
	return &FetchContentUseCase{
		fetchContentOutputPort: fetchContentOutputPort,
	}
}

func (u *FetchContentUseCase) FetchRootContents(ctx context.Context) ([]entity.Content, error) {
	contents, err := u.fetchContentOutputPort.FetchRootDirectoryContent(ctx, uuid.MustParse("5ba7229a-1272-4acf-9bbe-0e2da648c55d"))
	if err != nil {
		return nil, err
	}

	return contents, nil
}

func (u *FetchContentUseCase) FetchContent(ctx context.Context, id uuid.UUID) (entity.Content, error) {
	content, err := u.fetchContentOutputPort.FetchContent(ctx, id)
	if err != nil {
		return entity.Content{}, err
	}

	return content, nil
}
