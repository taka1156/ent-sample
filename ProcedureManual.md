## スキーマ作成~sql生成までの手順

1. こちらでスキーマの雛形作成
    `go run -mod=mod entgo.io/ent/cmd/ent init`

2. 以下のようなものが生成されるので、こちらに追記
    ```go
    package schema

    import "entgo.io/ent"

    // User holds the schema definition for the User entity.
    type User struct {
        ent.Schema
    }

    // Fields of the User.
    func (User) Fields() []ent.Field {
        return nil
    }

    // Edges of the User.
    func (User) Edges() []ent.Edge {
        return nil
    }
    
    ```

    追記例
    ```go
    // Edges of the User.
    func (User) Edges() []ent.Edge {
        return []ent.Edge{
            // M2M
            // ユーザーは0~Mのカテゴリーを持つ
            // 中間テーブルの名称変更
            // https://github.com/ent/ent/issues/2621#issuecomment-1169636111
            // PRIMARY KEY (`user_id`, `category_id`) 複合主キーから単一の主キーに変更する(今はできない?)
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
                Field("team_id").
                Required(),
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
    ```

3. `ent/generate.go`の`go:generate`部分に`--feature sql/versioned-migration`オプションがついてなければ、追記する。
   ```go
    package ent

    //go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/versioned-migration ./schema
   ```

4. [Versioned Migrations - official document](https://entgo.io/docs/versioned-migrations)を参考に、ソースコードをコピペ

> 2. Create a main.go file under the migrate/ent package and customize the migration generation for your project.
上記のgolang-migrate/migrateタブをクリック

```go
//go:build ignore

package main

import (
    "context"
    "log"
    "os"

    "<project>/ent/migrate"

    "ariga.io/atlas/sql/sqltool"
    "entgo.io/ent/dialect"
    "entgo.io/ent/dialect/sql/schema"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    ctx := context.Background()
    // Create a local migration directory able to understand golang-migrate migration file format for replay.
    dir, err := sqltool.NewGolangMigrateDir("ent/migrate/migrations")
    if err != nil {
        log.Fatalf("failed creating atlas migration directory: %v", err)
    }
    // Migrate diff options.
    opts := []schema.MigrateOption{
        schema.WithDir(dir),                         // provide migration directory
        schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
        schema.WithDialect(dialect.MySQL),           // Ent dialect to use
    }
    if len(os.Args) != 2 {
        log.Fatalln("migration name is required. Use: 'go run -mod=mod ent/migrate/main.go <name>'")
    }
    // Generate migrations using Atlas support for MySQL (note the Ent dialect option passed above).
    err = migrate.NamedDiff(ctx, "ここはローカルの環境に合わせる:mysql://root:pass@localhost:3306/test", os.Args[1], opts...)
    if err != nil {
        log.Fatalf("failed generating migration file: %v", err)
    }
}
```

5. 以下で、SQLが生成(その時点の全部のテーブルの変更や追加がまとめて反映される)
`go run -mod=mod ent/migrate/main.go 名前`

6. 生成されたSQLを実行

## その他使えそうなコマンド

`go run -mod=mod entgo.io/ent/cmd/ent describe schemaディレクトリまでのパス`を実行

以下のようにスキーマ一覧が見れます。

```
User:
        +------------+-----------+--------+----------+----------+---------+---------------+-----------+-----------------------------+------------+
        |   Field    |   Type    | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |          StructTag          | Validators |
        +------------+-----------+--------+----------+----------+---------+---------------+-----------+-----------------------------+------------+
        | id         | int       | false  | false    | false    | false   | false         | false     | json:"id,omitempty"         |          0 |
        | role_id    | int       | false  | true     | false    | false   | false         | false     | json:"role_id,omitempty"    |          0 |
        | team_id    | int       | false  | false    | false    | false   | false         | false     | json:"team_id,omitempty"    |          0 |
        | goal_id    | int       | false  | false    | false    | false   | false         | false     | json:"goal_id,omitempty"    |          0 |
        | name       | string    | false  | false    | false    | false   | false         | false     | json:"name,omitempty"       |          0 |
        | email      | string    | false  | false    | false    | false   | false         | false     | json:"email,omitempty"      |          0 |
        | user_icon  | string    | false  | false    | false    | false   | false         | false     | json:"user_icon,omitempty"  |          0 |
        | user_memo  | string    | false  | false    | false    | false   | false         | false     | json:"user_memo,omitempty"  |          0 |
        | created_at | time.Time | false  | false    | false    | false   | false         | false     | json:"created_at,omitempty" |          0 |
        | updated_at | time.Time | false  | false    | false    | false   | false         | false     | json:"updated_at,omitempty" |          0 |
        | deleted_at | time.Time | false  | false    | false    | false   | false         | false     | json:"deleted_at,omitempty" |          0 |
        +------------+-----------+--------+----------+----------+---------+---------------+-----------+-----------------------------+------------+
        +----------+----------+---------+---------+----------+--------+----------+
        |   Edge   |   Type   | Inverse | BackRef | Relation | Unique | Optional |
        +----------+----------+---------+---------+----------+--------+----------+
        | category | Category | false   |         | M2M      | false  | true     |
        | role     | Role     | true    | user    | M2O      | true   | true     |
        | team     | Team     | true    | user    | M2O      | true   | false    |
        | goal     | Goal     | true    | user    | O2O      | true   | false    |
        +----------+----------+---------+---------+----------+--------+----------+
```

`go get -u github.com/a8m/enter`をインストール
`go  run -mod=mod github.com/a8m/enter schemaディレクトリまでのパス`を実行

このようなER図をHTMLとして吐き出してくれます。

<img width="584" alt="スクリーンショット 2022-10-30 12 02 35" src="https://user-images.githubusercontent.com/47517002/198860620-c121d127-ed92-4541-bc7b-d767f31b164d.png">

## makefile

```Makefile
name=sample
MYSQL_URL="mysql://user:password@tcp(127.0.0.1:3306)/sample"

# スキーマの一覧、構造確認
ent-list:
	go run -mod=mod entgo.io/ent/cmd/ent describe ./migrations/ent/schema

# ER図吐き出し(html)
ent-ergen:
    go  run -mod=mod github.com/a8m/enter ./migrations/ent/schema

# 新規スキーマ生成
ent-schema:
		go run -mod=mod entgo.io/ent/cmd/ent init --target ./migrations/ent/schema --template ./migrations/ent/templ/schema.tmpl ${name}

# sql、その他関連コード生成
ent-gen:
	go generate ./migrations/ent
	go run -mod=mod ./migrations/ent/migrate/main.go ${name}

# 用途不明?
sqllint:
	go run -mod=mod ariga.io/atlas/cmd/atlas@master migrate lint \
	--dev-url="mysql://user:password@localhost:3306/sample" \
	--dir="file://migrations/ent/migrate/sql" \
	--dir-format="golang-migrate" \
	--latest=1

# 生成したsqlを走らせる
migrate:
	migrate -path ./migrations/ent/migrate/sql -database ${MYSQL_URL} up

migrate-create:
	migrate create -ext sql -dir ./migrations/ent/migrate/sql -tz asia/tokyo ${name}
```
