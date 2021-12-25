// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/internationalization/pkg/db/ent/lang"
	"github.com/google/uuid"
)

// LangCreate is the builder for creating a Lang entity.
type LangCreate struct {
	config
	mutation *LangMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetLang sets the "lang" field.
func (lc *LangCreate) SetLang(s string) *LangCreate {
	lc.mutation.SetLang(s)
	return lc
}

// SetCreateAt sets the "create_at" field.
func (lc *LangCreate) SetCreateAt(u uint32) *LangCreate {
	lc.mutation.SetCreateAt(u)
	return lc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (lc *LangCreate) SetNillableCreateAt(u *uint32) *LangCreate {
	if u != nil {
		lc.SetCreateAt(*u)
	}
	return lc
}

// SetUpdateAt sets the "update_at" field.
func (lc *LangCreate) SetUpdateAt(u uint32) *LangCreate {
	lc.mutation.SetUpdateAt(u)
	return lc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (lc *LangCreate) SetNillableUpdateAt(u *uint32) *LangCreate {
	if u != nil {
		lc.SetUpdateAt(*u)
	}
	return lc
}

// SetDeleteAt sets the "delete_at" field.
func (lc *LangCreate) SetDeleteAt(u uint32) *LangCreate {
	lc.mutation.SetDeleteAt(u)
	return lc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (lc *LangCreate) SetNillableDeleteAt(u *uint32) *LangCreate {
	if u != nil {
		lc.SetDeleteAt(*u)
	}
	return lc
}

// SetID sets the "id" field.
func (lc *LangCreate) SetID(u uuid.UUID) *LangCreate {
	lc.mutation.SetID(u)
	return lc
}

// Mutation returns the LangMutation object of the builder.
func (lc *LangCreate) Mutation() *LangMutation {
	return lc.mutation
}

// Save creates the Lang in the database.
func (lc *LangCreate) Save(ctx context.Context) (*Lang, error) {
	var (
		err  error
		node *Lang
	)
	lc.defaults()
	if len(lc.hooks) == 0 {
		if err = lc.check(); err != nil {
			return nil, err
		}
		node, err = lc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LangMutation)
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
func (lc *LangCreate) SaveX(ctx context.Context) *Lang {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LangCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LangCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LangCreate) defaults() {
	if _, ok := lc.mutation.CreateAt(); !ok {
		v := lang.DefaultCreateAt()
		lc.mutation.SetCreateAt(v)
	}
	if _, ok := lc.mutation.UpdateAt(); !ok {
		v := lang.DefaultUpdateAt()
		lc.mutation.SetUpdateAt(v)
	}
	if _, ok := lc.mutation.DeleteAt(); !ok {
		v := lang.DefaultDeleteAt()
		lc.mutation.SetDeleteAt(v)
	}
	if _, ok := lc.mutation.ID(); !ok {
		v := lang.DefaultID()
		lc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LangCreate) check() error {
	if _, ok := lc.mutation.Lang(); !ok {
		return &ValidationError{Name: "lang", err: errors.New(`ent: missing required field "lang"`)}
	}
	if _, ok := lc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := lc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "update_at"`)}
	}
	if _, ok := lc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "delete_at"`)}
	}
	return nil
}

func (lc *LangCreate) sqlSave(ctx context.Context) (*Lang, error) {
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (lc *LangCreate) createSpec() (*Lang, *sqlgraph.CreateSpec) {
	var (
		_node = &Lang{config: lc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: lang.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: lang.FieldID,
			},
		}
	)
	_spec.OnConflict = lc.conflict
	if id, ok := lc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := lc.mutation.Lang(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lang.FieldLang,
		})
		_node.Lang = value
	}
	if value, ok := lc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: lang.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := lc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: lang.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := lc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: lang.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Lang.Create().
//		SetLang(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LangUpsert) {
//			SetLang(v+v).
//		}).
//		Exec(ctx)
//
func (lc *LangCreate) OnConflict(opts ...sql.ConflictOption) *LangUpsertOne {
	lc.conflict = opts
	return &LangUpsertOne{
		create: lc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Lang.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (lc *LangCreate) OnConflictColumns(columns ...string) *LangUpsertOne {
	lc.conflict = append(lc.conflict, sql.ConflictColumns(columns...))
	return &LangUpsertOne{
		create: lc,
	}
}

type (
	// LangUpsertOne is the builder for "upsert"-ing
	//  one Lang node.
	LangUpsertOne struct {
		create *LangCreate
	}

	// LangUpsert is the "OnConflict" setter.
	LangUpsert struct {
		*sql.UpdateSet
	}
)

// SetLang sets the "lang" field.
func (u *LangUpsert) SetLang(v string) *LangUpsert {
	u.Set(lang.FieldLang, v)
	return u
}

// UpdateLang sets the "lang" field to the value that was provided on create.
func (u *LangUpsert) UpdateLang() *LangUpsert {
	u.SetExcluded(lang.FieldLang)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *LangUpsert) SetCreateAt(v uint32) *LangUpsert {
	u.Set(lang.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *LangUpsert) UpdateCreateAt() *LangUpsert {
	u.SetExcluded(lang.FieldCreateAt)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *LangUpsert) SetUpdateAt(v uint32) *LangUpsert {
	u.Set(lang.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *LangUpsert) UpdateUpdateAt() *LangUpsert {
	u.SetExcluded(lang.FieldUpdateAt)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *LangUpsert) SetDeleteAt(v uint32) *LangUpsert {
	u.Set(lang.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *LangUpsert) UpdateDeleteAt() *LangUpsert {
	u.SetExcluded(lang.FieldDeleteAt)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Lang.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(lang.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *LangUpsertOne) UpdateNewValues() *LangUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(lang.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Lang.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *LangUpsertOne) Ignore() *LangUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LangUpsertOne) DoNothing() *LangUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LangCreate.OnConflict
// documentation for more info.
func (u *LangUpsertOne) Update(set func(*LangUpsert)) *LangUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LangUpsert{UpdateSet: update})
	}))
	return u
}

// SetLang sets the "lang" field.
func (u *LangUpsertOne) SetLang(v string) *LangUpsertOne {
	return u.Update(func(s *LangUpsert) {
		s.SetLang(v)
	})
}

// UpdateLang sets the "lang" field to the value that was provided on create.
func (u *LangUpsertOne) UpdateLang() *LangUpsertOne {
	return u.Update(func(s *LangUpsert) {
		s.UpdateLang()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *LangUpsertOne) SetCreateAt(v uint32) *LangUpsertOne {
	return u.Update(func(s *LangUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *LangUpsertOne) UpdateCreateAt() *LangUpsertOne {
	return u.Update(func(s *LangUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *LangUpsertOne) SetUpdateAt(v uint32) *LangUpsertOne {
	return u.Update(func(s *LangUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *LangUpsertOne) UpdateUpdateAt() *LangUpsertOne {
	return u.Update(func(s *LangUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *LangUpsertOne) SetDeleteAt(v uint32) *LangUpsertOne {
	return u.Update(func(s *LangUpsert) {
		s.SetDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *LangUpsertOne) UpdateDeleteAt() *LangUpsertOne {
	return u.Update(func(s *LangUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *LangUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LangCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LangUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *LangUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: LangUpsertOne.ID is not supported by MySQL driver. Use LangUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *LangUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// LangCreateBulk is the builder for creating many Lang entities in bulk.
type LangCreateBulk struct {
	config
	builders []*LangCreate
	conflict []sql.ConflictOption
}

// Save creates the Lang entities in the database.
func (lcb *LangCreateBulk) Save(ctx context.Context) ([]*Lang, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Lang, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LangMutation)
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
					spec.OnConflict = lcb.conflict
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
func (lcb *LangCreateBulk) SaveX(ctx context.Context) []*Lang {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LangCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LangCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Lang.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LangUpsert) {
//			SetLang(v+v).
//		}).
//		Exec(ctx)
//
func (lcb *LangCreateBulk) OnConflict(opts ...sql.ConflictOption) *LangUpsertBulk {
	lcb.conflict = opts
	return &LangUpsertBulk{
		create: lcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Lang.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (lcb *LangCreateBulk) OnConflictColumns(columns ...string) *LangUpsertBulk {
	lcb.conflict = append(lcb.conflict, sql.ConflictColumns(columns...))
	return &LangUpsertBulk{
		create: lcb,
	}
}

// LangUpsertBulk is the builder for "upsert"-ing
// a bulk of Lang nodes.
type LangUpsertBulk struct {
	create *LangCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Lang.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(lang.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *LangUpsertBulk) UpdateNewValues() *LangUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(lang.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Lang.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *LangUpsertBulk) Ignore() *LangUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LangUpsertBulk) DoNothing() *LangUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LangCreateBulk.OnConflict
// documentation for more info.
func (u *LangUpsertBulk) Update(set func(*LangUpsert)) *LangUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LangUpsert{UpdateSet: update})
	}))
	return u
}

// SetLang sets the "lang" field.
func (u *LangUpsertBulk) SetLang(v string) *LangUpsertBulk {
	return u.Update(func(s *LangUpsert) {
		s.SetLang(v)
	})
}

// UpdateLang sets the "lang" field to the value that was provided on create.
func (u *LangUpsertBulk) UpdateLang() *LangUpsertBulk {
	return u.Update(func(s *LangUpsert) {
		s.UpdateLang()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *LangUpsertBulk) SetCreateAt(v uint32) *LangUpsertBulk {
	return u.Update(func(s *LangUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *LangUpsertBulk) UpdateCreateAt() *LangUpsertBulk {
	return u.Update(func(s *LangUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *LangUpsertBulk) SetUpdateAt(v uint32) *LangUpsertBulk {
	return u.Update(func(s *LangUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *LangUpsertBulk) UpdateUpdateAt() *LangUpsertBulk {
	return u.Update(func(s *LangUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *LangUpsertBulk) SetDeleteAt(v uint32) *LangUpsertBulk {
	return u.Update(func(s *LangUpsert) {
		s.SetDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *LangUpsertBulk) UpdateDeleteAt() *LangUpsertBulk {
	return u.Update(func(s *LangUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *LangUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the LangCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LangCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LangUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
