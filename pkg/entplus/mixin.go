package entplus

import (
	"context"
	"fmt"

	"github.com/facebook/ent"
	"github.com/facebook/ent/examples/start/ent/hook"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
	"github.com/google/uuid"
)

type UUIDMixin struct {
	mixin.Schema
}

func (UUIDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable(),
	}
}

type VersionMixin struct {
	LengthVersionField int
	mixin.Schema
}

func (vm VersionMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("version").
			Default(string(randASCIIBytes(vm.LengthVersionField))),
	}
}

func (vm VersionMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(vm.VersionHook(), ent.OpUpdateOne),
	}
}

func (vm VersionMixin) VersionHook() ent.Hook {
	type OldSetVersion interface {
		SetVersion(string)
		Version() (string, bool)
		OldVersion(context.Context) (string, error)
	}
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			ver, ok := m.(OldSetVersion)
			if !ok {
				return next.Mutate(ctx, m)
			}
			oldV, err := ver.OldVersion(ctx)
			if err != nil {
				return nil, err
			}
			curV, exists := ver.Version()
			if !exists {
				return nil, fmt.Errorf("ent: version field is required in update mutation")
			}
			if oldV != curV {
				return nil, fmt.Errorf("ent: version not valid")
			}
			ver.SetVersion(string(randASCIIBytes(vm.LengthVersionField)))
			return next.Mutate(ctx, m)
		})
	}
}
