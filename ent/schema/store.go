package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Store holds the schema definition for the Store entity.
type Store struct {
	ent.Schema
}

func (Store) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UsedMixin{},
		TimeMixin{},
	}
}

// Fields of the Store.
func (Store) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),    //seq
		field.String("store_id"),       // 스토어 아이디
		field.String("app_access_key"), // 앱 엑세스 키
	}
}

// Edges of the Store.
func (Store) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
	}
}
