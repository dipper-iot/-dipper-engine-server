// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dipper-iot/dipper-engine-server/ent/predicate"
	"github.com/dipper-iot/dipper-engine-server/ent/rulechan"
	"github.com/dipper-iot/dipper-engine-server/ent/rulenode"
)

// RuleNodeUpdate is the builder for updating RuleNode entities.
type RuleNodeUpdate struct {
	config
	hooks    []Hook
	mutation *RuleNodeMutation
}

// Where appends a list predicates to the RuleNodeUpdate builder.
func (rnu *RuleNodeUpdate) Where(ps ...predicate.RuleNode) *RuleNodeUpdate {
	rnu.mutation.Where(ps...)
	return rnu
}

// SetChainID sets the "chain_id" field.
func (rnu *RuleNodeUpdate) SetChainID(u uint64) *RuleNodeUpdate {
	rnu.mutation.SetChainID(u)
	return rnu
}

// SetNillableChainID sets the "chain_id" field if the given value is not nil.
func (rnu *RuleNodeUpdate) SetNillableChainID(u *uint64) *RuleNodeUpdate {
	if u != nil {
		rnu.SetChainID(*u)
	}
	return rnu
}

// ClearChainID clears the value of the "chain_id" field.
func (rnu *RuleNodeUpdate) ClearChainID() *RuleNodeUpdate {
	rnu.mutation.ClearChainID()
	return rnu
}

// SetNodeID sets the "node_id" field.
func (rnu *RuleNodeUpdate) SetNodeID(s string) *RuleNodeUpdate {
	rnu.mutation.SetNodeID(s)
	return rnu
}

// SetRuleID sets the "rule_id" field.
func (rnu *RuleNodeUpdate) SetRuleID(s string) *RuleNodeUpdate {
	rnu.mutation.SetRuleID(s)
	return rnu
}

// SetOption sets the "option" field.
func (rnu *RuleNodeUpdate) SetOption(m map[string]interface{}) *RuleNodeUpdate {
	rnu.mutation.SetOption(m)
	return rnu
}

// SetDebug sets the "debug" field.
func (rnu *RuleNodeUpdate) SetDebug(b bool) *RuleNodeUpdate {
	rnu.mutation.SetDebug(b)
	return rnu
}

// SetNillableDebug sets the "debug" field if the given value is not nil.
func (rnu *RuleNodeUpdate) SetNillableDebug(b *bool) *RuleNodeUpdate {
	if b != nil {
		rnu.SetDebug(*b)
	}
	return rnu
}

// SetEnd sets the "end" field.
func (rnu *RuleNodeUpdate) SetEnd(b bool) *RuleNodeUpdate {
	rnu.mutation.SetEnd(b)
	return rnu
}

// SetNillableEnd sets the "end" field if the given value is not nil.
func (rnu *RuleNodeUpdate) SetNillableEnd(b *bool) *RuleNodeUpdate {
	if b != nil {
		rnu.SetEnd(*b)
	}
	return rnu
}

// SetCreatedAt sets the "created_at" field.
func (rnu *RuleNodeUpdate) SetCreatedAt(t time.Time) *RuleNodeUpdate {
	rnu.mutation.SetCreatedAt(t)
	return rnu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rnu *RuleNodeUpdate) SetNillableCreatedAt(t *time.Time) *RuleNodeUpdate {
	if t != nil {
		rnu.SetCreatedAt(*t)
	}
	return rnu
}

// SetUpdatedAt sets the "updated_at" field.
func (rnu *RuleNodeUpdate) SetUpdatedAt(t time.Time) *RuleNodeUpdate {
	rnu.mutation.SetUpdatedAt(t)
	return rnu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rnu *RuleNodeUpdate) SetNillableUpdatedAt(t *time.Time) *RuleNodeUpdate {
	if t != nil {
		rnu.SetUpdatedAt(*t)
	}
	return rnu
}

// SetChain sets the "chain" edge to the RuleChan entity.
func (rnu *RuleNodeUpdate) SetChain(r *RuleChan) *RuleNodeUpdate {
	return rnu.SetChainID(r.ID)
}

// Mutation returns the RuleNodeMutation object of the builder.
func (rnu *RuleNodeUpdate) Mutation() *RuleNodeMutation {
	return rnu.mutation
}

// ClearChain clears the "chain" edge to the RuleChan entity.
func (rnu *RuleNodeUpdate) ClearChain() *RuleNodeUpdate {
	rnu.mutation.ClearChain()
	return rnu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rnu *RuleNodeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rnu.hooks) == 0 {
		if err = rnu.check(); err != nil {
			return 0, err
		}
		affected, err = rnu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RuleNodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rnu.check(); err != nil {
				return 0, err
			}
			rnu.mutation = mutation
			affected, err = rnu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rnu.hooks) - 1; i >= 0; i-- {
			if rnu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rnu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rnu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (rnu *RuleNodeUpdate) SaveX(ctx context.Context) int {
	affected, err := rnu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rnu *RuleNodeUpdate) Exec(ctx context.Context) error {
	_, err := rnu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rnu *RuleNodeUpdate) ExecX(ctx context.Context) {
	if err := rnu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rnu *RuleNodeUpdate) check() error {
	if v, ok := rnu.mutation.NodeID(); ok {
		if err := rulenode.NodeIDValidator(v); err != nil {
			return &ValidationError{Name: "node_id", err: fmt.Errorf(`ent: validator failed for field "RuleNode.node_id": %w`, err)}
		}
	}
	if v, ok := rnu.mutation.RuleID(); ok {
		if err := rulenode.RuleIDValidator(v); err != nil {
			return &ValidationError{Name: "rule_id", err: fmt.Errorf(`ent: validator failed for field "RuleNode.rule_id": %w`, err)}
		}
	}
	return nil
}

func (rnu *RuleNodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rulenode.Table,
			Columns: rulenode.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: rulenode.FieldID,
			},
		},
	}
	if ps := rnu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rnu.mutation.NodeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rulenode.FieldNodeID,
		})
	}
	if value, ok := rnu.mutation.RuleID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rulenode.FieldRuleID,
		})
	}
	if value, ok := rnu.mutation.Option(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: rulenode.FieldOption,
		})
	}
	if value, ok := rnu.mutation.Debug(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: rulenode.FieldDebug,
		})
	}
	if value, ok := rnu.mutation.End(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: rulenode.FieldEnd,
		})
	}
	if value, ok := rnu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: rulenode.FieldCreatedAt,
		})
	}
	if value, ok := rnu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: rulenode.FieldUpdatedAt,
		})
	}
	if rnu.mutation.ChainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rulenode.ChainTable,
			Columns: []string{rulenode.ChainColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: rulechan.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rnu.mutation.ChainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rulenode.ChainTable,
			Columns: []string{rulenode.ChainColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: rulechan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rnu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rulenode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// RuleNodeUpdateOne is the builder for updating a single RuleNode entity.
type RuleNodeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RuleNodeMutation
}

// SetChainID sets the "chain_id" field.
func (rnuo *RuleNodeUpdateOne) SetChainID(u uint64) *RuleNodeUpdateOne {
	rnuo.mutation.SetChainID(u)
	return rnuo
}

// SetNillableChainID sets the "chain_id" field if the given value is not nil.
func (rnuo *RuleNodeUpdateOne) SetNillableChainID(u *uint64) *RuleNodeUpdateOne {
	if u != nil {
		rnuo.SetChainID(*u)
	}
	return rnuo
}

// ClearChainID clears the value of the "chain_id" field.
func (rnuo *RuleNodeUpdateOne) ClearChainID() *RuleNodeUpdateOne {
	rnuo.mutation.ClearChainID()
	return rnuo
}

// SetNodeID sets the "node_id" field.
func (rnuo *RuleNodeUpdateOne) SetNodeID(s string) *RuleNodeUpdateOne {
	rnuo.mutation.SetNodeID(s)
	return rnuo
}

// SetRuleID sets the "rule_id" field.
func (rnuo *RuleNodeUpdateOne) SetRuleID(s string) *RuleNodeUpdateOne {
	rnuo.mutation.SetRuleID(s)
	return rnuo
}

// SetOption sets the "option" field.
func (rnuo *RuleNodeUpdateOne) SetOption(m map[string]interface{}) *RuleNodeUpdateOne {
	rnuo.mutation.SetOption(m)
	return rnuo
}

// SetDebug sets the "debug" field.
func (rnuo *RuleNodeUpdateOne) SetDebug(b bool) *RuleNodeUpdateOne {
	rnuo.mutation.SetDebug(b)
	return rnuo
}

// SetNillableDebug sets the "debug" field if the given value is not nil.
func (rnuo *RuleNodeUpdateOne) SetNillableDebug(b *bool) *RuleNodeUpdateOne {
	if b != nil {
		rnuo.SetDebug(*b)
	}
	return rnuo
}

// SetEnd sets the "end" field.
func (rnuo *RuleNodeUpdateOne) SetEnd(b bool) *RuleNodeUpdateOne {
	rnuo.mutation.SetEnd(b)
	return rnuo
}

// SetNillableEnd sets the "end" field if the given value is not nil.
func (rnuo *RuleNodeUpdateOne) SetNillableEnd(b *bool) *RuleNodeUpdateOne {
	if b != nil {
		rnuo.SetEnd(*b)
	}
	return rnuo
}

// SetCreatedAt sets the "created_at" field.
func (rnuo *RuleNodeUpdateOne) SetCreatedAt(t time.Time) *RuleNodeUpdateOne {
	rnuo.mutation.SetCreatedAt(t)
	return rnuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rnuo *RuleNodeUpdateOne) SetNillableCreatedAt(t *time.Time) *RuleNodeUpdateOne {
	if t != nil {
		rnuo.SetCreatedAt(*t)
	}
	return rnuo
}

// SetUpdatedAt sets the "updated_at" field.
func (rnuo *RuleNodeUpdateOne) SetUpdatedAt(t time.Time) *RuleNodeUpdateOne {
	rnuo.mutation.SetUpdatedAt(t)
	return rnuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rnuo *RuleNodeUpdateOne) SetNillableUpdatedAt(t *time.Time) *RuleNodeUpdateOne {
	if t != nil {
		rnuo.SetUpdatedAt(*t)
	}
	return rnuo
}

// SetChain sets the "chain" edge to the RuleChan entity.
func (rnuo *RuleNodeUpdateOne) SetChain(r *RuleChan) *RuleNodeUpdateOne {
	return rnuo.SetChainID(r.ID)
}

// Mutation returns the RuleNodeMutation object of the builder.
func (rnuo *RuleNodeUpdateOne) Mutation() *RuleNodeMutation {
	return rnuo.mutation
}

// ClearChain clears the "chain" edge to the RuleChan entity.
func (rnuo *RuleNodeUpdateOne) ClearChain() *RuleNodeUpdateOne {
	rnuo.mutation.ClearChain()
	return rnuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rnuo *RuleNodeUpdateOne) Select(field string, fields ...string) *RuleNodeUpdateOne {
	rnuo.fields = append([]string{field}, fields...)
	return rnuo
}

// Save executes the query and returns the updated RuleNode entity.
func (rnuo *RuleNodeUpdateOne) Save(ctx context.Context) (*RuleNode, error) {
	var (
		err  error
		node *RuleNode
	)
	if len(rnuo.hooks) == 0 {
		if err = rnuo.check(); err != nil {
			return nil, err
		}
		node, err = rnuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RuleNodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rnuo.check(); err != nil {
				return nil, err
			}
			rnuo.mutation = mutation
			node, err = rnuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rnuo.hooks) - 1; i >= 0; i-- {
			if rnuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rnuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rnuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*RuleNode)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RuleNodeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (rnuo *RuleNodeUpdateOne) SaveX(ctx context.Context) *RuleNode {
	node, err := rnuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rnuo *RuleNodeUpdateOne) Exec(ctx context.Context) error {
	_, err := rnuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rnuo *RuleNodeUpdateOne) ExecX(ctx context.Context) {
	if err := rnuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rnuo *RuleNodeUpdateOne) check() error {
	if v, ok := rnuo.mutation.NodeID(); ok {
		if err := rulenode.NodeIDValidator(v); err != nil {
			return &ValidationError{Name: "node_id", err: fmt.Errorf(`ent: validator failed for field "RuleNode.node_id": %w`, err)}
		}
	}
	if v, ok := rnuo.mutation.RuleID(); ok {
		if err := rulenode.RuleIDValidator(v); err != nil {
			return &ValidationError{Name: "rule_id", err: fmt.Errorf(`ent: validator failed for field "RuleNode.rule_id": %w`, err)}
		}
	}
	return nil
}

func (rnuo *RuleNodeUpdateOne) sqlSave(ctx context.Context) (_node *RuleNode, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rulenode.Table,
			Columns: rulenode.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: rulenode.FieldID,
			},
		},
	}
	id, ok := rnuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RuleNode.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rnuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rulenode.FieldID)
		for _, f := range fields {
			if !rulenode.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != rulenode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rnuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rnuo.mutation.NodeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rulenode.FieldNodeID,
		})
	}
	if value, ok := rnuo.mutation.RuleID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rulenode.FieldRuleID,
		})
	}
	if value, ok := rnuo.mutation.Option(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: rulenode.FieldOption,
		})
	}
	if value, ok := rnuo.mutation.Debug(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: rulenode.FieldDebug,
		})
	}
	if value, ok := rnuo.mutation.End(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: rulenode.FieldEnd,
		})
	}
	if value, ok := rnuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: rulenode.FieldCreatedAt,
		})
	}
	if value, ok := rnuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: rulenode.FieldUpdatedAt,
		})
	}
	if rnuo.mutation.ChainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rulenode.ChainTable,
			Columns: []string{rulenode.ChainColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: rulechan.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rnuo.mutation.ChainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rulenode.ChainTable,
			Columns: []string{rulenode.ChainColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: rulechan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &RuleNode{config: rnuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rnuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rulenode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
