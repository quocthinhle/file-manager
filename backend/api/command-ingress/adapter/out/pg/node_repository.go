package pgoutadapter

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/quocthinhle/file-manager-api/command-ingress/application/domain/entity"

	pgdbgenerated "github.com/quocthinhle/file-manager-api/internal/database/postgres/generated"
)

type NodeOutputAdapter struct {
	db    *pgxpool.Pool
	query *pgdbgenerated.Queries
}

func NewNodeOutputAdapter(db *pgxpool.Pool) *NodeOutputAdapter {
	return &NodeOutputAdapter{
		db:    db,
		query: pgdbgenerated.New(db),
	}
}

func (n *NodeOutputAdapter) FetchRootDirectoryContent(
	ctx context.Context,
	ownerID uuid.UUID,
) ([]entity.Content, error) {
	nodes, err := n.query.GetParentNodes(ctx, toPgUUID(ownerID))
	if err != nil {
		return nil, err
	}

	return toContents(nodes), nil
}

func (n *NodeOutputAdapter) Create(
	ctx context.Context,
	content entity.Content,
) (contentCreated entity.Content, err error) {
	var node pgdbgenerated.Node

	node, err = n.query.CreateNode(ctx, pgdbgenerated.CreateNodeParams{
		ID:       toPgUUID(content.ID),
		Type:     content.Type,
		Name:     content.Name,
		ParentID: toPgUUID(content.ParentID),
		OwnerID:  toPgUUID(content.OwnerID),
	})
	if err != nil {
		return
	}

	return toContent(node), nil
}
