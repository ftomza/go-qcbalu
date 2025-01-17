// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/ftomza/go-qcbalu/wallet/repository/ent/wallet"
	"github.com/google/uuid"
)

// WalletCreate is the builder for creating a Wallet entity.
type WalletCreate struct {
	config
	mutation *WalletMutation
	hooks    []Hook
}

// SetCreateTime sets the create_time field.
func (wc *WalletCreate) SetCreateTime(t time.Time) *WalletCreate {
	wc.mutation.SetCreateTime(t)
	return wc
}

// SetNillableCreateTime sets the create_time field if the given value is not nil.
func (wc *WalletCreate) SetNillableCreateTime(t *time.Time) *WalletCreate {
	if t != nil {
		wc.SetCreateTime(*t)
	}
	return wc
}

// SetUpdateTime sets the update_time field.
func (wc *WalletCreate) SetUpdateTime(t time.Time) *WalletCreate {
	wc.mutation.SetUpdateTime(t)
	return wc
}

// SetNillableUpdateTime sets the update_time field if the given value is not nil.
func (wc *WalletCreate) SetNillableUpdateTime(t *time.Time) *WalletCreate {
	if t != nil {
		wc.SetUpdateTime(*t)
	}
	return wc
}

// SetVersion sets the version field.
func (wc *WalletCreate) SetVersion(s string) *WalletCreate {
	wc.mutation.SetVersion(s)
	return wc
}

// SetNillableVersion sets the version field if the given value is not nil.
func (wc *WalletCreate) SetNillableVersion(s *string) *WalletCreate {
	if s != nil {
		wc.SetVersion(*s)
	}
	return wc
}

// SetUserID sets the user_id field.
func (wc *WalletCreate) SetUserID(u uuid.UUID) *WalletCreate {
	wc.mutation.SetUserID(u)
	return wc
}

// SetLock sets the lock field.
func (wc *WalletCreate) SetLock(b bool) *WalletCreate {
	wc.mutation.SetLock(b)
	return wc
}

// SetNillableLock sets the lock field if the given value is not nil.
func (wc *WalletCreate) SetNillableLock(b *bool) *WalletCreate {
	if b != nil {
		wc.SetLock(*b)
	}
	return wc
}

// SetBalance sets the balance field.
func (wc *WalletCreate) SetBalance(i int) *WalletCreate {
	wc.mutation.SetBalance(i)
	return wc
}

// SetNillableBalance sets the balance field if the given value is not nil.
func (wc *WalletCreate) SetNillableBalance(i *int) *WalletCreate {
	if i != nil {
		wc.SetBalance(*i)
	}
	return wc
}

// SetID sets the id field.
func (wc *WalletCreate) SetID(u uuid.UUID) *WalletCreate {
	wc.mutation.SetID(u)
	return wc
}

// Mutation returns the WalletMutation object of the builder.
func (wc *WalletCreate) Mutation() *WalletMutation {
	return wc.mutation
}

// Save creates the Wallet in the database.
func (wc *WalletCreate) Save(ctx context.Context) (*Wallet, error) {
	var (
		err  error
		node *Wallet
	)
	wc.defaults()
	if len(wc.hooks) == 0 {
		if err = wc.check(); err != nil {
			return nil, err
		}
		node, err = wc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WalletMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wc.check(); err != nil {
				return nil, err
			}
			wc.mutation = mutation
			node, err = wc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wc.hooks) - 1; i >= 0; i-- {
			mut = wc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WalletCreate) SaveX(ctx context.Context) *Wallet {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (wc *WalletCreate) defaults() {
	if _, ok := wc.mutation.CreateTime(); !ok {
		v := wallet.DefaultCreateTime()
		wc.mutation.SetCreateTime(v)
	}
	if _, ok := wc.mutation.UpdateTime(); !ok {
		v := wallet.DefaultUpdateTime()
		wc.mutation.SetUpdateTime(v)
	}
	if _, ok := wc.mutation.Version(); !ok {
		v := wallet.DefaultVersion
		wc.mutation.SetVersion(v)
	}
	if _, ok := wc.mutation.Lock(); !ok {
		v := wallet.DefaultLock
		wc.mutation.SetLock(v)
	}
	if _, ok := wc.mutation.Balance(); !ok {
		v := wallet.DefaultBalance
		wc.mutation.SetBalance(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WalletCreate) check() error {
	if _, ok := wc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := wc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := wc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New("ent: missing required field \"version\"")}
	}
	if _, ok := wc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New("ent: missing required field \"user_id\"")}
	}
	if _, ok := wc.mutation.Lock(); !ok {
		return &ValidationError{Name: "lock", err: errors.New("ent: missing required field \"lock\"")}
	}
	if _, ok := wc.mutation.Balance(); !ok {
		return &ValidationError{Name: "balance", err: errors.New("ent: missing required field \"balance\"")}
	}
	if v, ok := wc.mutation.Balance(); ok {
		if err := wallet.BalanceValidator(v); err != nil {
			return &ValidationError{Name: "balance", err: fmt.Errorf("ent: validator failed for field \"balance\": %w", err)}
		}
	}
	return nil
}

func (wc *WalletCreate) sqlSave(ctx context.Context) (*Wallet, error) {
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (wc *WalletCreate) createSpec() (*Wallet, *sqlgraph.CreateSpec) {
	var (
		_node = &Wallet{config: wc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: wallet.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: wallet.FieldID,
			},
		}
	)
	if id, ok := wc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := wc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: wallet.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := wc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: wallet.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := wc.mutation.Version(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: wallet.FieldVersion,
		})
		_node.Version = value
	}
	if value, ok := wc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: wallet.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := wc.mutation.Lock(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: wallet.FieldLock,
		})
		_node.Lock = value
	}
	if value, ok := wc.mutation.Balance(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: wallet.FieldBalance,
		})
		_node.Balance = value
	}
	return _node, _spec
}

// WalletCreateBulk is the builder for creating a bulk of Wallet entities.
type WalletCreateBulk struct {
	config
	builders []*WalletCreate
}

// Save creates the Wallet entities in the database.
func (wcb *WalletCreateBulk) Save(ctx context.Context) ([]*Wallet, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Wallet, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WalletMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (wcb *WalletCreateBulk) SaveX(ctx context.Context) []*Wallet {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
