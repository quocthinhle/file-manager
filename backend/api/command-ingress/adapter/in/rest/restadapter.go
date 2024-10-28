package restadapter

import (
	"context"

	"github.com/quocthinhle/file-manager-api/command-ingress/application/domain/entity"
	inputport "github.com/quocthinhle/file-manager-api/command-ingress/application/port/in"
)

type FileManagerRestAdapter struct {
	fetchContentUseCase inputport.FetchContentUseCase
}

func (r *FileManagerRestAdapter) CreateContent(ctx context.Context, request CreateContentRequestObject) (CreateContentResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func NewFileManagerRestAdapter(fetchContentUseCase inputport.FetchContentUseCase) *FileManagerRestAdapter {
	return &FileManagerRestAdapter{
		fetchContentUseCase: fetchContentUseCase,
	}
}

func (r *FileManagerRestAdapter) GetHomeDirectory(ctx context.Context, request GetHomeDirectoryRequestObject) (GetHomeDirectoryResponseObject, error) {
	contents, err := r.fetchContentUseCase.FetchRootContents(ctx)
	if err != nil {
		return nil, err
	}

	contentResponses := toContentResponse(contents)

	return GetHomeDirectory200JSONResponse(contentResponses), nil
}

func toContentResponse(contents []entity.Content) []Content {
	var contentResponses []Content
	for _, content := range contents {
		contentResponses = append(contentResponses, Content{
			Id:       content.ID,
			Name:     content.Name,
			ParentID: content.ParentID,
			Type:     content.Type,
		})
	}

	return contentResponses
}
