// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/blob"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BlobCreate is the builder for creating a Blob entity.
type BlobCreate struct {
	config
	mutation *BlobMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (bc *BlobCreate) SetUUID(u uuid.UUID) *BlobCreate {
	bc.mutation.SetUUID(u)
	return bc
}

// SetID sets the "id" field.
func (bc *BlobCreate) SetID(u uuid.UUID) *BlobCreate {
	bc.mutation.SetID(u)
	return bc
}

// SetParentID sets the "parent" edge to the Blob entity by ID.
func (bc *BlobCreate) SetParentID(id uuid.UUID) *BlobCreate {
	bc.mutation.SetParentID(id)
	return bc
}

// SetNillableParentID sets the "parent" edge to the Blob entity by ID if the given value is not nil.
func (bc *BlobCreate) SetNillableParentID(id *uuid.UUID) *BlobCreate {
	if id != nil {
		bc = bc.SetParentID(*id)
	}
	return bc
}

// SetParent sets the "parent" edge to the Blob entity.
func (bc *BlobCreate) SetParent(b *Blob) *BlobCreate {
	return bc.SetParentID(b.ID)
}

// AddLinkIDs adds the "links" edge to the Blob entity by IDs.
func (bc *BlobCreate) AddLinkIDs(ids ...uuid.UUID) *BlobCreate {
	bc.mutation.AddLinkIDs(ids...)
	return bc
}

// AddLinks adds the "links" edges to the Blob entity.
func (bc *BlobCreate) AddLinks(b ...*Blob) *BlobCreate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bc.AddLinkIDs(ids...)
}

// Mutation returns the BlobMutation object of the builder.
func (bc *BlobCreate) Mutation() *BlobMutation {
	return bc.mutation
}

// Save creates the Blob in the database.
func (bc *BlobCreate) Save(ctx context.Context) (*Blob, error) {
	var (
		err  error
		node *Blob
	)
	bc.defaults()
	if len(bc.hooks) == 0 {
		if err = bc.check(); err != nil {
			return nil, err
		}
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BlobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bc.check(); err != nil {
				return nil, err
			}
			bc.mutation = mutation
			if node, err = bc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BlobCreate) SaveX(ctx context.Context) *Blob {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (bc *BlobCreate) defaults() {
	if _, ok := bc.mutation.UUID(); !ok {
		v := blob.DefaultUUID()
		bc.mutation.SetUUID(v)
	}
	if _, ok := bc.mutation.ID(); !ok {
		v := blob.DefaultID()
		bc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BlobCreate) check() error {
	if _, ok := bc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New("ent: missing required field \"uuid\"")}
	}
	return nil
}

func (bc *BlobCreate) sqlSave(ctx context.Context) (*Blob, error) {
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (bc *BlobCreate) createSpec() (*Blob, *sqlgraph.CreateSpec) {
	var (
		_node = &Blob{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: blob.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: blob.FieldID,
			},
		}
	)
	if id, ok := bc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := bc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: blob.FieldUUID,
		})
		_node.UUID = value
	}
	if nodes := bc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   blob.ParentTable,
			Columns: []string{blob.ParentColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: blob.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.blob_parent = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.LinksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   blob.LinksTable,
			Columns: blob.LinksPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: blob.FieldID,
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

// BlobCreateBulk is the builder for creating many Blob entities in bulk.
type BlobCreateBulk struct {
	config
	builders []*BlobCreate
}

// Save creates the Blob entities in the database.
func (bcb *BlobCreateBulk) Save(ctx context.Context) ([]*Blob, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Blob, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BlobMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BlobCreateBulk) SaveX(ctx context.Context) []*Blob {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
