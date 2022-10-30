## 懸念点

- Q: 中間テーブルの名称変更について
- A: 以下のようにTo側からStorageKeyを使うことによって可能
```go
edge.To("category", Category.Type).
        StorageKey(edge.Table("j_user_with_category")), 
```

- Q: FromとToの違いは何?
- A: エッジでそれぞれのノードをつなぐという考え方なので、ユーザー、とチームを例に挙げて考えると
  - Teamノードから見たらuserというエッジを利用してUserノードに繋ぐ **To**
  - Userノードから見たらteamという名前のエッジを利用してTeamノードに繋ぐ(逆向きのエッジ作成)  **From**

  こちらをすることで、お互いに参照するような関係ができる?

  <small>※ 逆向きのエッジを作る際には、`Ref()`で参照するエッジを指定する。
  O2Mの場合は、From側にUniqueがつき、O2Oの場合は、両方につく<small>

```golang
// Edges of the User.
func (User) Edges() []ent.Edge {
        return []ent.Edge{
                // O2M(One側)
                // ユーザーは、0~1つのチームに属する
				// (必須の場合は、EdgesにRequired()がつき、必要ない場合は、FieldsにOptional()をつける)
                edge.From("team", Team.Type).
                        Ref("user").
                        Unique().
                        Field("team_id"),
        }
}
```

```golang
// Edges of the Team.
func (Team) Edges() []ent.Edge {
	return []ent.Edge{
		// O2M(Many側)
		edge.To("user", User.Type),
	}
}
```

- Q: 外部キー制約を持つカラムに対して、変更を加えるとき、一旦外部キーを削除して貼り直すという動作をしない(マイグレーション時にそのままエラー)
-A:  https://github.com/ent/ent/issues/2388 ここで指摘がある

  現状は、sql生成後に削除を手動で挿入する、もしくは、hookで生成の処理に割り込ませる、該当のリレーションを一度コメントアウトするなどが考えられる(中間テーブル消える?)

  (現在、キー制約が存在するかを確認して、存在すれば落として、再生成という処理が実現可能か調査中)

 ** ※ 1対多の時は問題なく貼り直してくれるので、中間テーブルの挙動回りで確認が必要**

- Q: あとからテーブル名を変更すると新規作成扱いになる(rename tableのサポート)
- A: 調査中

```go
func (Category) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "goal",
		},
	}
}
```

- Q: https://github.com/ent/ent/issues/252 ソフトデリートについて
- A: 調査中(普通のアップデート構文でできるので特に問題ない?)

```
User(id=1, role_id=0, team_id=0, goal_id=1, user_name=taka, email=example@example.com, user_icon=, user_memo=, created_at=Sun Oct 30 14:41:43 2022, updated_at=Sun Oct 30 14:41:43 2022, deleted_at=Sun Oct 30 14:41:43 2022)
```

- Q: https://zenn.dev/taxio/articles/32274944ae61fd 意図しないUPDATE/DELETE ALLが発生する可能性について
- A: mysqlの--safe-updatesを使用

- Q: https://github.com/ent/ent/issues/1986 テーブル名の単数系サポート
- A: Annotationsを使う
