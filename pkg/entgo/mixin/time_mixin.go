package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// CreateTime adds created at time field.
type CreateTime struct{ ent.Schema }

// Fields of the creation time mixin.
func (CreateTime) Fields() []ent.Field {
	return []ent.Field{
		field.Time("create_time").
			Default(time.Now).
			Immutable(),
	}
}

// create time mixin must implement `Mixin` interface.
var _ ent.Mixin = (*CreateTime)(nil)

// UpdateTime adds updated at time field.
type UpdateTime struct{ ent.Schema }

// Fields of the update time mixin.
func (UpdateTime) Fields() []ent.Field {
	return []ent.Field{
		field.Time("update_time").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// update time mixin must implement `Mixin` interface.
var _ ent.Mixin = (*UpdateTime)(nil)

// Time composes create/update time mixin.
type Time struct{ ent.Schema }

// Fields of the time mixin.
func (Time) Fields() []ent.Field {
	return append(
		CreateTime{}.Fields(),
		UpdateTime{}.Fields()...,
	)
}

// time mixin must implement `Mixin` interface.
var _ ent.Mixin = (*Time)(nil)
