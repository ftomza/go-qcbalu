// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/ftomza/go-qcbalu/wallet/repository/ent/predicate"
	"github.com/ftomza/go-qcbalu/wallet/repository/ent/wallet"
	"github.com/google/uuid"
)

// WalletUpdate is the builder for updating Wallet entities.
type WalletUpdate struct {
	config
	hooks      []Hook
	mutation   *WalletMutation
	predicates []predicate.Wallet
}

// Where adds a new predicate for the builder.
func (wu *WalletUpdate) Where(ps ...predicate.Wallet) *WalletUpdate {
	wu.predicates = append(wu.predicates, ps...)
	return wu
}

// SetVersion sets the version field.
func (wu *WalletUpdate) SetVersion(s string) *WalletUpdate {
	wu.mutation.SetVersion(s)
	return wu
}

// SetNillableVersion sets the version field if the given value is not nil.
func (wu *WalletUpdate) SetNillableVersion(s *string) *WalletUpdate {
	if s != nil {
		wu.SetVersion(*s)
	}
	return wu
}

// SetUserID sets the user_id field.
func (wu *WalletUpdate) SetUserID(u uuid.UUID) *WalletUpdate {
	wu.mutation.SetUserID(u)
	return wu
}

// SetLock sets the lock field.
func (wu *WalletUpdate) SetLock(b bool) *WalletUpdate {
	wu.mutation.SetLock(b)
	return wu
}

// SetNillableLock sets the lock field if the given value is not nil.
func (wu *WalletUpdate) SetNillableLock(b *bool) *WalletUpdate {
	if b != nil {
		wu.SetLock(*b)
	}
	return wu
}

// SetBalance sets the balance field.
func (wu *WalletUpdate) SetBalance(i int) *WalletUpdate {
	wu.mutation.ResetBalance()
	wu.mutation.SetBalance(i)
	return wu
}

// SetNillableBalance sets the balance field if the given value is not nil.
func (wu *WalletUpdate) SetNillableBalance(i *int) *WalletUpdate {
	if i != nil {
		wu.SetBalance(*i)
	}
	return wu
}

// AddBalance adds i to balance.
func (wu *WalletUpdate) AddBalance(i int) *WalletUpdate {
	wu.mutation.AddBalance(i)
	return wu
}

// Mutation returns the WalletMutation object of the builder.
func (wu *WalletUpdate) Mutation() *WalletMutation {
	return wu.mutation
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (wu *WalletUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	wu.defaults()
	if len(wu.hooks) == 0 {
		if err = wu.check(); err != nil {
			return 0, err
		}
		affected, err = wu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WalletMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wu.check(); err != nil {
				return 0, err
			}
			wu.mutation = mutation
			affected, err = wu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wu.hooks) - 1; i >= 0; i-- {
			mut = wu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WalletUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WalletUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WalletUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wu *WalletUpdate) defaults() {
	if _, ok := wu.mutation.UpdateTime(); !ok {
		v := wallet.UpdateDefaultUpdateTime()
		wu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wu *WalletUpdate) check() error {
	if v, ok := wu.mutation.Balance(); ok {
		if err := wallet.BalanceValidator(v); err != nil {
			return &ValidationError{Name: "balance", err: fmt.Errorf("ent: validator failed for field \"balance\": %w", err)}
		}
	}
	return nil
}

func (wu *WalletUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   wallet.Table,
			Columns: wallet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: wallet.FieldID,
			},
		},
	}
	if ps := wu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: wallet.FieldUpdateTime,
		})
	}
	if value, ok := wu.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: wallet.FieldVersion,
		})
	}
	if value, ok := wu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: wallet.FieldUserID,
		})
	}
	if value, ok := wu.mutation.Lock(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: wallet.FieldLock,
		})
	}
	if value, ok := wu.mutation.Balance(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: wallet.FieldBalance,
		})
	}
	if value, ok := wu.mutation.AddedBalance(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: wallet.FieldBalance,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wallet.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// WalletUpdateOne is the builder for updating a single Wallet entity.
type WalletUpdateOne struct {
	config
	hooks    []Hook
	mutation *WalletMutation
}

// SetVersion sets the version field.
func (wuo *WalletUpdateOne) SetVersion(s string) *WalletUpdateOne {
	wuo.mutation.SetVersion(s)
	return wuo
}

// SetNillableVersion sets the version field if the given value is not nil.
func (wuo *WalletUpdateOne) SetNillableVersion(s *string) *WalletUpdateOne {
	if s != nil {
		wuo.SetVersion(*s)
	}
	return wuo
}

// SetUserID sets the user_id field.
func (wuo *WalletUpdateOne) SetUserID(u uuid.UUID) *WalletUpdateOne {
	wuo.mutation.SetUserID(u)
	return wuo
}

// SetLock sets the lock field.
func (wuo *WalletUpdateOne) SetLock(b bool) *WalletUpdateOne {
	wuo.mutation.SetLock(b)
	return wuo
}

// SetNillableLock sets the lock field if the given value is not nil.
func (wuo *WalletUpdateOne) SetNillableLock(b *bool) *WalletUpdateOne {
	if b != nil {
		wuo.SetLock(*b)
	}
	return wuo
}

// SetBalance sets the balance field.
func (wuo *WalletUpdateOne) SetBalance(i int) *WalletUpdateOne {
	wuo.mutation.ResetBalance()
	wuo.mutation.SetBalance(i)
	return wuo
}

// SetNillableBalance sets the balance field if the given value is not nil.
func (wuo *WalletUpdateOne) SetNillableBalance(i *int) *WalletUpdateOne {
	if i != nil {
		wuo.SetBalance(*i)
	}
	return wuo
}

// AddBalance adds i to balance.
func (wuo *WalletUpdateOne) AddBalance(i int) *WalletUpdateOne {
	wuo.mutation.AddBalance(i)
	return wuo
}

// Mutation returns the WalletMutation object of the builder.
func (wuo *WalletUpdateOne) Mutation() *WalletMutation {
	return wuo.mutation
}

// Save executes the query and returns the updated entity.
func (wuo *WalletUpdateOne) Save(ctx context.Context) (*Wallet, error) {
	var (
		err  error
		node *Wallet
	)
	wuo.defaults()
	if len(wuo.hooks) == 0 {
		if err = wuo.check(); err != nil {
			return nil, err
		}
		node, err = wuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WalletMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wuo.check(); err != nil {
				return nil, err
			}
			wuo.mutation = mutation
			node, err = wuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wuo.hooks) - 1; i >= 0; i-- {
			mut = wuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WalletUpdateOne) SaveX(ctx context.Context) *Wallet {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WalletUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WalletUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wuo *WalletUpdateOne) defaults() {
	if _, ok := wuo.mutation.UpdateTime(); !ok {
		v := wallet.UpdateDefaultUpdateTime()
		wuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuo *WalletUpdateOne) check() error {
	if v, ok := wuo.mutation.Balance(); ok {
		if err := wallet.BalanceValidator(v); err != nil {
			return &ValidationError{Name: "balance", err: fmt.Errorf("ent: validator failed for field \"balance\": %w", err)}
		}
	}
	return nil
}

func (wuo *WalletUpdateOne) sqlSave(ctx context.Context) (_node *Wallet, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   wallet.Table,
			Columns: wallet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: wallet.FieldID,
			},
		},
	}
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Wallet.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := wuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: wallet.FieldUpdateTime,
		})
	}
	if value, ok := wuo.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: wallet.FieldVersion,
		})
	}
	if value, ok := wuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: wallet.FieldUserID,
		})
	}
	if value, ok := wuo.mutation.Lock(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: wallet.FieldLock,
		})
	}
	if value, ok := wuo.mutation.Balance(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: wallet.FieldBalance,
		})
	}
	if value, ok := wuo.mutation.AddedBalance(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: wallet.FieldBalance,
		})
	}
	_node = &Wallet{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wallet.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
