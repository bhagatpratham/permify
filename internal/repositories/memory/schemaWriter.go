package memory

import (
	"context"
	"errors"

	"github.com/rs/xid"

	"github.com/Permify/permify/internal/repositories"
	db "github.com/Permify/permify/pkg/database/memory"
	base "github.com/Permify/permify/pkg/pb/base/v1"
)

// SchemaWriter - Structure for Schema Writer
type SchemaWriter struct {
	database *db.Memory
}

// NewSchemaWriter creates a new SchemaWriter
func NewSchemaWriter(database *db.Memory) *SchemaWriter {
	return &SchemaWriter{
		database: database,
	}
}

// WriteSchema - Write Schema to repository
func (r *SchemaWriter) WriteSchema(ctx context.Context, definitions []repositories.SchemaDefinition) (string, error) {
	id := xid.New()
	var err error
	txn := r.database.DB.Txn(true)
	defer txn.Abort()
	for _, definition := range definitions {
		definition.Version = id.String()
		if err = txn.Insert(SchemaDefinitionTable, definition); err != nil {
			return "", errors.New(base.ErrorCode_ERROR_CODE_EXECUTION.String())
		}
	}
	txn.Commit()
	return id.String(), nil
}
