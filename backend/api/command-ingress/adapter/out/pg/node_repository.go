package pgoutadapter

import (
	"context"
	"encoding/json"
	"fmt"

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

type ChildContent struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	OwnerID  uuid.UUID `json:"owner_id"`
	ParentID uuid.UUID `json:"parent_id"`
	Type     string    `json:"type"`
}

func (s ChildContent) toContent() entity.Content {
	return entity.Content{
		ID:       s.ID,
		Name:     s.Name,
		OwnerID:  s.OwnerID,
		ParentID: s.ParentID,
		Type:     s.Type,
	}
}

func (n *NodeOutputAdapter) FetchContent(
	ctx context.Context,
	id uuid.UUID,
) (content entity.Content, err error) {
	node, err := n.query.GetNode(ctx, toPgUUID(id))
	if err != nil {
		return
	}

	var childrenContent []ChildContent

	err = json.Unmarshal([]byte(node.Children), &childrenContent)

	if err != nil {
		return entity.Content{}, fmt.Errorf("cannot convert children to []interface{}")
	}

	children := make([]entity.Content, len(childrenContent))
	for i, child := range childrenContent {
		children[i] = child.toContent()
	}

	return entity.Content{
		ID:       node.ID.Bytes,
		Name:     node.Name,
		Type:     node.Type,
		OwnerID:  node.OwnerID.Bytes,
		ParentID: node.ParentID.Bytes,
		Children: children,
	}, nil
}
