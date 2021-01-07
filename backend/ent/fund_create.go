// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/mmildd_s/app/ent/coveredperson"
	"github.com/mmildd_s/app/ent/fund"
)

// FundCreate is the builder for creating a Fund entity.
type FundCreate struct {
	config
	mutation *FundMutation
	hooks    []Hook
}

// SetFundName sets the Fund_Name field.
func (fc *FundCreate) SetFundName(s string) *FundCreate {
	fc.mutation.SetFundName(s)
	return fc
}

// AddFundCoveredPersonIDs adds the Fund_CoveredPerson edge to CoveredPerson by ids.
func (fc *FundCreate) AddFundCoveredPersonIDs(ids ...int) *FundCreate {
	fc.mutation.AddFundCoveredPersonIDs(ids...)
	return fc
}

// AddFundCoveredPerson adds the Fund_CoveredPerson edges to CoveredPerson.
func (fc *FundCreate) AddFundCoveredPerson(c ...*CoveredPerson) *FundCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return fc.AddFundCoveredPersonIDs(ids...)
}

// Mutation returns the FundMutation object of the builder.
func (fc *FundCreate) Mutation() *FundMutation {
	return fc.mutation
}

// Save creates the Fund in the database.
func (fc *FundCreate) Save(ctx context.Context) (*Fund, error) {
	if _, ok := fc.mutation.FundName(); !ok {
		return nil, &ValidationError{Name: "Fund_Name", err: errors.New("ent: missing required field \"Fund_Name\"")}
	}
	if v, ok := fc.mutation.FundName(); ok {
		if err := fund.FundNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Fund_Name", err: fmt.Errorf("ent: validator failed for field \"Fund_Name\": %w", err)}
		}
	}
	var (
		err  error
		node *Fund
	)
	if len(fc.hooks) == 0 {
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FundMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fc.mutation = mutation
			node, err = fc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fc.hooks) - 1; i >= 0; i-- {
			mut = fc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FundCreate) SaveX(ctx context.Context) *Fund {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fc *FundCreate) sqlSave(ctx context.Context) (*Fund, error) {
	f, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	f.ID = int(id)
	return f, nil
}

func (fc *FundCreate) createSpec() (*Fund, *sqlgraph.CreateSpec) {
	var (
		f     = &Fund{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: fund.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: fund.FieldID,
			},
		}
	)
	if value, ok := fc.mutation.FundName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fund.FieldFundName,
		})
		f.FundName = value
	}
	if nodes := fc.mutation.FundCoveredPersonIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return f, _spec
}