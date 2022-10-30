package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StorageKey("user_id").
			SchemaType(map[string]string{
				dialect.MySQL: "int",
			}),
		field.Int("role_id").
			Optional(),
		field.Int("team_id").
			Optional(),
		field.Int("goal_id"),
		field.String("user_name").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(20)",
			}),
		field.String("email").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(127)",
			}),
		field.String("user_icon").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(127)",
			}).
			Optional(),
		field.String("user_memo").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(300)",
			}).
			Optional(),
		field.Time("created_at"),
		field.Time("updated_at"),
		field.Time("deleted_at").
			Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// M2M
		// ユーザーは0~Mのカテゴリーを持つ
		// 中間テーブルの名称変更
		// https://github.com/ent/ent/issues/2621#issuecomment-1169636111
		// PRIMARY KEY (`user_id`, `category_id`) 複合主キーから単一の主キーを追加する
		// https://github.com/ent/ent/issues/438#issuecomment-616066789
		edge.To("category", Category.Type).
			StorageKey(edge.Table("j_user_with_category")),
		// O2M(One側)
		// ユーザーは0~1のロールを持つ
		edge.From("role", Role.Type).
			Ref("user").
			Unique().
			Field("role_id"),
		// O2M(One側)
		// ユーザーは、0~1つのチームに属する
		// (必須の場合は、EdgesにRequired()がつき、必要ない場合は、FieldsにOptional()をつける)
		edge.From("team", Team.Type).
			Ref("user").
			Unique().
			Field("team_id"),
		// O2O
		// ユーザーは、一つの目標を持つ
		edge.From("goal", Goal.Type).
			Ref("user").
			Unique().
			Field("goal_id").
			Required(),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "user",
		},
	}
}
