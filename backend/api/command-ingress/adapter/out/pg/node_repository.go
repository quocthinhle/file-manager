package pgoutadapter

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	pgdbgenerated "github.com/quocthinhle/file-manager-api/internal/database/postgres/generated"
)

type NodeOutputAdapter struct {
	db    *pgx.Conn
	query *pgdbgenerated.Queries
}

func NewNodeOutputAdapter(db *pgx.Conn) *NodeOutputAdapter {
	return &NodeOutputAdapter{
		db:    db,
		query: pgdbgenerated.New(db),
	}
}

func (n *NodeOutputAdapter) GetHomeDirectory(ctx context.Context, userID uuid.UUID) (any, error) {
	return nil, nil
}
