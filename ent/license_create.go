// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/thuanvh/golicenseserver/ent/license"
)

// LicenseCreate is the builder for creating a License entity.
type LicenseCreate struct {
	config
	mutation *LicenseMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (lc *LicenseCreate) SetName(s string) *LicenseCreate {
	lc.mutation.SetName(s)
	return lc
}

// SetDuration sets the "duration" field.
func (lc *LicenseCreate) SetDuration(i int) *LicenseCreate {
	lc.mutation.SetDuration(i)
	return lc
}

// SetCPU sets the "cpu" field.
func (lc *LicenseCreate) SetCPU(i int) *LicenseCreate {
	lc.mutation.SetCPU(i)
	return lc
}

// SetNillableCPU sets the "cpu" field if the given value is not nil.
func (lc *LicenseCreate) SetNillableCPU(i *int) *LicenseCreate {
	if i != nil {
		lc.SetCPU(*i)
	}
	return lc
}

// SetStorage sets the "storage" field.
func (lc *LicenseCreate) SetStorage(i int) *LicenseCreate {
	lc.mutation.SetStorage(i)
	return lc
}

// SetNillableStorage sets the "storage" field if the given value is not nil.
func (lc *LicenseCreate) SetNillableStorage(i *int) *LicenseCreate {
	if i != nil {
		lc.SetStorage(*i)
	}
	return lc
}

// SetNumberOfNodes sets the "number_of_nodes" field.
func (lc *LicenseCreate) SetNumberOfNodes(i int) *LicenseCreate {
	lc.mutation.SetNumberOfNodes(i)
	return lc
}

// SetNillableNumberOfNodes sets the "number_of_nodes" field if the given value is not nil.
func (lc *LicenseCreate) SetNillableNumberOfNodes(i *int) *LicenseCreate {
	if i != nil {
		lc.SetNumberOfNodes(*i)
	}
	return lc
}

// Mutation returns the LicenseMutation object of the builder.
func (lc *LicenseCreate) Mutation() *LicenseMutation {
	return lc.mutation
}

// Save creates the License in the database.
func (lc *LicenseCreate) Save(ctx context.Context) (*License, error) {
	var (
		err  error
		node *License
	)
	lc.defaults()
	if len(lc.hooks) == 0 {
		if err = lc.check(); err != nil {
			return nil, err
		}
		node, err = lc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LicenseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lc.check(); err != nil {
				return nil, err
			}
			lc.mutation = mutation
			if node, err = lc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(lc.hooks) - 1; i >= 0; i-- {
			if lc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LicenseCreate) SaveX(ctx context.Context) *License {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LicenseCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LicenseCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LicenseCreate) defaults() {
	if _, ok := lc.mutation.CPU(); !ok {
		v := license.DefaultCPU
		lc.mutation.SetCPU(v)
	}
	if _, ok := lc.mutation.Storage(); !ok {
		v := license.DefaultStorage
		lc.mutation.SetStorage(v)
	}
	if _, ok := lc.mutation.NumberOfNodes(); !ok {
		v := license.DefaultNumberOfNodes
		lc.mutation.SetNumberOfNodes(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LicenseCreate) check() error {
	if _, ok := lc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "License.name"`)}
	}
	if v, ok := lc.mutation.Name(); ok {
		if err := license.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "License.name": %w`, err)}
		}
	}
	if _, ok := lc.mutation.Duration(); !ok {
		return &ValidationError{Name: "duration", err: errors.New(`ent: missing required field "License.duration"`)}
	}
	if v, ok := lc.mutation.Duration(); ok {
		if err := license.DurationValidator(v); err != nil {
			return &ValidationError{Name: "duration", err: fmt.Errorf(`ent: validator failed for field "License.duration": %w`, err)}
		}
	}
	if _, ok := lc.mutation.CPU(); !ok {
		return &ValidationError{Name: "cpu", err: errors.New(`ent: missing required field "License.cpu"`)}
	}
	if _, ok := lc.mutation.Storage(); !ok {
		return &ValidationError{Name: "storage", err: errors.New(`ent: missing required field "License.storage"`)}
	}
	if _, ok := lc.mutation.NumberOfNodes(); !ok {
		return &ValidationError{Name: "number_of_nodes", err: errors.New(`ent: missing required field "License.number_of_nodes"`)}
	}
	return nil
}

func (lc *LicenseCreate) sqlSave(ctx context.Context) (*License, error) {
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (lc *LicenseCreate) createSpec() (*License, *sqlgraph.CreateSpec) {
	var (
		_node = &License{config: lc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: license.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: license.FieldID,
			},
		}
	)
	if value, ok := lc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: license.FieldName,
		})
		_node.Name = value
	}
	if value, ok := lc.mutation.Duration(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: license.FieldDuration,
		})
		_node.Duration = value
	}
	if value, ok := lc.mutation.CPU(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: license.FieldCPU,
		})
		_node.CPU = value
	}
	if value, ok := lc.mutation.Storage(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: license.FieldStorage,
		})
		_node.Storage = value
	}
	if value, ok := lc.mutation.NumberOfNodes(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: license.FieldNumberOfNodes,
		})
		_node.NumberOfNodes = value
	}
	return _node, _spec
}

// LicenseCreateBulk is the builder for creating many License entities in bulk.
type LicenseCreateBulk struct {
	config
	builders []*LicenseCreate
}

// Save creates the License entities in the database.
func (lcb *LicenseCreateBulk) Save(ctx context.Context) ([]*License, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*License, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LicenseMutation)
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
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LicenseCreateBulk) SaveX(ctx context.Context) []*License {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LicenseCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LicenseCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}
