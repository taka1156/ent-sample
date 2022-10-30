package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StorageKey("role_id").
			SchemaType(map[string]string{
				dialect.MySQL: "int",
			}),
		field.String("role_name").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(127)",
			}),
		field.String("role_icon").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(127)",
			}),
		field.String("role_memo").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(300)",
			}),
		field.Time("created_at"),
		field.Time("updated_at"),
		field.Time("deleted_at").
			Optional(),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		// O2M(Many側)
		edge.To("user", User.Type),
	}
}

func (Role) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "role",
		},
	}
}
