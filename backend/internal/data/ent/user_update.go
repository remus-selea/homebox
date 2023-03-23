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
	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/ent/authtokens"
	"github.com/hay-kot/homebox/backend/internal/data/ent/group"
	"github.com/hay-kot/homebox/backend/internal/data/ent/notifier"
	"github.com/hay-kot/homebox/backend/internal/data/ent/predicate"
	"github.com/hay-kot/homebox/backend/internal/data/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetIsSuperuser sets the "is_superuser" field.
func (uu *UserUpdate) SetIsSuperuser(b bool) *UserUpdate {
	uu.mutation.SetIsSuperuser(b)
	return uu
}

// SetNillableIsSuperuser sets the "is_superuser" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIsSuperuser(b *bool) *UserUpdate {
	if b != nil {
		uu.SetIsSuperuser(*b)
	}
	return uu
}

// SetSuperuser sets the "superuser" field.
func (uu *UserUpdate) SetSuperuser(b bool) *UserUpdate {
	uu.mutation.SetSuperuser(b)
	return uu
}

// SetNillableSuperuser sets the "superuser" field if the given value is not nil.
func (uu *UserUpdate) SetNillableSuperuser(b *bool) *UserUpdate {
	if b != nil {
		uu.SetSuperuser(*b)
	}
	return uu
}

// SetRole sets the "role" field.
func (uu *UserUpdate) SetRole(u user.Role) *UserUpdate {
	uu.mutation.SetRole(u)
	return uu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRole(u *user.Role) *UserUpdate {
	if u != nil {
		uu.SetRole(*u)
	}
	return uu
}

// SetActivatedOn sets the "activated_on" field.
func (uu *UserUpdate) SetActivatedOn(t time.Time) *UserUpdate {
	uu.mutation.SetActivatedOn(t)
	return uu
}

// SetNillableActivatedOn sets the "activated_on" field if the given value is not nil.
func (uu *UserUpdate) SetNillableActivatedOn(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetActivatedOn(*t)
	}
	return uu
}

// ClearActivatedOn clears the value of the "activated_on" field.
func (uu *UserUpdate) ClearActivatedOn() *UserUpdate {
	uu.mutation.ClearActivatedOn()
	return uu
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (uu *UserUpdate) SetGroupID(id uuid.UUID) *UserUpdate {
	uu.mutation.SetGroupID(id)
	return uu
}

// SetGroup sets the "group" edge to the Group entity.
func (uu *UserUpdate) SetGroup(g *Group) *UserUpdate {
	return uu.SetGroupID(g.ID)
}

// AddAuthTokenIDs adds the "auth_tokens" edge to the AuthTokens entity by IDs.
func (uu *UserUpdate) AddAuthTokenIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddAuthTokenIDs(ids...)
	return uu
}

// AddAuthTokens adds the "auth_tokens" edges to the AuthTokens entity.
func (uu *UserUpdate) AddAuthTokens(a ...*AuthTokens) *UserUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.AddAuthTokenIDs(ids...)
}

// AddNotifierIDs adds the "notifiers" edge to the Notifier entity by IDs.
func (uu *UserUpdate) AddNotifierIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddNotifierIDs(ids...)
	return uu
}

// AddNotifiers adds the "notifiers" edges to the Notifier entity.
func (uu *UserUpdate) AddNotifiers(n ...*Notifier) *UserUpdate {
	ids := make([]uuid.UUID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return uu.AddNotifierIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearGroup clears the "group" edge to the Group entity.
func (uu *UserUpdate) ClearGroup() *UserUpdate {
	uu.mutation.ClearGroup()
	return uu
}

// ClearAuthTokens clears all "auth_tokens" edges to the AuthTokens entity.
func (uu *UserUpdate) ClearAuthTokens() *UserUpdate {
	uu.mutation.ClearAuthTokens()
	return uu
}

// RemoveAuthTokenIDs removes the "auth_tokens" edge to AuthTokens entities by IDs.
func (uu *UserUpdate) RemoveAuthTokenIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveAuthTokenIDs(ids...)
	return uu
}

// RemoveAuthTokens removes "auth_tokens" edges to AuthTokens entities.
func (uu *UserUpdate) RemoveAuthTokens(a ...*AuthTokens) *UserUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.RemoveAuthTokenIDs(ids...)
}

// ClearNotifiers clears all "notifiers" edges to the Notifier entity.
func (uu *UserUpdate) ClearNotifiers() *UserUpdate {
	uu.mutation.ClearNotifiers()
	return uu
}

// RemoveNotifierIDs removes the "notifiers" edge to Notifier entities by IDs.
func (uu *UserUpdate) RemoveNotifierIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveNotifierIDs(ids...)
	return uu
}

// RemoveNotifiers removes "notifiers" edges to Notifier entities.
func (uu *UserUpdate) RemoveNotifiers(n ...*Notifier) *UserUpdate {
	ids := make([]uuid.UUID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return uu.RemoveNotifierIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks[int, UserMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Role(); ok {
		if err := user.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "User.role": %w`, err)}
		}
	}
	if _, ok := uu.mutation.GroupID(); uu.mutation.GroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "User.group"`)
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.IsSuperuser(); ok {
		_spec.SetField(user.FieldIsSuperuser, field.TypeBool, value)
	}
	if value, ok := uu.mutation.Superuser(); ok {
		_spec.SetField(user.FieldSuperuser, field.TypeBool, value)
	}
	if value, ok := uu.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeEnum, value)
	}
	if value, ok := uu.mutation.ActivatedOn(); ok {
		_spec.SetField(user.FieldActivatedOn, field.TypeTime, value)
	}
	if uu.mutation.ActivatedOnCleared() {
		_spec.ClearField(user.FieldActivatedOn, field.TypeTime)
	}
	if uu.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.GroupTable,
			Columns: []string{user.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.GroupTable,
			Columns: []string{user.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.AuthTokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AuthTokensTable,
			Columns: []string{user.AuthTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authtokens.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedAuthTokensIDs(); len(nodes) > 0 && !uu.mutation.AuthTokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AuthTokensTable,
			Columns: []string{user.AuthTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authtokens.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.AuthTokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AuthTokensTable,
			Columns: []string{user.AuthTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authtokens.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.NotifiersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.NotifiersTable,
			Columns: []string{user.NotifiersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifier.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedNotifiersIDs(); len(nodes) > 0 && !uu.mutation.NotifiersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.NotifiersTable,
			Columns: []string{user.NotifiersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifier.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.NotifiersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.NotifiersTable,
			Columns: []string{user.NotifiersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifier.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetIsSuperuser sets the "is_superuser" field.
func (uuo *UserUpdateOne) SetIsSuperuser(b bool) *UserUpdateOne {
	uuo.mutation.SetIsSuperuser(b)
	return uuo
}

// SetNillableIsSuperuser sets the "is_superuser" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIsSuperuser(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetIsSuperuser(*b)
	}
	return uuo
}

// SetSuperuser sets the "superuser" field.
func (uuo *UserUpdateOne) SetSuperuser(b bool) *UserUpdateOne {
	uuo.mutation.SetSuperuser(b)
	return uuo
}

// SetNillableSuperuser sets the "superuser" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSuperuser(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetSuperuser(*b)
	}
	return uuo
}

// SetRole sets the "role" field.
func (uuo *UserUpdateOne) SetRole(u user.Role) *UserUpdateOne {
	uuo.mutation.SetRole(u)
	return uuo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRole(u *user.Role) *UserUpdateOne {
	if u != nil {
		uuo.SetRole(*u)
	}
	return uuo
}

// SetActivatedOn sets the "activated_on" field.
func (uuo *UserUpdateOne) SetActivatedOn(t time.Time) *UserUpdateOne {
	uuo.mutation.SetActivatedOn(t)
	return uuo
}

// SetNillableActivatedOn sets the "activated_on" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableActivatedOn(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetActivatedOn(*t)
	}
	return uuo
}

// ClearActivatedOn clears the value of the "activated_on" field.
func (uuo *UserUpdateOne) ClearActivatedOn() *UserUpdateOne {
	uuo.mutation.ClearActivatedOn()
	return uuo
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (uuo *UserUpdateOne) SetGroupID(id uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetGroupID(id)
	return uuo
}

// SetGroup sets the "group" edge to the Group entity.
func (uuo *UserUpdateOne) SetGroup(g *Group) *UserUpdateOne {
	return uuo.SetGroupID(g.ID)
}

// AddAuthTokenIDs adds the "auth_tokens" edge to the AuthTokens entity by IDs.
func (uuo *UserUpdateOne) AddAuthTokenIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddAuthTokenIDs(ids...)
	return uuo
}

// AddAuthTokens adds the "auth_tokens" edges to the AuthTokens entity.
func (uuo *UserUpdateOne) AddAuthTokens(a ...*AuthTokens) *UserUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.AddAuthTokenIDs(ids...)
}

// AddNotifierIDs adds the "notifiers" edge to the Notifier entity by IDs.
func (uuo *UserUpdateOne) AddNotifierIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddNotifierIDs(ids...)
	return uuo
}

// AddNotifiers adds the "notifiers" edges to the Notifier entity.
func (uuo *UserUpdateOne) AddNotifiers(n ...*Notifier) *UserUpdateOne {
	ids := make([]uuid.UUID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return uuo.AddNotifierIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearGroup clears the "group" edge to the Group entity.
func (uuo *UserUpdateOne) ClearGroup() *UserUpdateOne {
	uuo.mutation.ClearGroup()
	return uuo
}

// ClearAuthTokens clears all "auth_tokens" edges to the AuthTokens entity.
func (uuo *UserUpdateOne) ClearAuthTokens() *UserUpdateOne {
	uuo.mutation.ClearAuthTokens()
	return uuo
}

// RemoveAuthTokenIDs removes the "auth_tokens" edge to AuthTokens entities by IDs.
func (uuo *UserUpdateOne) RemoveAuthTokenIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveAuthTokenIDs(ids...)
	return uuo
}

// RemoveAuthTokens removes "auth_tokens" edges to AuthTokens entities.
func (uuo *UserUpdateOne) RemoveAuthTokens(a ...*AuthTokens) *UserUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.RemoveAuthTokenIDs(ids...)
}

// ClearNotifiers clears all "notifiers" edges to the Notifier entity.
func (uuo *UserUpdateOne) ClearNotifiers() *UserUpdateOne {
	uuo.mutation.ClearNotifiers()
	return uuo
}

// RemoveNotifierIDs removes the "notifiers" edge to Notifier entities by IDs.
func (uuo *UserUpdateOne) RemoveNotifierIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveNotifierIDs(ids...)
	return uuo
}

// RemoveNotifiers removes "notifiers" edges to Notifier entities.
func (uuo *UserUpdateOne) RemoveNotifiers(n ...*Notifier) *UserUpdateOne {
	ids := make([]uuid.UUID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return uuo.RemoveNotifierIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	uuo.defaults()
	return withHooks[*User, UserMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Role(); ok {
		if err := user.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "User.role": %w`, err)}
		}
	}
	if _, ok := uuo.mutation.GroupID(); uuo.mutation.GroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "User.group"`)
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.IsSuperuser(); ok {
		_spec.SetField(user.FieldIsSuperuser, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.Superuser(); ok {
		_spec.SetField(user.FieldSuperuser, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeEnum, value)
	}
	if value, ok := uuo.mutation.ActivatedOn(); ok {
		_spec.SetField(user.FieldActivatedOn, field.TypeTime, value)
	}
	if uuo.mutation.ActivatedOnCleared() {
		_spec.ClearField(user.FieldActivatedOn, field.TypeTime)
	}
	if uuo.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.GroupTable,
			Columns: []string{user.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.GroupTable,
			Columns: []string{user.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.AuthTokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AuthTokensTable,
			Columns: []string{user.AuthTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authtokens.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedAuthTokensIDs(); len(nodes) > 0 && !uuo.mutation.AuthTokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AuthTokensTable,
			Columns: []string{user.AuthTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authtokens.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.AuthTokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AuthTokensTable,
			Columns: []string{user.AuthTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authtokens.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.NotifiersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.NotifiersTable,
			Columns: []string{user.NotifiersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifier.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedNotifiersIDs(); len(nodes) > 0 && !uuo.mutation.NotifiersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.NotifiersTable,
			Columns: []string{user.NotifiersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifier.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.NotifiersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.NotifiersTable,
			Columns: []string{user.NotifiersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifier.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
