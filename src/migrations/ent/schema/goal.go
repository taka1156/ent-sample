package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Goal holds the schema definition for the Goal entity.
type Goal struct {
	ent.Schema
}

// Fields of the Goal.
func (Goal) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StorageKey("goal_id").
			SchemaType(map[string]string{
				dialect.MySQL: "int",
			}),
		field.String("goal").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(300)",
			}),
		field.Time("created_at"),
		field.Time("updated_at"),
		field.Time("deleted_at").
			Optional(),
	}
}

// Edges of the Goal.
func (Goal) Edges() []ent.Edge {
	return []ent.Edge{
		// O2O
		edge.To("user", User.Type).
			Unique(),
	}
}

func (Goal) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "goal",
		},
	}
}
