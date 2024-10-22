// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/ftomza/go-qcbalu/wallet/repository/ent/migrate"
	"github.com/google/uuid"

	"github.com/ftomza/go-qcbalu/wallet/repository/ent/wallet"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Wallet is the client for interacting with the Wallet builders.
	Wallet *WalletClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Wallet = NewWalletClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Wallet: NewWalletClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config: cfg,
		Wallet: NewWalletClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Wallet.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Wallet.Use(hooks...)
}

// WalletClient is a client for the Wallet schema.
type WalletClient struct {
	config
}

// NewWalletClient returns a client for the Wallet from the given config.
func NewWalletClient(c config) *WalletClient {
	return &WalletClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `wallet.Hooks(f(g(h())))`.
func (c *WalletClient) Use(hooks ...Hook) {
	c.hooks.Wallet = append(c.hooks.Wallet, hooks...)
}

// Create returns a create builder for Wallet.
func (c *WalletClient) Create() *WalletCreate {
	mutation := newWalletMutation(c.config, OpCreate)
	return &WalletCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Wallet entities.
func (c *WalletClient) CreateBulk(builders ...*WalletCreate) *WalletCreateBulk {
	return &WalletCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Wallet.
func (c *WalletClient) Update() *WalletUpdate {
	mutation := newWalletMutation(c.config, OpUpdate)
	return &WalletUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WalletClient) UpdateOne(w *Wallet) *WalletUpdateOne {
	mutation := newWalletMutation(c.config, OpUpdateOne, withWallet(w))
	return &WalletUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WalletClient) UpdateOneID(id uuid.UUID) *WalletUpdateOne {
	mutation := newWalletMutation(c.config, OpUpdateOne, withWalletID(id))
	return &WalletUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Wallet.
func (c *WalletClient) Delete() *WalletDelete {
	mutation := newWalletMutation(c.config, OpDelete)
	return &WalletDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *WalletClient) DeleteOne(w *Wallet) *WalletDeleteOne {
	return c.DeleteOneID(w.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *WalletClient) DeleteOneID(id uuid.UUID) *WalletDeleteOne {
	builder := c.Delete().Where(wallet.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WalletDeleteOne{builder}
}

// Query returns a query builder for Wallet.
func (c *WalletClient) Query() *WalletQuery {
	return &WalletQuery{config: c.config}
}

// Get returns a Wallet entity by its id.
func (c *WalletClient) Get(ctx context.Context, id uuid.UUID) (*Wallet, error) {
	return c.Query().Where(wallet.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WalletClient) GetX(ctx context.Context, id uuid.UUID) *Wallet {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *WalletClient) Hooks() []Hook {
	hooks := c.hooks.Wallet
	return append(hooks[:len(hooks):len(hooks)], wallet.Hooks[:]...)
}
