package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Team holds the schema definition for the Team entity.
type Team struct {
	ent.Schema
}

// Fields of the Team.
func (Team) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StorageKey("team_id").
			SchemaType(map[string]string{
				dialect.MySQL: "int",
			}),
		field.String("team_name").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(127)",
			}),
		field.String("team_icon").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(127)",
			}),
		field.String("team_memo").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(300)",
			}),
		field.Time("created_at"),
		field.Time("updated_at"),
		field.Time("deleted_at").
			Optional(),
	}
}

// Edges of the Team.
func (Team) Edges() []ent.Edge {
	return []ent.Edge{
		// O2M(Many側)
		edge.To("user", User.Type),
	}
}

func (Team) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "team",
		},
	}
}
