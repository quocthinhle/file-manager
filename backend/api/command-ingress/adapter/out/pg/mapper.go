package pgoutadapter

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/quocthinhle/file-manager-api/command-ingress/application/domain/entity"
	pgdbgenerated "github.com/quocthinhle/file-manager-api/internal/database/postgres/generated"
)

func toContent(node pgdbgenerated.Node) entity.Content {
	return entity.Content{
		ID:          node.ID.Bytes,
		Name:        node.Name,
		Description: "",
		Type:        node.Type,
		ParentID:    node.ParentID.Bytes,
		OwnerID:     node.OwnerID.Bytes,
		Children:    make([]entity.Content, 0),
	}
}

func toContents(nodes []pgdbgenerated.Node) []entity.Content {
	contents := make([]entity.Content, 0)
	for _, node := range nodes {
		content := toContent(node)
		contents = append(contents, content)
	}

	return contents
}

func toPgUUID(uuid uuid.UUID) pgtype.UUID {
	return pgtype.UUID{
		Bytes: uuid,
		Valid: true,
	}
}
