package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UsedMixin{},
		TimeMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(), //seq
		field.String("user_id"),     // 사용자 아이디
		field.String("user_name"),   // 반려묘/견 이름
		field.String("birthday"),    // 생일
		field.Float("weight"),       //몸무게
		field.String("type"),        //종류
		field.String("store_id"),    //스토어아이디
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("mystore", Store.Type).
			Ref("users"),
	}
}

//사용자아이디로 구매내역을 쌓음
//createday로 구매일을 확인
//주기계산은 사용자 별 카테고리별 심바 - 사료 / 심바 - 간식
