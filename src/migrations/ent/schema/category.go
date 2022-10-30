package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StorageKey("category_id").
			SchemaType(map[string]string{
				dialect.MySQL: "int",
			}),
		field.String("category_name").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(127)",
			}),
		field.String("category_icon").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(127)",
			}),
		field.String("category_memo").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(300)",
			}),
		field.Time("created_at"),
		field.Time("updated_at"),
		field.Time("deleted_at").
			Optional(),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		// M2M
		// カテゴリーは、0~Mのユーザーに参照される
		edge.From("user", User.Type).
			Ref("category"),
	}
}

func (Category) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "category",
		},
	}
}
