// Code generated by ent, DO NOT EDIT.

package category

import (
	"sample/migrations/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CategoryName applies equality check predicate on the "category_name" field. It's identical to CategoryNameEQ.
func CategoryName(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCategoryName), v))
	})
}

// CategoryIcon applies equality check predicate on the "category_icon" field. It's identical to CategoryIconEQ.
func CategoryIcon(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCategoryIcon), v))
	})
}

// CategoryMemo applies equality check predicate on the "category_memo" field. It's identical to CategoryMemoEQ.
func CategoryMemo(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCategoryMemo), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// CategoryNameEQ applies the EQ predicate on the "category_name" field.
func CategoryNameEQ(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCategoryName), v))
	})
}

// CategoryNameNEQ applies the NEQ predicate on the "category_name" field.
func CategoryNameNEQ(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCategoryName), v))
	})
}

// CategoryNameIn applies the In predicate on the "category_name" field.
func CategoryNameIn(vs ...string) predicate.Category {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCategoryName), v...))
	})
}

// CategoryNameNotIn applies the NotIn predicate on the "category_name" field.
func CategoryNameNotIn(vs ...string) predicate.Category {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCategoryName), v...))
	})
}

// CategoryNameGT applies the GT predicate on the "category_name" field.
func CategoryNameGT(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCategoryName), v))
	})
}

// CategoryNameGTE applies the GTE predicate on the "category_name" field.
func CategoryNameGTE(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCategoryName), v))
	})
}

// CategoryNameLT applies the LT predicate on the "category_name" field.
func CategoryNameLT(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCategoryName), v))
	})
}

// CategoryNameLTE applies the LTE predicate on the "category_name" field.
func CategoryNameLTE(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCategoryName), v))
	})
}

// CategoryNameContains applies the Contains predicate on the "category_name" field.
func CategoryNameContains(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCategoryName), v))
	})
}

// CategoryNameHasPrefix applies the HasPrefix predicate on the "category_name" field.
func CategoryNameHasPrefix(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCategoryName), v))
	})
}

// CategoryNameHasSuffix applies the HasSuffix predicate on the "category_name" field.
func CategoryNameHasSuffix(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCategoryName), v))
	})
}

// CategoryNameEqualFold applies the EqualFold predicate on the "category_name" field.
func CategoryNameEqualFold(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCategoryName), v))
	})
}

// CategoryNameContainsFold applies the ContainsFold predicate on the "category_name" field.
func CategoryNameContainsFold(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCategoryName), v))
	})
}

// CategoryIconEQ applies the EQ predicate on the "category_icon" field.
func CategoryIconEQ(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCategoryIcon), v))
	})
}

// CategoryIconNEQ applies the NEQ predicate on the "category_icon" field.
func CategoryIconNEQ(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCategoryIcon), v))
	})
}

// CategoryIconIn applies the In predicate on the "category_icon" field.
func CategoryIconIn(vs ...string) predicate.Category {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCategoryIcon), v...))
	})
}

// CategoryIconNotIn applies the NotIn predicate on the "category_icon" field.
func CategoryIconNotIn(vs ...string) predicate.Category {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCategoryIcon), v...))
	})
}

// CategoryIconGT applies the GT predicate on the "category_icon" field.
func CategoryIconGT(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCategoryIcon), v))
	})
}

// CategoryIconGTE applies the GTE predicate on the "category_icon" field.
func CategoryIconGTE(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCategoryIcon), v))
	})
}

// CategoryIconLT applies the LT predicate on the "category_icon" field.
func CategoryIconLT(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCategoryIcon), v))
	})
}

// CategoryIconLTE applies the LTE predicate on the "category_icon" field.
func CategoryIconLTE(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCategoryIcon), v))
	})
}

// CategoryIconContains applies the Contains predicate on the "category_icon" field.
func CategoryIconContains(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCategoryIcon), v))
	})
}

// CategoryIconHasPrefix applies the HasPrefix predicate on the "category_icon" field.
func CategoryIconHasPrefix(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCategoryIcon), v))
	})
}

// CategoryIconHasSuffix applies the HasSuffix predicate on the "category_icon" field.
func CategoryIconHasSuffix(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCategoryIcon), v))
	})
}

// CategoryIconEqualFold applies the EqualFold predicate on the "category_icon" field.
func CategoryIconEqualFold(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCategoryIcon), v))
	})
}

// CategoryIconContainsFold applies the ContainsFold predicate on the "category_icon" field.
func CategoryIconContainsFold(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCategoryIcon), v))
	})
}

// CategoryMemoEQ applies the EQ predicate on the "category_memo" field.
func CategoryMemoEQ(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCategoryMemo), v))
	})
}

// CategoryMemoNEQ applies the NEQ predicate on the "category_memo" field.
func CategoryMemoNEQ(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCategoryMemo), v))
	})
}

// CategoryMemoIn applies the In predicate on the "category_memo" field.
func CategoryMemoIn(vs ...string) predicate.Category {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCategoryMemo), v...))
	})
}

// CategoryMemoNotIn applies the NotIn predicate on the "category_memo" field.
func CategoryMemoNotIn(vs ...string) predicate.Category {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCategoryMemo), v...))
	})
}

// CategoryMemoGT applies the GT predicate on the "category_memo" field.
func CategoryMemoGT(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCategoryMemo), v))
	})
}

// CategoryMemoGTE applies the GTE predicate on the "category_memo" field.
func CategoryMemoGTE(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCategoryMemo), v))
	})
}

// CategoryMemoLT applies the LT predicate on the "category_memo" field.
func CategoryMemoLT(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCategoryMemo), v))
	})
}

// CategoryMemoLTE applies the LTE predicate on the "category_memo" field.
func CategoryMemoLTE(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCategoryMemo), v))
	})
}

// CategoryMemoContains applies the Contains predicate on the "category_memo" field.
func CategoryMemoContains(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCategoryMemo), v))
	})
}

// CategoryMemoHasPrefix applies the HasPrefix predicate on the "category_memo" field.
func CategoryMemoHasPrefix(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCategoryMemo), v))
	})
}

// CategoryMemoHasSuffix applies the HasSuffix predicate on the "category_memo" field.
func CategoryMemoHasSuffix(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCategoryMemo), v))
	})
}

// CategoryMemoEqualFold applies the EqualFold predicate on the "category_memo" field.
func CategoryMemoEqualFold(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCategoryMemo), v))
	})
}

// CategoryMemoContainsFold applies the ContainsFold predicate on the "category_memo" field.
func CategoryMemoContainsFold(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCategoryMemo), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Category {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Category {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Category {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Category {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Category {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Category {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, UserFieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UserTable, UserPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, UserFieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UserTable, UserPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Category) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Category) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
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
func Not(p predicate.Category) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		p(s.Not())
	})
}
