// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/ftomza/go-qcbalu/wallet/repository/ent/wallet"
	"github.com/google/uuid"
)

// Wallet is the model entity for the Wallet schema.
type Wallet struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Version holds the value of the "version" field.
	Version string `json:"version,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// Lock holds the value of the "lock" field.
	Lock bool `json:"lock,omitempty"`
	// Balance holds the value of the "balance" field.
	Balance int `json:"balance,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Wallet) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},      // id
		&sql.NullTime{},   // create_time
		&sql.NullTime{},   // update_time
		&sql.NullString{}, // version
		&uuid.UUID{},      // user_id
		&sql.NullBool{},   // lock
		&sql.NullInt64{},  // balance
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Wallet fields.
func (w *Wallet) assignValues(values ...interface{}) error {
	if m, n := len(values), len(wallet.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		w.ID = *value
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field create_time", values[0])
	} else if value.Valid {
		w.CreateTime = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field update_time", values[1])
	} else if value.Valid {
		w.UpdateTime = value.Time
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field version", values[2])
	} else if value.Valid {
		w.Version = value.String
	}
	if value, ok := values[3].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field user_id", values[3])
	} else if value != nil {
		w.UserID = *value
	}
	if value, ok := values[4].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field lock", values[4])
	} else if value.Valid {
		w.Lock = value.Bool
	}
	if value, ok := values[5].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field balance", values[5])
	} else if value.Valid {
		w.Balance = int(value.Int64)
	}
	return nil
}

// Update returns a builder for updating this Wallet.
// Note that, you need to call Wallet.Unwrap() before calling this method, if this Wallet
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Wallet) Update() *WalletUpdateOne {
	return (&WalletClient{config: w.config}).UpdateOne(w)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (w *Wallet) Unwrap() *Wallet {
	tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Wallet is not a transactional entity")
	}
	w.config.driver = tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Wallet) String() string {
	var builder strings.Builder
	builder.WriteString("Wallet(")
	builder.WriteString(fmt.Sprintf("id=%v", w.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(w.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(w.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", version=")
	builder.WriteString(w.Version)
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", w.UserID))
	builder.WriteString(", lock=")
	builder.WriteString(fmt.Sprintf("%v", w.Lock))
	builder.WriteString(", balance=")
	builder.WriteString(fmt.Sprintf("%v", w.Balance))
	builder.WriteByte(')')
	return builder.String()
}

// Wallets is a parsable slice of Wallet.
type Wallets []*Wallet

func (w Wallets) config(cfg config) {
	for _i := range w {
		w[_i].config = cfg
	}
}