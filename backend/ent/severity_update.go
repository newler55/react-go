// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/newler55/app/ent/disease"
	"github.com/newler55/app/ent/predicate"
	"github.com/newler55/app/ent/severity"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// SeverityUpdate is the builder for updating Severity entities.
type SeverityUpdate struct {
	config
	hooks      []Hook
	mutation   *SeverityMutation
	predicates []predicate.Severity
}

// Where adds a new predicate for the builder.
func (su *SeverityUpdate) Where(ps ...predicate.Severity) *SeverityUpdate {
	su.predicates = append(su.predicates, ps...)
	return su
}

// SetName sets the Name field.
func (su *SeverityUpdate) SetName(s string) *SeverityUpdate {
	su.mutation.SetName(s)
	return su
}

// AddDiseaseIDs adds the disease edge to Disease by ids.
func (su *SeverityUpdate) AddDiseaseIDs(ids ...int) *SeverityUpdate {
	su.mutation.AddDiseaseIDs(ids...)
	return su
}

// AddDisease adds the disease edges to Disease.
func (su *SeverityUpdate) AddDisease(d ...*Disease) *SeverityUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return su.AddDiseaseIDs(ids...)
}

// Mutation returns the SeverityMutation object of the builder.
func (su *SeverityUpdate) Mutation() *SeverityMutation {
	return su.mutation
}

// RemoveDiseaseIDs removes the disease edge to Disease by ids.
func (su *SeverityUpdate) RemoveDiseaseIDs(ids ...int) *SeverityUpdate {
	su.mutation.RemoveDiseaseIDs(ids...)
	return su
}

// RemoveDisease removes disease edges to Disease.
func (su *SeverityUpdate) RemoveDisease(d ...*Disease) *SeverityUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return su.RemoveDiseaseIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (su *SeverityUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SeverityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SeverityUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SeverityUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SeverityUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SeverityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   severity.Table,
			Columns: severity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: severity.FieldID,
			},
		},
	}
	if ps := su.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: severity.FieldName,
		})
	}
	if nodes := su.mutation.RemovedDiseaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   severity.DiseaseTable,
			Columns: []string{severity.DiseaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: disease.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.DiseaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   severity.DiseaseTable,
			Columns: []string{severity.DiseaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: disease.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{severity.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// SeverityUpdateOne is the builder for updating a single Severity entity.
type SeverityUpdateOne struct {
	config
	hooks    []Hook
	mutation *SeverityMutation
}

// SetName sets the Name field.
func (suo *SeverityUpdateOne) SetName(s string) *SeverityUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// AddDiseaseIDs adds the disease edge to Disease by ids.
func (suo *SeverityUpdateOne) AddDiseaseIDs(ids ...int) *SeverityUpdateOne {
	suo.mutation.AddDiseaseIDs(ids...)
	return suo
}

// AddDisease adds the disease edges to Disease.
func (suo *SeverityUpdateOne) AddDisease(d ...*Disease) *SeverityUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return suo.AddDiseaseIDs(ids...)
}

// Mutation returns the SeverityMutation object of the builder.
func (suo *SeverityUpdateOne) Mutation() *SeverityMutation {
	return suo.mutation
}

// RemoveDiseaseIDs removes the disease edge to Disease by ids.
func (suo *SeverityUpdateOne) RemoveDiseaseIDs(ids ...int) *SeverityUpdateOne {
	suo.mutation.RemoveDiseaseIDs(ids...)
	return suo
}

// RemoveDisease removes disease edges to Disease.
func (suo *SeverityUpdateOne) RemoveDisease(d ...*Disease) *SeverityUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return suo.RemoveDiseaseIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (suo *SeverityUpdateOne) Save(ctx context.Context) (*Severity, error) {

	var (
		err  error
		node *Severity
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SeverityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SeverityUpdateOne) SaveX(ctx context.Context) *Severity {
	s, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// Exec executes the query on the entity.
func (suo *SeverityUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SeverityUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SeverityUpdateOne) sqlSave(ctx context.Context) (s *Severity, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   severity.Table,
			Columns: severity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: severity.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Severity.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := suo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: severity.FieldName,
		})
	}
	if nodes := suo.mutation.RemovedDiseaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   severity.DiseaseTable,
			Columns: []string{severity.DiseaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: disease.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.DiseaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   severity.DiseaseTable,
			Columns: []string{severity.DiseaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: disease.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	s = &Severity{config: suo.config}
	_spec.Assign = s.assignValues
	_spec.ScanValues = s.scanValues()
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{severity.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return s, nil
}
