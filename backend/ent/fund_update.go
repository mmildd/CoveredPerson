// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/mmildd_s/app/ent/coveredperson"
	"github.com/mmildd_s/app/ent/fund"
	"github.com/mmildd_s/app/ent/predicate"
)

// FundUpdate is the builder for updating Fund entities.
type FundUpdate struct {
	config
	hooks      []Hook
	mutation   *FundMutation
	predicates []predicate.Fund
}

// Where adds a new predicate for the builder.
func (fu *FundUpdate) Where(ps ...predicate.Fund) *FundUpdate {
	fu.predicates = append(fu.predicates, ps...)
	return fu
}

// SetFundName sets the Fund_Name field.
func (fu *FundUpdate) SetFundName(s string) *FundUpdate {
	fu.mutation.SetFundName(s)
	return fu
}

// AddFundCoveredPersonIDs adds the Fund_CoveredPerson edge to CoveredPerson by ids.
func (fu *FundUpdate) AddFundCoveredPersonIDs(ids ...int) *FundUpdate {
	fu.mutation.AddFundCoveredPersonIDs(ids...)
	return fu
}

// AddFundCoveredPerson adds the Fund_CoveredPerson edges to CoveredPerson.
func (fu *FundUpdate) AddFundCoveredPerson(c ...*CoveredPerson) *FundUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return fu.AddFundCoveredPersonIDs(ids...)
}

// Mutation returns the FundMutation object of the builder.
func (fu *FundUpdate) Mutation() *FundMutation {
	return fu.mutation
}

// RemoveFundCoveredPersonIDs removes the Fund_CoveredPerson edge to CoveredPerson by ids.
func (fu *FundUpdate) RemoveFundCoveredPersonIDs(ids ...int) *FundUpdate {
	fu.mutation.RemoveFundCoveredPersonIDs(ids...)
	return fu
}

// RemoveFundCoveredPerson removes Fund_CoveredPerson edges to CoveredPerson.
func (fu *FundUpdate) RemoveFundCoveredPerson(c ...*CoveredPerson) *FundUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return fu.RemoveFundCoveredPersonIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (fu *FundUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := fu.mutation.FundName(); ok {
		if err := fund.FundNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "Fund_Name", err: fmt.Errorf("ent: validator failed for field \"Fund_Name\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(fu.hooks) == 0 {
		affected, err = fu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FundMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fu.mutation = mutation
			affected, err = fu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fu.hooks) - 1; i >= 0; i-- {
			mut = fu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FundUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FundUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FundUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fu *FundUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fund.Table,
			Columns: fund.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: fund.FieldID,
			},
		},
	}
	if ps := fu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.FundName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fund.FieldFundName,
		})
	}
	if nodes := fu.mutation.RemovedFundCoveredPersonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fund.FundCoveredPersonTable,
			Columns: []string{fund.FundCoveredPersonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coveredperson.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.FundCoveredPersonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fund.FundCoveredPersonTable,
			Columns: []string{fund.FundCoveredPersonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coveredperson.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fund.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// FundUpdateOne is the builder for updating a single Fund entity.
type FundUpdateOne struct {
	config
	hooks    []Hook
	mutation *FundMutation
}

// SetFundName sets the Fund_Name field.
func (fuo *FundUpdateOne) SetFundName(s string) *FundUpdateOne {
	fuo.mutation.SetFundName(s)
	return fuo
}

// AddFundCoveredPersonIDs adds the Fund_CoveredPerson edge to CoveredPerson by ids.
func (fuo *FundUpdateOne) AddFundCoveredPersonIDs(ids ...int) *FundUpdateOne {
	fuo.mutation.AddFundCoveredPersonIDs(ids...)
	return fuo
}

// AddFundCoveredPerson adds the Fund_CoveredPerson edges to CoveredPerson.
func (fuo *FundUpdateOne) AddFundCoveredPerson(c ...*CoveredPerson) *FundUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return fuo.AddFundCoveredPersonIDs(ids...)
}

// Mutation returns the FundMutation object of the builder.
func (fuo *FundUpdateOne) Mutation() *FundMutation {
	return fuo.mutation
}

// RemoveFundCoveredPersonIDs removes the Fund_CoveredPerson edge to CoveredPerson by ids.
func (fuo *FundUpdateOne) RemoveFundCoveredPersonIDs(ids ...int) *FundUpdateOne {
	fuo.mutation.RemoveFundCoveredPersonIDs(ids...)
	return fuo
}

// RemoveFundCoveredPerson removes Fund_CoveredPerson edges to CoveredPerson.
func (fuo *FundUpdateOne) RemoveFundCoveredPerson(c ...*CoveredPerson) *FundUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return fuo.RemoveFundCoveredPersonIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (fuo *FundUpdateOne) Save(ctx context.Context) (*Fund, error) {
	if v, ok := fuo.mutation.FundName(); ok {
		if err := fund.FundNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Fund_Name", err: fmt.Errorf("ent: validator failed for field \"Fund_Name\": %w", err)}
		}
	}

	var (
		err  error
		node *Fund
	)
	if len(fuo.hooks) == 0 {
		node, err = fuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FundMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fuo.mutation = mutation
			node, err = fuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fuo.hooks) - 1; i >= 0; i-- {
			mut = fuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FundUpdateOne) SaveX(ctx context.Context) *Fund {
	f, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return f
}

// Exec executes the query on the entity.
func (fuo *FundUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FundUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fuo *FundUpdateOne) sqlSave(ctx context.Context) (f *Fund, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fund.Table,
			Columns: fund.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: fund.FieldID,
			},
		},
	}
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Fund.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := fuo.mutation.FundName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fund.FieldFundName,
		})
	}
	if nodes := fuo.mutation.RemovedFundCoveredPersonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fund.FundCoveredPersonTable,
			Columns: []string{fund.FundCoveredPersonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coveredperson.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.FundCoveredPersonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fund.FundCoveredPersonTable,
			Columns: []string{fund.FundCoveredPersonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coveredperson.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	f = &Fund{config: fuo.config}
	_spec.Assign = f.assignValues
	_spec.ScanValues = f.scanValues()
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fund.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return f, nil
}
