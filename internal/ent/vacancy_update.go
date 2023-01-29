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
	"github.com/logn-soft/logn-back/internal/ent/area"
	"github.com/logn-soft/logn-back/internal/ent/location"
	"github.com/logn-soft/logn-back/internal/ent/predicate"
	"github.com/logn-soft/logn-back/internal/ent/technology"
	"github.com/logn-soft/logn-back/internal/ent/technologylevel"
	"github.com/logn-soft/logn-back/internal/ent/vacancy"
)

// VacancyUpdate is the builder for updating Vacancy entities.
type VacancyUpdate struct {
	config
	hooks    []Hook
	mutation *VacancyMutation
}

// Where appends a list predicates to the VacancyUpdate builder.
func (vu *VacancyUpdate) Where(ps ...predicate.Vacancy) *VacancyUpdate {
	vu.mutation.Where(ps...)
	return vu
}

// SetName sets the "name" field.
func (vu *VacancyUpdate) SetName(s string) *VacancyUpdate {
	vu.mutation.SetName(s)
	return vu
}

// SetDescription sets the "description" field.
func (vu *VacancyUpdate) SetDescription(s string) *VacancyUpdate {
	vu.mutation.SetDescription(s)
	return vu
}

// SetIsNegotiate sets the "is_negotiate" field.
func (vu *VacancyUpdate) SetIsNegotiate(b bool) *VacancyUpdate {
	vu.mutation.SetIsNegotiate(b)
	return vu
}

// SetMinSalary sets the "min_salary" field.
func (vu *VacancyUpdate) SetMinSalary(i int) *VacancyUpdate {
	vu.mutation.ResetMinSalary()
	vu.mutation.SetMinSalary(i)
	return vu
}

// AddMinSalary adds i to the "min_salary" field.
func (vu *VacancyUpdate) AddMinSalary(i int) *VacancyUpdate {
	vu.mutation.AddMinSalary(i)
	return vu
}

// SetMaxSalary sets the "max_salary" field.
func (vu *VacancyUpdate) SetMaxSalary(i int) *VacancyUpdate {
	vu.mutation.ResetMaxSalary()
	vu.mutation.SetMaxSalary(i)
	return vu
}

// AddMaxSalary adds i to the "max_salary" field.
func (vu *VacancyUpdate) AddMaxSalary(i int) *VacancyUpdate {
	vu.mutation.AddMaxSalary(i)
	return vu
}

// SetIsRemote sets the "is_remote" field.
func (vu *VacancyUpdate) SetIsRemote(b bool) *VacancyUpdate {
	vu.mutation.SetIsRemote(b)
	return vu
}

// SetViews sets the "views" field.
func (vu *VacancyUpdate) SetViews(i int) *VacancyUpdate {
	vu.mutation.ResetViews()
	vu.mutation.SetViews(i)
	return vu
}

// AddViews adds i to the "views" field.
func (vu *VacancyUpdate) AddViews(i int) *VacancyUpdate {
	vu.mutation.AddViews(i)
	return vu
}

// SetCreatedAt sets the "created_at" field.
func (vu *VacancyUpdate) SetCreatedAt(t time.Time) *VacancyUpdate {
	vu.mutation.SetCreatedAt(t)
	return vu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vu *VacancyUpdate) SetNillableCreatedAt(t *time.Time) *VacancyUpdate {
	if t != nil {
		vu.SetCreatedAt(*t)
	}
	return vu
}

// AddTechnologyIDs adds the "technologies" edge to the Technology entity by IDs.
func (vu *VacancyUpdate) AddTechnologyIDs(ids ...int) *VacancyUpdate {
	vu.mutation.AddTechnologyIDs(ids...)
	return vu
}

// AddTechnologies adds the "technologies" edges to the Technology entity.
func (vu *VacancyUpdate) AddTechnologies(t ...*Technology) *VacancyUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return vu.AddTechnologyIDs(ids...)
}

// AddLocationIDs adds the "locations" edge to the Location entity by IDs.
func (vu *VacancyUpdate) AddLocationIDs(ids ...int) *VacancyUpdate {
	vu.mutation.AddLocationIDs(ids...)
	return vu
}

// AddLocations adds the "locations" edges to the Location entity.
func (vu *VacancyUpdate) AddLocations(l ...*Location) *VacancyUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return vu.AddLocationIDs(ids...)
}

// AddAreaIDs adds the "areas" edge to the Area entity by IDs.
func (vu *VacancyUpdate) AddAreaIDs(ids ...int) *VacancyUpdate {
	vu.mutation.AddAreaIDs(ids...)
	return vu
}

// AddAreas adds the "areas" edges to the Area entity.
func (vu *VacancyUpdate) AddAreas(a ...*Area) *VacancyUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return vu.AddAreaIDs(ids...)
}

// AddTechnologyLevelIDs adds the "technology_levels" edge to the TechnologyLevel entity by IDs.
func (vu *VacancyUpdate) AddTechnologyLevelIDs(ids ...int) *VacancyUpdate {
	vu.mutation.AddTechnologyLevelIDs(ids...)
	return vu
}

// AddTechnologyLevels adds the "technology_levels" edges to the TechnologyLevel entity.
func (vu *VacancyUpdate) AddTechnologyLevels(t ...*TechnologyLevel) *VacancyUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return vu.AddTechnologyLevelIDs(ids...)
}

// Mutation returns the VacancyMutation object of the builder.
func (vu *VacancyUpdate) Mutation() *VacancyMutation {
	return vu.mutation
}

// ClearTechnologies clears all "technologies" edges to the Technology entity.
func (vu *VacancyUpdate) ClearTechnologies() *VacancyUpdate {
	vu.mutation.ClearTechnologies()
	return vu
}

// RemoveTechnologyIDs removes the "technologies" edge to Technology entities by IDs.
func (vu *VacancyUpdate) RemoveTechnologyIDs(ids ...int) *VacancyUpdate {
	vu.mutation.RemoveTechnologyIDs(ids...)
	return vu
}

// RemoveTechnologies removes "technologies" edges to Technology entities.
func (vu *VacancyUpdate) RemoveTechnologies(t ...*Technology) *VacancyUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return vu.RemoveTechnologyIDs(ids...)
}

// ClearLocations clears all "locations" edges to the Location entity.
func (vu *VacancyUpdate) ClearLocations() *VacancyUpdate {
	vu.mutation.ClearLocations()
	return vu
}

// RemoveLocationIDs removes the "locations" edge to Location entities by IDs.
func (vu *VacancyUpdate) RemoveLocationIDs(ids ...int) *VacancyUpdate {
	vu.mutation.RemoveLocationIDs(ids...)
	return vu
}

// RemoveLocations removes "locations" edges to Location entities.
func (vu *VacancyUpdate) RemoveLocations(l ...*Location) *VacancyUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return vu.RemoveLocationIDs(ids...)
}

// ClearAreas clears all "areas" edges to the Area entity.
func (vu *VacancyUpdate) ClearAreas() *VacancyUpdate {
	vu.mutation.ClearAreas()
	return vu
}

// RemoveAreaIDs removes the "areas" edge to Area entities by IDs.
func (vu *VacancyUpdate) RemoveAreaIDs(ids ...int) *VacancyUpdate {
	vu.mutation.RemoveAreaIDs(ids...)
	return vu
}

// RemoveAreas removes "areas" edges to Area entities.
func (vu *VacancyUpdate) RemoveAreas(a ...*Area) *VacancyUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return vu.RemoveAreaIDs(ids...)
}

// ClearTechnologyLevels clears all "technology_levels" edges to the TechnologyLevel entity.
func (vu *VacancyUpdate) ClearTechnologyLevels() *VacancyUpdate {
	vu.mutation.ClearTechnologyLevels()
	return vu
}

// RemoveTechnologyLevelIDs removes the "technology_levels" edge to TechnologyLevel entities by IDs.
func (vu *VacancyUpdate) RemoveTechnologyLevelIDs(ids ...int) *VacancyUpdate {
	vu.mutation.RemoveTechnologyLevelIDs(ids...)
	return vu
}

// RemoveTechnologyLevels removes "technology_levels" edges to TechnologyLevel entities.
func (vu *VacancyUpdate) RemoveTechnologyLevels(t ...*TechnologyLevel) *VacancyUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return vu.RemoveTechnologyLevelIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vu *VacancyUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, VacancyMutation](ctx, vu.sqlSave, vu.mutation, vu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vu *VacancyUpdate) SaveX(ctx context.Context) int {
	affected, err := vu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vu *VacancyUpdate) Exec(ctx context.Context) error {
	_, err := vu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vu *VacancyUpdate) ExecX(ctx context.Context) {
	if err := vu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vu *VacancyUpdate) check() error {
	if v, ok := vu.mutation.Name(); ok {
		if err := vacancy.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Vacancy.name": %w`, err)}
		}
	}
	if v, ok := vu.mutation.Description(); ok {
		if err := vacancy.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Vacancy.description": %w`, err)}
		}
	}
	if v, ok := vu.mutation.MinSalary(); ok {
		if err := vacancy.MinSalaryValidator(v); err != nil {
			return &ValidationError{Name: "min_salary", err: fmt.Errorf(`ent: validator failed for field "Vacancy.min_salary": %w`, err)}
		}
	}
	return nil
}

func (vu *VacancyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := vu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   vacancy.Table,
			Columns: vacancy.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: vacancy.FieldID,
			},
		},
	}
	if ps := vu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vu.mutation.Name(); ok {
		_spec.SetField(vacancy.FieldName, field.TypeString, value)
	}
	if value, ok := vu.mutation.Description(); ok {
		_spec.SetField(vacancy.FieldDescription, field.TypeString, value)
	}
	if value, ok := vu.mutation.IsNegotiate(); ok {
		_spec.SetField(vacancy.FieldIsNegotiate, field.TypeBool, value)
	}
	if value, ok := vu.mutation.MinSalary(); ok {
		_spec.SetField(vacancy.FieldMinSalary, field.TypeInt, value)
	}
	if value, ok := vu.mutation.AddedMinSalary(); ok {
		_spec.AddField(vacancy.FieldMinSalary, field.TypeInt, value)
	}
	if value, ok := vu.mutation.MaxSalary(); ok {
		_spec.SetField(vacancy.FieldMaxSalary, field.TypeInt, value)
	}
	if value, ok := vu.mutation.AddedMaxSalary(); ok {
		_spec.AddField(vacancy.FieldMaxSalary, field.TypeInt, value)
	}
	if value, ok := vu.mutation.IsRemote(); ok {
		_spec.SetField(vacancy.FieldIsRemote, field.TypeBool, value)
	}
	if value, ok := vu.mutation.Views(); ok {
		_spec.SetField(vacancy.FieldViews, field.TypeInt, value)
	}
	if value, ok := vu.mutation.AddedViews(); ok {
		_spec.AddField(vacancy.FieldViews, field.TypeInt, value)
	}
	if value, ok := vu.mutation.CreatedAt(); ok {
		_spec.SetField(vacancy.FieldCreatedAt, field.TypeTime, value)
	}
	if vu.mutation.TechnologiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vacancy.TechnologiesTable,
			Columns: vacancy.TechnologiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: technology.FieldID,
				},
			},
		}
		createE := &TechnologyLevelCreate{config: vu.config, mutation: newTechnologyLevelMutation(vu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.RemovedTechnologiesIDs(); len(nodes) > 0 && !vu.mutation.TechnologiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vacancy.TechnologiesTable,
			Columns: vacancy.TechnologiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: technology.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TechnologyLevelCreate{config: vu.config, mutation: newTechnologyLevelMutation(vu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.TechnologiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vacancy.TechnologiesTable,
			Columns: vacancy.TechnologiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: technology.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TechnologyLevelCreate{config: vu.config, mutation: newTechnologyLevelMutation(vu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vu.mutation.LocationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vacancy.LocationsTable,
			Columns: vacancy.LocationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: location.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.RemovedLocationsIDs(); len(nodes) > 0 && !vu.mutation.LocationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vacancy.LocationsTable,
			Columns: vacancy.LocationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: location.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.LocationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vacancy.LocationsTable,
			Columns: vacancy.LocationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: location.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vu.mutation.AreasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vacancy.AreasTable,
			Columns: vacancy.AreasPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: area.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.RemovedAreasIDs(); len(nodes) > 0 && !vu.mutation.AreasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vacancy.AreasTable,
			Columns: vacancy.AreasPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: area.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.AreasIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vacancy.AreasTable,
			Columns: vacancy.AreasPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: area.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vu.mutation.TechnologyLevelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   vacancy.TechnologyLevelsTable,
			Columns: []string{vacancy.TechnologyLevelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: technologylevel.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.RemovedTechnologyLevelsIDs(); len(nodes) > 0 && !vu.mutation.TechnologyLevelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   vacancy.TechnologyLevelsTable,
			Columns: []string{vacancy.TechnologyLevelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: technologylevel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.TechnologyLevelsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   vacancy.TechnologyLevelsTable,
			Columns: []string{vacancy.TechnologyLevelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: technologylevel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vacancy.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	vu.mutation.done = true
	return n, nil
}

// VacancyUpdateOne is the builder for updating a single Vacancy entity.
type VacancyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VacancyMutation
}

// SetName sets the "name" field.
func (vuo *VacancyUpdateOne) SetName(s string) *VacancyUpdateOne {
	vuo.mutation.SetName(s)
	return vuo
}

// SetDescription sets the "description" field.
func (vuo *VacancyUpdateOne) SetDescription(s string) *VacancyUpdateOne {
	vuo.mutation.SetDescription(s)
	return vuo
}

// SetIsNegotiate sets the "is_negotiate" field.
func (vuo *VacancyUpdateOne) SetIsNegotiate(b bool) *VacancyUpdateOne {
	vuo.mutation.SetIsNegotiate(b)
	return vuo
}

// SetMinSalary sets the "min_salary" field.
func (vuo *VacancyUpdateOne) SetMinSalary(i int) *VacancyUpdateOne {
	vuo.mutation.ResetMinSalary()
	vuo.mutation.SetMinSalary(i)
	return vuo
}

// AddMinSalary adds i to the "min_salary" field.
func (vuo *VacancyUpdateOne) AddMinSalary(i int) *VacancyUpdateOne {
	vuo.mutation.AddMinSalary(i)
	return vuo
}

// SetMaxSalary sets the "max_salary" field.
func (vuo *VacancyUpdateOne) SetMaxSalary(i int) *VacancyUpdateOne {
	vuo.mutation.ResetMaxSalary()
	vuo.mutation.SetMaxSalary(i)
	return vuo
}

// AddMaxSalary adds i to the "max_salary" field.
func (vuo *VacancyUpdateOne) AddMaxSalary(i int) *VacancyUpdateOne {
	vuo.mutation.AddMaxSalary(i)
	return vuo
}

// SetIsRemote sets the "is_remote" field.
func (vuo *VacancyUpdateOne) SetIsRemote(b bool) *VacancyUpdateOne {
	vuo.mutation.SetIsRemote(b)
	return vuo
}

// SetViews sets the "views" field.
func (vuo *VacancyUpdateOne) SetViews(i int) *VacancyUpdateOne {
	vuo.mutation.ResetViews()
	vuo.mutation.SetViews(i)
	return vuo
}

// AddViews adds i to the "views" field.
func (vuo *VacancyUpdateOne) AddViews(i int) *VacancyUpdateOne {
	vuo.mutation.AddViews(i)
	return vuo
}

// SetCreatedAt sets the "created_at" field.
func (vuo *VacancyUpdateOne) SetCreatedAt(t time.Time) *VacancyUpdateOne {
	vuo.mutation.SetCreatedAt(t)
	return vuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vuo *VacancyUpdateOne) SetNillableCreatedAt(t *time.Time) *VacancyUpdateOne {
	if t != nil {
		vuo.SetCreatedAt(*t)
	}
	return vuo
}

// AddTechnologyIDs adds the "technologies" edge to the Technology entity by IDs.
func (vuo *VacancyUpdateOne) AddTechnologyIDs(ids ...int) *VacancyUpdateOne {
	vuo.mutation.AddTechnologyIDs(ids...)
	return vuo
}

// AddTechnologies adds the "technologies" edges to the Technology entity.
func (vuo *VacancyUpdateOne) AddTechnologies(t ...*Technology) *VacancyUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return vuo.AddTechnologyIDs(ids...)
}

// AddLocationIDs adds the "locations" edge to the Location entity by IDs.
func (vuo *VacancyUpdateOne) AddLocationIDs(ids ...int) *VacancyUpdateOne {
	vuo.mutation.AddLocationIDs(ids...)
	return vuo
}

// AddLocations adds the "locations" edges to the Location entity.
func (vuo *VacancyUpdateOne) AddLocations(l ...*Location) *VacancyUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return vuo.AddLocationIDs(ids...)
}

// AddAreaIDs adds the "areas" edge to the Area entity by IDs.
func (vuo *VacancyUpdateOne) AddAreaIDs(ids ...int) *VacancyUpdateOne {
	vuo.mutation.AddAreaIDs(ids...)
	return vuo
}

// AddAreas adds the "areas" edges to the Area entity.
func (vuo *VacancyUpdateOne) AddAreas(a ...*Area) *VacancyUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return vuo.AddAreaIDs(ids...)
}

// AddTechnologyLevelIDs adds the "technology_levels" edge to the TechnologyLevel entity by IDs.
func (vuo *VacancyUpdateOne) AddTechnologyLevelIDs(ids ...int) *VacancyUpdateOne {
	vuo.mutation.AddTechnologyLevelIDs(ids...)
	return vuo
}

// AddTechnologyLevels adds the "technology_levels" edges to the TechnologyLevel entity.
func (vuo *VacancyUpdateOne) AddTechnologyLevels(t ...*TechnologyLevel) *VacancyUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return vuo.AddTechnologyLevelIDs(ids...)
}

// Mutation returns the VacancyMutation object of the builder.
func (vuo *VacancyUpdateOne) Mutation() *VacancyMutation {
	return vuo.mutation
}

// ClearTechnologies clears all "technologies" edges to the Technology entity.
func (vuo *VacancyUpdateOne) ClearTechnologies() *VacancyUpdateOne {
	vuo.mutation.ClearTechnologies()
	return vuo
}

// RemoveTechnologyIDs removes the "technologies" edge to Technology entities by IDs.
func (vuo *VacancyUpdateOne) RemoveTechnologyIDs(ids ...int) *VacancyUpdateOne {
	vuo.mutation.RemoveTechnologyIDs(ids...)
	return vuo
}

// RemoveTechnologies removes "technologies" edges to Technology entities.
func (vuo *VacancyUpdateOne) RemoveTechnologies(t ...*Technology) *VacancyUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return vuo.RemoveTechnologyIDs(ids...)
}

// ClearLocations clears all "locations" edges to the Location entity.
func (vuo *VacancyUpdateOne) ClearLocations() *VacancyUpdateOne {
	vuo.mutation.ClearLocations()
	return vuo
}

// RemoveLocationIDs removes the "locations" edge to Location entities by IDs.
func (vuo *VacancyUpdateOne) RemoveLocationIDs(ids ...int) *VacancyUpdateOne {
	vuo.mutation.RemoveLocationIDs(ids...)
	return vuo
}

// RemoveLocations removes "locations" edges to Location entities.
func (vuo *VacancyUpdateOne) RemoveLocations(l ...*Location) *VacancyUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return vuo.RemoveLocationIDs(ids...)
}

// ClearAreas clears all "areas" edges to the Area entity.
func (vuo *VacancyUpdateOne) ClearAreas() *VacancyUpdateOne {
	vuo.mutation.ClearAreas()
	return vuo
}

// RemoveAreaIDs removes the "areas" edge to Area entities by IDs.
func (vuo *VacancyUpdateOne) RemoveAreaIDs(ids ...int) *VacancyUpdateOne {
	vuo.mutation.RemoveAreaIDs(ids...)
	return vuo
}

// RemoveAreas removes "areas" edges to Area entities.
func (vuo *VacancyUpdateOne) RemoveAreas(a ...*Area) *VacancyUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return vuo.RemoveAreaIDs(ids...)
}

// ClearTechnologyLevels clears all "technology_levels" edges to the TechnologyLevel entity.
func (vuo *VacancyUpdateOne) ClearTechnologyLevels() *VacancyUpdateOne {
	vuo.mutation.ClearTechnologyLevels()
	return vuo
}

// RemoveTechnologyLevelIDs removes the "technology_levels" edge to TechnologyLevel entities by IDs.
func (vuo *VacancyUpdateOne) RemoveTechnologyLevelIDs(ids ...int) *VacancyUpdateOne {
	vuo.mutation.RemoveTechnologyLevelIDs(ids...)
	return vuo
}

// RemoveTechnologyLevels removes "technology_levels" edges to TechnologyLevel entities.
func (vuo *VacancyUpdateOne) RemoveTechnologyLevels(t ...*TechnologyLevel) *VacancyUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return vuo.RemoveTechnologyLevelIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vuo *VacancyUpdateOne) Select(field string, fields ...string) *VacancyUpdateOne {
	vuo.fields = append([]string{field}, fields...)
	return vuo
}

// Save executes the query and returns the updated Vacancy entity.
func (vuo *VacancyUpdateOne) Save(ctx context.Context) (*Vacancy, error) {
	return withHooks[*Vacancy, VacancyMutation](ctx, vuo.sqlSave, vuo.mutation, vuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vuo *VacancyUpdateOne) SaveX(ctx context.Context) *Vacancy {
	node, err := vuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vuo *VacancyUpdateOne) Exec(ctx context.Context) error {
	_, err := vuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuo *VacancyUpdateOne) ExecX(ctx context.Context) {
	if err := vuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vuo *VacancyUpdateOne) check() error {
	if v, ok := vuo.mutation.Name(); ok {
		if err := vacancy.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Vacancy.name": %w`, err)}
		}
	}
	if v, ok := vuo.mutation.Description(); ok {
		if err := vacancy.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Vacancy.description": %w`, err)}
		}
	}
	if v, ok := vuo.mutation.MinSalary(); ok {
		if err := vacancy.MinSalaryValidator(v); err != nil {
			return &ValidationError{Name: "min_salary", err: fmt.Errorf(`ent: validator failed for field "Vacancy.min_salary": %w`, err)}
		}
	}
	return nil
}

func (vuo *VacancyUpdateOne) sqlSave(ctx context.Context) (_node *Vacancy, err error) {
	if err := vuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   vacancy.Table,
			Columns: vacancy.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: vacancy.FieldID,
			},
		},
	}
	id, ok := vuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Vacancy.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vacancy.FieldID)
		for _, f := range fields {
			if !vacancy.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != vacancy.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vuo.mutation.Name(); ok {
		_spec.SetField(vacancy.FieldName, field.TypeString, value)
	}
	if value, ok := vuo.mutation.Description(); ok {
		_spec.SetField(vacancy.FieldDescription, field.TypeString, value)
	}
	if value, ok := vuo.mutation.IsNegotiate(); ok {
		_spec.SetField(vacancy.FieldIsNegotiate, field.TypeBool, value)
	}
	if value, ok := vuo.mutation.MinSalary(); ok {
		_spec.SetField(vacancy.FieldMinSalary, field.TypeInt, value)
	}
	if value, ok := vuo.mutation.AddedMinSalary(); ok {
		_spec.AddField(vacancy.FieldMinSalary, field.TypeInt, value)
	}
	if value, ok := vuo.mutation.MaxSalary(); ok {
		_spec.SetField(vacancy.FieldMaxSalary, field.TypeInt, value)
	}
	if value, ok := vuo.mutation.AddedMaxSalary(); ok {
		_spec.AddField(vacancy.FieldMaxSalary, field.TypeInt, value)
	}
	if value, ok := vuo.mutation.IsRemote(); ok {
		_spec.SetField(vacancy.FieldIsRemote, field.TypeBool, value)
	}
	if value, ok := vuo.mutation.Views(); ok {
		_spec.SetField(vacancy.FieldViews, field.TypeInt, value)
	}
	if value, ok := vuo.mutation.AddedViews(); ok {
		_spec.AddField(vacancy.FieldViews, field.TypeInt, value)
	}
	if value, ok := vuo.mutation.CreatedAt(); ok {
		_spec.SetField(vacancy.FieldCreatedAt, field.TypeTime, value)
	}
	if vuo.mutation.TechnologiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vacancy.TechnologiesTable,
			Columns: vacancy.TechnologiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: technology.FieldID,
				},
			},
		}
		createE := &TechnologyLevelCreate{config: vuo.config, mutation: newTechnologyLevelMutation(vuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.RemovedTechnologiesIDs(); len(nodes) > 0 && !vuo.mutation.TechnologiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vacancy.TechnologiesTable,
			Columns: vacancy.TechnologiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: technology.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TechnologyLevelCreate{config: vuo.config, mutation: newTechnologyLevelMutation(vuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.TechnologiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vacancy.TechnologiesTable,
			Columns: vacancy.TechnologiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: technology.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TechnologyLevelCreate{config: vuo.config, mutation: newTechnologyLevelMutation(vuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vuo.mutation.LocationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vacancy.LocationsTable,
			Columns: vacancy.LocationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: location.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.RemovedLocationsIDs(); len(nodes) > 0 && !vuo.mutation.LocationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vacancy.LocationsTable,
			Columns: vacancy.LocationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: location.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.LocationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vacancy.LocationsTable,
			Columns: vacancy.LocationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: location.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vuo.mutation.AreasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vacancy.AreasTable,
			Columns: vacancy.AreasPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: area.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.RemovedAreasIDs(); len(nodes) > 0 && !vuo.mutation.AreasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vacancy.AreasTable,
			Columns: vacancy.AreasPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: area.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.AreasIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vacancy.AreasTable,
			Columns: vacancy.AreasPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: area.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vuo.mutation.TechnologyLevelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   vacancy.TechnologyLevelsTable,
			Columns: []string{vacancy.TechnologyLevelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: technologylevel.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.RemovedTechnologyLevelsIDs(); len(nodes) > 0 && !vuo.mutation.TechnologyLevelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   vacancy.TechnologyLevelsTable,
			Columns: []string{vacancy.TechnologyLevelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: technologylevel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.TechnologyLevelsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   vacancy.TechnologyLevelsTable,
			Columns: []string{vacancy.TechnologyLevelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: technologylevel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Vacancy{config: vuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vacancy.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	vuo.mutation.done = true
	return _node, nil
}