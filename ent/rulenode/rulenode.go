// Code generated by ent, DO NOT EDIT.

package rulenode

import (
	"time"
)

const (
	// Label holds the string label denoting the rulenode type in the database.
	Label = "rule_node"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldChainID holds the string denoting the chain_id field in the database.
	FieldChainID = "chain_id"
	// FieldNodeID holds the string denoting the node_id field in the database.
	FieldNodeID = "node_id"
	// FieldRuleID holds the string denoting the rule_id field in the database.
	FieldRuleID = "rule_id"
	// FieldOption holds the string denoting the option field in the database.
	FieldOption = "option"
	// FieldInfinite holds the string denoting the infinite field in the database.
	FieldInfinite = "infinite"
	// FieldDebug holds the string denoting the debug field in the database.
	FieldDebug = "debug"
	// FieldEnd holds the string denoting the end field in the database.
	FieldEnd = "end"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeChain holds the string denoting the chain edge name in mutations.
	EdgeChain = "chain"
	// Table holds the table name of the rulenode in the database.
	Table = "rule_nodes"
	// ChainTable is the table that holds the chain relation/edge.
	ChainTable = "rule_nodes"
	// ChainInverseTable is the table name for the RuleChan entity.
	// It exists in this package in order to avoid circular dependency with the "rulechan" package.
	ChainInverseTable = "rule_chan"
	// ChainColumn is the table column denoting the chain relation/edge.
	ChainColumn = "chain_id"
)

// Columns holds all SQL columns for rulenode fields.
var Columns = []string{
	FieldID,
	FieldChainID,
	FieldNodeID,
	FieldRuleID,
	FieldOption,
	FieldInfinite,
	FieldDebug,
	FieldEnd,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NodeIDValidator is a validator for the "node_id" field. It is called by the builders before save.
	NodeIDValidator func(string) error
	// RuleIDValidator is a validator for the "rule_id" field. It is called by the builders before save.
	RuleIDValidator func(string) error
	// DefaultOption holds the default value on creation for the "option" field.
	DefaultOption map[string]interface{}
	// DefaultInfinite holds the default value on creation for the "infinite" field.
	DefaultInfinite bool
	// DefaultDebug holds the default value on creation for the "debug" field.
	DefaultDebug bool
	// DefaultEnd holds the default value on creation for the "end" field.
	DefaultEnd bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)
