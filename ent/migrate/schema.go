// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CharmsColumns holds the columns for the "charms" table.
	CharmsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "model_id", Type: field.TypeUint64},
		{Name: "owner_id", Type: field.TypeUint64},
	}
	// CharmsTable holds the schema information for the "charms" table.
	CharmsTable = &schema.Table{
		Name:       "charms",
		Columns:    CharmsColumns,
		PrimaryKey: []*schema.Column{CharmsColumns[0]},
	}
	// CharmCreatorsColumns holds the columns for the "charm_creators" table.
	CharmCreatorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "type", Type: field.TypeUint32},
	}
	// CharmCreatorsTable holds the schema information for the "charm_creators" table.
	CharmCreatorsTable = &schema.Table{
		Name:       "charm_creators",
		Columns:    CharmCreatorsColumns,
		PrimaryKey: []*schema.Column{CharmCreatorsColumns[0]},
	}
	// CharmModelsColumns holds the columns for the "charm_models" table.
	CharmModelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "creator_id", Type: field.TypeUint64},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "type", Type: field.TypeUint32},
		{Name: "generation", Type: field.TypeUint32},
	}
	// CharmModelsTable holds the schema information for the "charm_models" table.
	CharmModelsTable = &schema.Table{
		Name:       "charm_models",
		Columns:    CharmModelsColumns,
		PrimaryKey: []*schema.Column{CharmModelsColumns[0]},
	}
	// GardensColumns holds the columns for the "gardens" table.
	GardensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "location", Type: field.TypeString},
		{Name: "mentorship_id", Type: field.TypeUint64},
		{Name: "legion_system", Type: field.TypeUint32},
	}
	// GardensTable holds the schema information for the "gardens" table.
	GardensTable = &schema.Table{
		Name:       "gardens",
		Columns:    GardensColumns,
		PrimaryKey: []*schema.Column{GardensColumns[0]},
	}
	// LiliesColumns holds the columns for the "lilies" table.
	LiliesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "cause_of_deletion", Type: field.TypeUint32, Default: 0},
		{Name: "birth", Type: field.TypeTime, Nullable: true},
		{Name: "first_name", Type: field.TypeString},
		{Name: "middle_name", Type: field.TypeString},
		{Name: "last_name", Type: field.TypeString},
		{Name: "rank", Type: field.TypeUint32, Default: 0},
		{Name: "garden_id", Type: field.TypeUint64, Default: 0},
		{Name: "legion_id", Type: field.TypeUint64, Default: 0},
	}
	// LiliesTable holds the schema information for the "lilies" table.
	LiliesTable = &schema.Table{
		Name:       "lilies",
		Columns:    LiliesColumns,
		PrimaryKey: []*schema.Column{LiliesColumns[0]},
	}
	// LilySkillsColumns holds the columns for the "lily_skills" table.
	LilySkillsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "lily_id", Type: field.TypeUint64},
		{Name: "skill_id", Type: field.TypeUint64},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// LilySkillsTable holds the schema information for the "lily_skills" table.
	LilySkillsTable = &schema.Table{
		Name:       "lily_skills",
		Columns:    LilySkillsColumns,
		PrimaryKey: []*schema.Column{LilySkillsColumns[0]},
	}
	// MentorshipsColumns holds the columns for the "mentorships" table.
	MentorshipsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// MentorshipsTable holds the schema information for the "mentorships" table.
	MentorshipsTable = &schema.Table{
		Name:       "mentorships",
		Columns:    MentorshipsColumns,
		PrimaryKey: []*schema.Column{MentorshipsColumns[0]},
	}
	// SkillsColumns holds the columns for the "skills" table.
	SkillsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "type", Type: field.TypeUint32, Default: 0},
	}
	// SkillsTable holds the schema information for the "skills" table.
	SkillsTable = &schema.Table{
		Name:       "skills",
		Columns:    SkillsColumns,
		PrimaryKey: []*schema.Column{SkillsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CharmsTable,
		CharmCreatorsTable,
		CharmModelsTable,
		GardensTable,
		LiliesTable,
		LilySkillsTable,
		MentorshipsTable,
		SkillsTable,
	}
)

func init() {
}
