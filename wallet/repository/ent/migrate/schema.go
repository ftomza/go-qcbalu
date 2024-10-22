// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// WalletsColumns holds the columns for the "wallets" table.
	WalletsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "version", Type: field.TypeString, Default: "ETShHAXBwlRpYUDkTYRzqduztFehbs"},
		{Name: "user_id", Type: field.TypeUUID, Unique: true},
		{Name: "lock", Type: field.TypeBool},
		{Name: "balance", Type: field.TypeInt},
	}
	// WalletsTable holds the schema information for the "wallets" table.
	WalletsTable = &schema.Table{
		Name:        "wallets",
		Columns:     WalletsColumns,
		PrimaryKey:  []*schema.Column{WalletsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		WalletsTable,
	}
)

func init() {
}
