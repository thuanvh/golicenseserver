// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/thuanvh/golicenseserver/ent/customer"
	"github.com/thuanvh/golicenseserver/ent/customerlicense"
	"github.com/thuanvh/golicenseserver/ent/license"
)

// CustomerLicenseCreate is the builder for creating a CustomerLicense entity.
type CustomerLicenseCreate struct {
	config
	mutation *CustomerLicenseMutation
	hooks    []Hook
}

// SetLicenseCode sets the "license_code" field.
func (clc *CustomerLicenseCreate) SetLicenseCode(s string) *CustomerLicenseCreate {
	clc.mutation.SetLicenseCode(s)
	return clc
}

// SetActive sets the "active" field.
func (clc *CustomerLicenseCreate) SetActive(b bool) *CustomerLicenseCreate {
	clc.mutation.SetActive(b)
	return clc
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (clc *CustomerLicenseCreate) SetNillableActive(b *bool) *CustomerLicenseCreate {
	if b != nil {
		clc.SetActive(*b)
	}
	return clc
}

// SetStartDate sets the "start_date" field.
func (clc *CustomerLicenseCreate) SetStartDate(t time.Time) *CustomerLicenseCreate {
	clc.mutation.SetStartDate(t)
	return clc
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (clc *CustomerLicenseCreate) SetNillableStartDate(t *time.Time) *CustomerLicenseCreate {
	if t != nil {
		clc.SetStartDate(*t)
	}
	return clc
}

// SetEndDate sets the "end_date" field.
func (clc *CustomerLicenseCreate) SetEndDate(t time.Time) *CustomerLicenseCreate {
	clc.mutation.SetEndDate(t)
	return clc
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (clc *CustomerLicenseCreate) SetNillableEndDate(t *time.Time) *CustomerLicenseCreate {
	if t != nil {
		clc.SetEndDate(*t)
	}
	return clc
}

// SetCPU sets the "cpu" field.
func (clc *CustomerLicenseCreate) SetCPU(i int) *CustomerLicenseCreate {
	clc.mutation.SetCPU(i)
	return clc
}

// SetNillableCPU sets the "cpu" field if the given value is not nil.
func (clc *CustomerLicenseCreate) SetNillableCPU(i *int) *CustomerLicenseCreate {
	if i != nil {
		clc.SetCPU(*i)
	}
	return clc
}

// SetStorage sets the "storage" field.
func (clc *CustomerLicenseCreate) SetStorage(i int) *CustomerLicenseCreate {
	clc.mutation.SetStorage(i)
	return clc
}

// SetNillableStorage sets the "storage" field if the given value is not nil.
func (clc *CustomerLicenseCreate) SetNillableStorage(i *int) *CustomerLicenseCreate {
	if i != nil {
		clc.SetStorage(*i)
	}
	return clc
}

// SetNumberOfNodes sets the "number_of_nodes" field.
func (clc *CustomerLicenseCreate) SetNumberOfNodes(i int) *CustomerLicenseCreate {
	clc.mutation.SetNumberOfNodes(i)
	return clc
}

// SetNillableNumberOfNodes sets the "number_of_nodes" field if the given value is not nil.
func (clc *CustomerLicenseCreate) SetNillableNumberOfNodes(i *int) *CustomerLicenseCreate {
	if i != nil {
		clc.SetNumberOfNodes(*i)
	}
	return clc
}

// SetCustomerID sets the "customer" edge to the Customer entity by ID.
func (clc *CustomerLicenseCreate) SetCustomerID(id int) *CustomerLicenseCreate {
	clc.mutation.SetCustomerID(id)
	return clc
}

// SetNillableCustomerID sets the "customer" edge to the Customer entity by ID if the given value is not nil.
func (clc *CustomerLicenseCreate) SetNillableCustomerID(id *int) *CustomerLicenseCreate {
	if id != nil {
		clc = clc.SetCustomerID(*id)
	}
	return clc
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (clc *CustomerLicenseCreate) SetCustomer(c *Customer) *CustomerLicenseCreate {
	return clc.SetCustomerID(c.ID)
}

// AddLicenseIDs adds the "license" edge to the License entity by IDs.
func (clc *CustomerLicenseCreate) AddLicenseIDs(ids ...int) *CustomerLicenseCreate {
	clc.mutation.AddLicenseIDs(ids...)
	return clc
}

// AddLicense adds the "license" edges to the License entity.
func (clc *CustomerLicenseCreate) AddLicense(l ...*License) *CustomerLicenseCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return clc.AddLicenseIDs(ids...)
}

// Mutation returns the CustomerLicenseMutation object of the builder.
func (clc *CustomerLicenseCreate) Mutation() *CustomerLicenseMutation {
	return clc.mutation
}

// Save creates the CustomerLicense in the database.
func (clc *CustomerLicenseCreate) Save(ctx context.Context) (*CustomerLicense, error) {
	var (
		err  error
		node *CustomerLicense
	)
	clc.defaults()
	if len(clc.hooks) == 0 {
		if err = clc.check(); err != nil {
			return nil, err
		}
		node, err = clc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomerLicenseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = clc.check(); err != nil {
				return nil, err
			}
			clc.mutation = mutation
			if node, err = clc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(clc.hooks) - 1; i >= 0; i-- {
			if clc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = clc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, clc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (clc *CustomerLicenseCreate) SaveX(ctx context.Context) *CustomerLicense {
	v, err := clc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (clc *CustomerLicenseCreate) Exec(ctx context.Context) error {
	_, err := clc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (clc *CustomerLicenseCreate) ExecX(ctx context.Context) {
	if err := clc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (clc *CustomerLicenseCreate) defaults() {
	if _, ok := clc.mutation.Active(); !ok {
		v := customerlicense.DefaultActive
		clc.mutation.SetActive(v)
	}
	if _, ok := clc.mutation.StartDate(); !ok {
		v := customerlicense.DefaultStartDate()
		clc.mutation.SetStartDate(v)
	}
	if _, ok := clc.mutation.EndDate(); !ok {
		v := customerlicense.DefaultEndDate()
		clc.mutation.SetEndDate(v)
	}
	if _, ok := clc.mutation.CPU(); !ok {
		v := customerlicense.DefaultCPU
		clc.mutation.SetCPU(v)
	}
	if _, ok := clc.mutation.Storage(); !ok {
		v := customerlicense.DefaultStorage
		clc.mutation.SetStorage(v)
	}
	if _, ok := clc.mutation.NumberOfNodes(); !ok {
		v := customerlicense.DefaultNumberOfNodes
		clc.mutation.SetNumberOfNodes(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (clc *CustomerLicenseCreate) check() error {
	if _, ok := clc.mutation.LicenseCode(); !ok {
		return &ValidationError{Name: "license_code", err: errors.New(`ent: missing required field "CustomerLicense.license_code"`)}
	}
	if v, ok := clc.mutation.LicenseCode(); ok {
		if err := customerlicense.LicenseCodeValidator(v); err != nil {
			return &ValidationError{Name: "license_code", err: fmt.Errorf(`ent: validator failed for field "CustomerLicense.license_code": %w`, err)}
		}
	}
	if _, ok := clc.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New(`ent: missing required field "CustomerLicense.active"`)}
	}
	if _, ok := clc.mutation.StartDate(); !ok {
		return &ValidationError{Name: "start_date", err: errors.New(`ent: missing required field "CustomerLicense.start_date"`)}
	}
	if _, ok := clc.mutation.EndDate(); !ok {
		return &ValidationError{Name: "end_date", err: errors.New(`ent: missing required field "CustomerLicense.end_date"`)}
	}
	if _, ok := clc.mutation.CPU(); !ok {
		return &ValidationError{Name: "cpu", err: errors.New(`ent: missing required field "CustomerLicense.cpu"`)}
	}
	if _, ok := clc.mutation.Storage(); !ok {
		return &ValidationError{Name: "storage", err: errors.New(`ent: missing required field "CustomerLicense.storage"`)}
	}
	if _, ok := clc.mutation.NumberOfNodes(); !ok {
		return &ValidationError{Name: "number_of_nodes", err: errors.New(`ent: missing required field "CustomerLicense.number_of_nodes"`)}
	}
	return nil
}

func (clc *CustomerLicenseCreate) sqlSave(ctx context.Context) (*CustomerLicense, error) {
	_node, _spec := clc.createSpec()
	if err := sqlgraph.CreateNode(ctx, clc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (clc *CustomerLicenseCreate) createSpec() (*CustomerLicense, *sqlgraph.CreateSpec) {
	var (
		_node = &CustomerLicense{config: clc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: customerlicense.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: customerlicense.FieldID,
			},
		}
	)
	if value, ok := clc.mutation.LicenseCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customerlicense.FieldLicenseCode,
		})
		_node.LicenseCode = value
	}
	if value, ok := clc.mutation.Active(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: customerlicense.FieldActive,
		})
		_node.Active = value
	}
	if value, ok := clc.mutation.StartDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customerlicense.FieldStartDate,
		})
		_node.StartDate = value
	}
	if value, ok := clc.mutation.EndDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customerlicense.FieldEndDate,
		})
		_node.EndDate = value
	}
	if value, ok := clc.mutation.CPU(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: customerlicense.FieldCPU,
		})
		_node.CPU = value
	}
	if value, ok := clc.mutation.Storage(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: customerlicense.FieldStorage,
		})
		_node.Storage = value
	}
	if value, ok := clc.mutation.NumberOfNodes(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: customerlicense.FieldNumberOfNodes,
		})
		_node.NumberOfNodes = value
	}
	if nodes := clc.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   customerlicense.CustomerTable,
			Columns: []string{customerlicense.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: customer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.customer_license_customer = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := clc.mutation.LicenseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customerlicense.LicenseTable,
			Columns: []string{customerlicense.LicenseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: license.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CustomerLicenseCreateBulk is the builder for creating many CustomerLicense entities in bulk.
type CustomerLicenseCreateBulk struct {
	config
	builders []*CustomerLicenseCreate
}

// Save creates the CustomerLicense entities in the database.
func (clcb *CustomerLicenseCreateBulk) Save(ctx context.Context) ([]*CustomerLicense, error) {
	specs := make([]*sqlgraph.CreateSpec, len(clcb.builders))
	nodes := make([]*CustomerLicense, len(clcb.builders))
	mutators := make([]Mutator, len(clcb.builders))
	for i := range clcb.builders {
		func(i int, root context.Context) {
			builder := clcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CustomerLicenseMutation)
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
					_, err = mutators[i+1].Mutate(root, clcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, clcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, clcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (clcb *CustomerLicenseCreateBulk) SaveX(ctx context.Context) []*CustomerLicense {
	v, err := clcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (clcb *CustomerLicenseCreateBulk) Exec(ctx context.Context) error {
	_, err := clcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (clcb *CustomerLicenseCreateBulk) ExecX(ctx context.Context) {
	if err := clcb.Exec(ctx); err != nil {
		panic(err)
	}
}
