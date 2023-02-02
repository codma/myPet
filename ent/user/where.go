// Code generated by ent, DO NOT EDIT.

package user

import (
	"main/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// IsUsed applies equality check predicate on the "is_used" field. It's identical to IsUsedEQ.
func IsUsed(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldIsUsed, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUserID, v))
}

// UserName applies equality check predicate on the "user_name" field. It's identical to UserNameEQ.
func UserName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUserName, v))
}

// Birthday applies equality check predicate on the "birthday" field. It's identical to BirthdayEQ.
func Birthday(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldBirthday, v))
}

// Weight applies equality check predicate on the "weight" field. It's identical to WeightEQ.
func Weight(v float64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldWeight, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldType, v))
}

// StoreID applies equality check predicate on the "store_id" field. It's identical to StoreIDEQ.
func StoreID(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldStoreID, v))
}

// IsUsedEQ applies the EQ predicate on the "is_used" field.
func IsUsedEQ(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldIsUsed, v))
}

// IsUsedNEQ applies the NEQ predicate on the "is_used" field.
func IsUsedNEQ(v bool) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldIsUsed, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUserID, v))
}

// UserNameEQ applies the EQ predicate on the "user_name" field.
func UserNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUserName, v))
}

// UserNameNEQ applies the NEQ predicate on the "user_name" field.
func UserNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUserName, v))
}

// UserNameIn applies the In predicate on the "user_name" field.
func UserNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUserName, vs...))
}

// UserNameNotIn applies the NotIn predicate on the "user_name" field.
func UserNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUserName, vs...))
}

// UserNameGT applies the GT predicate on the "user_name" field.
func UserNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUserName, v))
}

// UserNameGTE applies the GTE predicate on the "user_name" field.
func UserNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUserName, v))
}

// UserNameLT applies the LT predicate on the "user_name" field.
func UserNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUserName, v))
}

// UserNameLTE applies the LTE predicate on the "user_name" field.
func UserNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUserName, v))
}

// UserNameContains applies the Contains predicate on the "user_name" field.
func UserNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUserName, v))
}

// UserNameHasPrefix applies the HasPrefix predicate on the "user_name" field.
func UserNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUserName, v))
}

// UserNameHasSuffix applies the HasSuffix predicate on the "user_name" field.
func UserNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUserName, v))
}

// UserNameEqualFold applies the EqualFold predicate on the "user_name" field.
func UserNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUserName, v))
}

// UserNameContainsFold applies the ContainsFold predicate on the "user_name" field.
func UserNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUserName, v))
}

// BirthdayEQ applies the EQ predicate on the "birthday" field.
func BirthdayEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldBirthday, v))
}

// BirthdayNEQ applies the NEQ predicate on the "birthday" field.
func BirthdayNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldBirthday, v))
}

// BirthdayIn applies the In predicate on the "birthday" field.
func BirthdayIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldBirthday, vs...))
}

// BirthdayNotIn applies the NotIn predicate on the "birthday" field.
func BirthdayNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldBirthday, vs...))
}

// BirthdayGT applies the GT predicate on the "birthday" field.
func BirthdayGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldBirthday, v))
}

// BirthdayGTE applies the GTE predicate on the "birthday" field.
func BirthdayGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldBirthday, v))
}

// BirthdayLT applies the LT predicate on the "birthday" field.
func BirthdayLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldBirthday, v))
}

// BirthdayLTE applies the LTE predicate on the "birthday" field.
func BirthdayLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldBirthday, v))
}

// BirthdayContains applies the Contains predicate on the "birthday" field.
func BirthdayContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldBirthday, v))
}

// BirthdayHasPrefix applies the HasPrefix predicate on the "birthday" field.
func BirthdayHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldBirthday, v))
}

// BirthdayHasSuffix applies the HasSuffix predicate on the "birthday" field.
func BirthdayHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldBirthday, v))
}

// BirthdayEqualFold applies the EqualFold predicate on the "birthday" field.
func BirthdayEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldBirthday, v))
}

// BirthdayContainsFold applies the ContainsFold predicate on the "birthday" field.
func BirthdayContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldBirthday, v))
}

// WeightEQ applies the EQ predicate on the "weight" field.
func WeightEQ(v float64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldWeight, v))
}

// WeightNEQ applies the NEQ predicate on the "weight" field.
func WeightNEQ(v float64) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldWeight, v))
}

// WeightIn applies the In predicate on the "weight" field.
func WeightIn(vs ...float64) predicate.User {
	return predicate.User(sql.FieldIn(FieldWeight, vs...))
}

// WeightNotIn applies the NotIn predicate on the "weight" field.
func WeightNotIn(vs ...float64) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldWeight, vs...))
}

// WeightGT applies the GT predicate on the "weight" field.
func WeightGT(v float64) predicate.User {
	return predicate.User(sql.FieldGT(FieldWeight, v))
}

// WeightGTE applies the GTE predicate on the "weight" field.
func WeightGTE(v float64) predicate.User {
	return predicate.User(sql.FieldGTE(FieldWeight, v))
}

// WeightLT applies the LT predicate on the "weight" field.
func WeightLT(v float64) predicate.User {
	return predicate.User(sql.FieldLT(FieldWeight, v))
}

// WeightLTE applies the LTE predicate on the "weight" field.
func WeightLTE(v float64) predicate.User {
	return predicate.User(sql.FieldLTE(FieldWeight, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldType, v))
}

// StoreIDEQ applies the EQ predicate on the "store_id" field.
func StoreIDEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldStoreID, v))
}

// StoreIDNEQ applies the NEQ predicate on the "store_id" field.
func StoreIDNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldStoreID, v))
}

// StoreIDIn applies the In predicate on the "store_id" field.
func StoreIDIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldStoreID, vs...))
}

// StoreIDNotIn applies the NotIn predicate on the "store_id" field.
func StoreIDNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldStoreID, vs...))
}

// StoreIDGT applies the GT predicate on the "store_id" field.
func StoreIDGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldStoreID, v))
}

// StoreIDGTE applies the GTE predicate on the "store_id" field.
func StoreIDGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldStoreID, v))
}

// StoreIDLT applies the LT predicate on the "store_id" field.
func StoreIDLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldStoreID, v))
}

// StoreIDLTE applies the LTE predicate on the "store_id" field.
func StoreIDLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldStoreID, v))
}

// StoreIDContains applies the Contains predicate on the "store_id" field.
func StoreIDContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldStoreID, v))
}

// StoreIDHasPrefix applies the HasPrefix predicate on the "store_id" field.
func StoreIDHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldStoreID, v))
}

// StoreIDHasSuffix applies the HasSuffix predicate on the "store_id" field.
func StoreIDHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldStoreID, v))
}

// StoreIDEqualFold applies the EqualFold predicate on the "store_id" field.
func StoreIDEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldStoreID, v))
}

// StoreIDContainsFold applies the ContainsFold predicate on the "store_id" field.
func StoreIDContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldStoreID, v))
}

// HasMystore applies the HasEdge predicate on the "mystore" edge.
func HasMystore() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, MystoreTable, MystorePrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMystoreWith applies the HasEdge predicate on the "mystore" edge with a given conditions (other predicates).
func HasMystoreWith(preds ...predicate.Store) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MystoreInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, MystoreTable, MystorePrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
