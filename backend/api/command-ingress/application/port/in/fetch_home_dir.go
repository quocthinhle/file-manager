package inputport

import "github.com/quocthinhle/file-manager-api/command-ingress/application/domain/entity"

type FetchUserHomeDirUseCase interface {
	FetchUserHomeDir() ([]entity.Content, error)
}
