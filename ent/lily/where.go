// Code generated by entc, DO NOT EDIT.

package lily

import (
	"study-event-go/ent/predicate"
	"study-event-go/types"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id types.LilyID) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id types.LilyID) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id types.LilyID) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...types.LilyID) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...types.LilyID) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id types.LilyID) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id types.LilyID) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id types.LilyID) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id types.LilyID) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletionReason applies equality check predicate on the "deletion_reason" field. It's identical to DeletionReasonEQ.
func DeletionReason(v types.DeletionReason) predicate.Lily {
	vc := uint32(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletionReason), vc))
	})
}

// FirstName applies equality check predicate on the "first_name" field. It's identical to FirstNameEQ.
func FirstName(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFirstName), v))
	})
}

// MiddleName applies equality check predicate on the "middle_name" field. It's identical to MiddleNameEQ.
func MiddleName(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMiddleName), v))
	})
}

// LastName applies equality check predicate on the "last_name" field. It's identical to LastNameEQ.
func LastName(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastName), v))
	})
}

// Rank applies equality check predicate on the "rank" field. It's identical to RankEQ.
func Rank(v uint32) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRank), v))
	})
}

// MainCharmID applies equality check predicate on the "main_charm_id" field. It's identical to MainCharmIDEQ.
func MainCharmID(v types.CharmID) predicate.Lily {
	vc := uint64(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMainCharmID), vc))
	})
}

// SubCharmID applies equality check predicate on the "sub_charm_id" field. It's identical to SubCharmIDEQ.
func SubCharmID(v types.CharmID) predicate.Lily {
	vc := uint64(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSubCharmID), vc))
	})
}

// GardenID applies equality check predicate on the "garden_id" field. It's identical to GardenIDEQ.
func GardenID(v types.GardenID) predicate.Lily {
	vc := uint64(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGardenID), vc))
	})
}

// LegionID applies equality check predicate on the "legion_id" field. It's identical to LegionIDEQ.
func LegionID(v types.LegionID) predicate.Lily {
	vc := uint64(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLegionID), vc))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Lily {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Lily(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Lily {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Lily(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Lily {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Lily(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Lily {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Lily(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Lily {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Lily(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Lily {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Lily(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// DeletionReasonEQ applies the EQ predicate on the "deletion_reason" field.
func DeletionReasonEQ(v types.DeletionReason) predicate.Lily {
	vc := uint32(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletionReason), vc))
	})
}

// DeletionReasonNEQ applies the NEQ predicate on the "deletion_reason" field.
func DeletionReasonNEQ(v types.DeletionReason) predicate.Lily {
	vc := uint32(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletionReason), vc))
	})
}

// DeletionReasonIn applies the In predicate on the "deletion_reason" field.
func DeletionReasonIn(vs ...types.DeletionReason) predicate.Lily {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = uint32(vs[i])
	}
	return predicate.Lily(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeletionReason), v...))
	})
}

// DeletionReasonNotIn applies the NotIn predicate on the "deletion_reason" field.
func DeletionReasonNotIn(vs ...types.DeletionReason) predicate.Lily {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = uint32(vs[i])
	}
	return predicate.Lily(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeletionReason), v...))
	})
}

// DeletionReasonGT applies the GT predicate on the "deletion_reason" field.
func DeletionReasonGT(v types.DeletionReason) predicate.Lily {
	vc := uint32(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletionReason), vc))
	})
}

// DeletionReasonGTE applies the GTE predicate on the "deletion_reason" field.
func DeletionReasonGTE(v types.DeletionReason) predicate.Lily {
	vc := uint32(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletionReason), vc))
	})
}

// DeletionReasonLT applies the LT predicate on the "deletion_reason" field.
func DeletionReasonLT(v types.DeletionReason) predicate.Lily {
	vc := uint32(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletionReason), vc))
	})
}

// DeletionReasonLTE applies the LTE predicate on the "deletion_reason" field.
func DeletionReasonLTE(v types.DeletionReason) predicate.Lily {
	vc := uint32(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletionReason), vc))
	})
}

// FirstNameEQ applies the EQ predicate on the "first_name" field.
func FirstNameEQ(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFirstName), v))
	})
}

// FirstNameNEQ applies the NEQ predicate on the "first_name" field.
func FirstNameNEQ(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFirstName), v))
	})
}

// FirstNameIn applies the In predicate on the "first_name" field.
func FirstNameIn(vs ...string) predicate.Lily {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Lily(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFirstName), v...))
	})
}

// FirstNameNotIn applies the NotIn predicate on the "first_name" field.
func FirstNameNotIn(vs ...string) predicate.Lily {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Lily(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFirstName), v...))
	})
}

// FirstNameGT applies the GT predicate on the "first_name" field.
func FirstNameGT(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFirstName), v))
	})
}

// FirstNameGTE applies the GTE predicate on the "first_name" field.
func FirstNameGTE(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFirstName), v))
	})
}

// FirstNameLT applies the LT predicate on the "first_name" field.
func FirstNameLT(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFirstName), v))
	})
}

// FirstNameLTE applies the LTE predicate on the "first_name" field.
func FirstNameLTE(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFirstName), v))
	})
}

// FirstNameContains applies the Contains predicate on the "first_name" field.
func FirstNameContains(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFirstName), v))
	})
}

// FirstNameHasPrefix applies the HasPrefix predicate on the "first_name" field.
func FirstNameHasPrefix(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFirstName), v))
	})
}

// FirstNameHasSuffix applies the HasSuffix predicate on the "first_name" field.
func FirstNameHasSuffix(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFirstName), v))
	})
}

// FirstNameEqualFold applies the EqualFold predicate on the "first_name" field.
func FirstNameEqualFold(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFirstName), v))
	})
}

// FirstNameContainsFold applies the ContainsFold predicate on the "first_name" field.
func FirstNameContainsFold(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFirstName), v))
	})
}

// MiddleNameEQ applies the EQ predicate on the "middle_name" field.
func MiddleNameEQ(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMiddleName), v))
	})
}

// MiddleNameNEQ applies the NEQ predicate on the "middle_name" field.
func MiddleNameNEQ(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMiddleName), v))
	})
}

// MiddleNameIn applies the In predicate on the "middle_name" field.
func MiddleNameIn(vs ...string) predicate.Lily {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Lily(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMiddleName), v...))
	})
}

// MiddleNameNotIn applies the NotIn predicate on the "middle_name" field.
func MiddleNameNotIn(vs ...string) predicate.Lily {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Lily(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMiddleName), v...))
	})
}

// MiddleNameGT applies the GT predicate on the "middle_name" field.
func MiddleNameGT(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMiddleName), v))
	})
}

// MiddleNameGTE applies the GTE predicate on the "middle_name" field.
func MiddleNameGTE(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMiddleName), v))
	})
}

// MiddleNameLT applies the LT predicate on the "middle_name" field.
func MiddleNameLT(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMiddleName), v))
	})
}

// MiddleNameLTE applies the LTE predicate on the "middle_name" field.
func MiddleNameLTE(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMiddleName), v))
	})
}

// MiddleNameContains applies the Contains predicate on the "middle_name" field.
func MiddleNameContains(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMiddleName), v))
	})
}

// MiddleNameHasPrefix applies the HasPrefix predicate on the "middle_name" field.
func MiddleNameHasPrefix(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMiddleName), v))
	})
}

// MiddleNameHasSuffix applies the HasSuffix predicate on the "middle_name" field.
func MiddleNameHasSuffix(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMiddleName), v))
	})
}

// MiddleNameEqualFold applies the EqualFold predicate on the "middle_name" field.
func MiddleNameEqualFold(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMiddleName), v))
	})
}

// MiddleNameContainsFold applies the ContainsFold predicate on the "middle_name" field.
func MiddleNameContainsFold(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMiddleName), v))
	})
}

// LastNameEQ applies the EQ predicate on the "last_name" field.
func LastNameEQ(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastName), v))
	})
}

// LastNameNEQ applies the NEQ predicate on the "last_name" field.
func LastNameNEQ(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastName), v))
	})
}

// LastNameIn applies the In predicate on the "last_name" field.
func LastNameIn(vs ...string) predicate.Lily {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Lily(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLastName), v...))
	})
}

// LastNameNotIn applies the NotIn predicate on the "last_name" field.
func LastNameNotIn(vs ...string) predicate.Lily {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Lily(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLastName), v...))
	})
}

// LastNameGT applies the GT predicate on the "last_name" field.
func LastNameGT(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastName), v))
	})
}

// LastNameGTE applies the GTE predicate on the "last_name" field.
func LastNameGTE(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastName), v))
	})
}

// LastNameLT applies the LT predicate on the "last_name" field.
func LastNameLT(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastName), v))
	})
}

// LastNameLTE applies the LTE predicate on the "last_name" field.
func LastNameLTE(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastName), v))
	})
}

// LastNameContains applies the Contains predicate on the "last_name" field.
func LastNameContains(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLastName), v))
	})
}

// LastNameHasPrefix applies the HasPrefix predicate on the "last_name" field.
func LastNameHasPrefix(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLastName), v))
	})
}

// LastNameHasSuffix applies the HasSuffix predicate on the "last_name" field.
func LastNameHasSuffix(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLastName), v))
	})
}

// LastNameEqualFold applies the EqualFold predicate on the "last_name" field.
func LastNameEqualFold(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLastName), v))
	})
}

// LastNameContainsFold applies the ContainsFold predicate on the "last_name" field.
func LastNameContainsFold(v string) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLastName), v))
	})
}

// RankEQ applies the EQ predicate on the "rank" field.
func RankEQ(v uint32) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRank), v))
	})
}

// RankNEQ applies the NEQ predicate on the "rank" field.
func RankNEQ(v uint32) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRank), v))
	})
}

// RankIn applies the In predicate on the "rank" field.
func RankIn(vs ...uint32) predicate.Lily {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Lily(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRank), v...))
	})
}

// RankNotIn applies the NotIn predicate on the "rank" field.
func RankNotIn(vs ...uint32) predicate.Lily {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Lily(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRank), v...))
	})
}

// RankGT applies the GT predicate on the "rank" field.
func RankGT(v uint32) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRank), v))
	})
}

// RankGTE applies the GTE predicate on the "rank" field.
func RankGTE(v uint32) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRank), v))
	})
}

// RankLT applies the LT predicate on the "rank" field.
func RankLT(v uint32) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRank), v))
	})
}

// RankLTE applies the LTE predicate on the "rank" field.
func RankLTE(v uint32) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRank), v))
	})
}

// MainCharmIDEQ applies the EQ predicate on the "main_charm_id" field.
func MainCharmIDEQ(v types.CharmID) predicate.Lily {
	vc := uint64(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMainCharmID), vc))
	})
}

// MainCharmIDNEQ applies the NEQ predicate on the "main_charm_id" field.
func MainCharmIDNEQ(v types.CharmID) predicate.Lily {
	vc := uint64(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMainCharmID), vc))
	})
}

// MainCharmIDIn applies the In predicate on the "main_charm_id" field.
func MainCharmIDIn(vs ...types.CharmID) predicate.Lily {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.Lily(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMainCharmID), v...))
	})
}

// MainCharmIDNotIn applies the NotIn predicate on the "main_charm_id" field.
func MainCharmIDNotIn(vs ...types.CharmID) predicate.Lily {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.Lily(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMainCharmID), v...))
	})
}

// MainCharmIDGT applies the GT predicate on the "main_charm_id" field.
func MainCharmIDGT(v types.CharmID) predicate.Lily {
	vc := uint64(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMainCharmID), vc))
	})
}

// MainCharmIDGTE applies the GTE predicate on the "main_charm_id" field.
func MainCharmIDGTE(v types.CharmID) predicate.Lily {
	vc := uint64(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMainCharmID), vc))
	})
}

// MainCharmIDLT applies the LT predicate on the "main_charm_id" field.
func MainCharmIDLT(v types.CharmID) predicate.Lily {
	vc := uint64(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMainCharmID), vc))
	})
}

// MainCharmIDLTE applies the LTE predicate on the "main_charm_id" field.
func MainCharmIDLTE(v types.CharmID) predicate.Lily {
	vc := uint64(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMainCharmID), vc))
	})
}

// SubCharmIDEQ applies the EQ predicate on the "sub_charm_id" field.
func SubCharmIDEQ(v types.CharmID) predicate.Lily {
	vc := uint64(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSubCharmID), vc))
	})
}

// SubCharmIDNEQ applies the NEQ predicate on the "sub_charm_id" field.
func SubCharmIDNEQ(v types.CharmID) predicate.Lily {
	vc := uint64(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSubCharmID), vc))
	})
}

// SubCharmIDIn applies the In predicate on the "sub_charm_id" field.
func SubCharmIDIn(vs ...types.CharmID) predicate.Lily {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.Lily(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSubCharmID), v...))
	})
}

// SubCharmIDNotIn applies the NotIn predicate on the "sub_charm_id" field.
func SubCharmIDNotIn(vs ...types.CharmID) predicate.Lily {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.Lily(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSubCharmID), v...))
	})
}

// SubCharmIDGT applies the GT predicate on the "sub_charm_id" field.
func SubCharmIDGT(v types.CharmID) predicate.Lily {
	vc := uint64(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSubCharmID), vc))
	})
}

// SubCharmIDGTE applies the GTE predicate on the "sub_charm_id" field.
func SubCharmIDGTE(v types.CharmID) predicate.Lily {
	vc := uint64(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSubCharmID), vc))
	})
}

// SubCharmIDLT applies the LT predicate on the "sub_charm_id" field.
func SubCharmIDLT(v types.CharmID) predicate.Lily {
	vc := uint64(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSubCharmID), vc))
	})
}

// SubCharmIDLTE applies the LTE predicate on the "sub_charm_id" field.
func SubCharmIDLTE(v types.CharmID) predicate.Lily {
	vc := uint64(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSubCharmID), vc))
	})
}

// GardenIDEQ applies the EQ predicate on the "garden_id" field.
func GardenIDEQ(v types.GardenID) predicate.Lily {
	vc := uint64(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGardenID), vc))
	})
}

// GardenIDNEQ applies the NEQ predicate on the "garden_id" field.
func GardenIDNEQ(v types.GardenID) predicate.Lily {
	vc := uint64(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGardenID), vc))
	})
}

// GardenIDIn applies the In predicate on the "garden_id" field.
func GardenIDIn(vs ...types.GardenID) predicate.Lily {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.Lily(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGardenID), v...))
	})
}

// GardenIDNotIn applies the NotIn predicate on the "garden_id" field.
func GardenIDNotIn(vs ...types.GardenID) predicate.Lily {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.Lily(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGardenID), v...))
	})
}

// GardenIDGT applies the GT predicate on the "garden_id" field.
func GardenIDGT(v types.GardenID) predicate.Lily {
	vc := uint64(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGardenID), vc))
	})
}

// GardenIDGTE applies the GTE predicate on the "garden_id" field.
func GardenIDGTE(v types.GardenID) predicate.Lily {
	vc := uint64(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGardenID), vc))
	})
}

// GardenIDLT applies the LT predicate on the "garden_id" field.
func GardenIDLT(v types.GardenID) predicate.Lily {
	vc := uint64(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGardenID), vc))
	})
}

// GardenIDLTE applies the LTE predicate on the "garden_id" field.
func GardenIDLTE(v types.GardenID) predicate.Lily {
	vc := uint64(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGardenID), vc))
	})
}

// LegionIDEQ applies the EQ predicate on the "legion_id" field.
func LegionIDEQ(v types.LegionID) predicate.Lily {
	vc := uint64(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLegionID), vc))
	})
}

// LegionIDNEQ applies the NEQ predicate on the "legion_id" field.
func LegionIDNEQ(v types.LegionID) predicate.Lily {
	vc := uint64(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLegionID), vc))
	})
}

// LegionIDIn applies the In predicate on the "legion_id" field.
func LegionIDIn(vs ...types.LegionID) predicate.Lily {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.Lily(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLegionID), v...))
	})
}

// LegionIDNotIn applies the NotIn predicate on the "legion_id" field.
func LegionIDNotIn(vs ...types.LegionID) predicate.Lily {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.Lily(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLegionID), v...))
	})
}

// LegionIDGT applies the GT predicate on the "legion_id" field.
func LegionIDGT(v types.LegionID) predicate.Lily {
	vc := uint64(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLegionID), vc))
	})
}

// LegionIDGTE applies the GTE predicate on the "legion_id" field.
func LegionIDGTE(v types.LegionID) predicate.Lily {
	vc := uint64(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLegionID), vc))
	})
}

// LegionIDLT applies the LT predicate on the "legion_id" field.
func LegionIDLT(v types.LegionID) predicate.Lily {
	vc := uint64(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLegionID), vc))
	})
}

// LegionIDLTE applies the LTE predicate on the "legion_id" field.
func LegionIDLTE(v types.LegionID) predicate.Lily {
	vc := uint64(v)
	return predicate.Lily(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLegionID), vc))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Lily) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Lily) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
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
func Not(p predicate.Lily) predicate.Lily {
	return predicate.Lily(func(s *sql.Selector) {
		p(s.Not())
	})
}
