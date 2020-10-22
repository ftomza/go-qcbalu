package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
	"github.com/ftomza/go-qcbalu/pkg/entplus"
	"github.com/google/uuid"
)

// Wallet holds the schema definition for the Wallet entity.
type Wallet struct {
	ent.Schema
}

// Fields of the Wallet.
func (Wallet) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.UUID{}).Unique(),
		field.Bool("lock").Default(false),
		field.Int("balance").Min(0).Default(0),
	}
}

// Edges of the Wallet.
func (Wallet) Edges() []ent.Edge {
	return nil
}

// Mixin of the Wallet.
func (Wallet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entplus.UUIDMixin{},
		mixin.Time{},
		entplus.VersionMixin{
			LengthVersionField: 30,
		},
	}
}

// Indexes of the Wallet.
func (Wallet) Indexes() []ent.Index {
	return nil
}
