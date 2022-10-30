package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sample/migrations/ent"
	"sample/migrations/ent/user"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	erHandler := func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("er.html")
		if err != nil {
			panic(err.Error())
		}
		if err := t.Execute(w, nil); err != nil {
			panic(err.Error())
		}

		ctx := context.Background()
		client := conn()
		defer client.Close()
		// ソフトデリート挙動確認
		softDeleteSample(ctx, client)
	}

	// er図生成ツールは二種類ある
	// https://github.com/a8m/enter
	// https://github.com/hedwigz/entviz
	http.HandleFunc("/er", erHandler)
	http.ListenAndServe("localhost:3002", nil /*ent.ServeEntviz()*/)
}

func softDeleteSample(ctx context.Context, client *ent.Client) {
	entGoal, _ := createGoal(ctx, client)
	entUser, _ := createUser(ctx, client, entGoal.ID)
	fmt.Print(entUser.String())
	_ = deleteUser(ctx, client, 1) //　ソフトデリート

	entUserInfo, _ := getUser(ctx, client, 1)

	for _, userInfo := range entUserInfo {
		fmt.Printf("%s", userInfo.String())
	}

}

func conn() *ent.Client {
	client, err := ent.Open("mysql", "user:password@tcp(localhost:3306)/sample?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client

}

func createGoal(ctx context.Context, client *ent.Client) (*ent.Goal, error) {
	currentTime := time.Now()
	entGoal, err := client.
		Debug().
		Goal.
		Create().
		SetGoal("foobar").
		SetCreatedAt(currentTime).
		SetUpdatedAt(currentTime).
		Save(ctx) // saveXは実行できない場合にパニックを起こす
	if err != nil {
		log.Fatalf(err.Error())
	}
	return entGoal, err
}

func createUser(ctx context.Context, client *ent.Client, goalId int) (*ent.User, error) {
	currentTime := time.Now()
	entUser, err := client.
		Debug().
		User.
		Create().
		SetUserName("taka").
		SetEmail("example@example.com").
		SetGoalID(goalId).
		SetCreatedAt(currentTime).
		SetUpdatedAt(currentTime).
		Save(ctx) // saveXは実行できない場合にパニックを起こす
	if err != nil {
		log.Fatalf(err.Error())
	}
	return entUser, err
}

func deleteUser(ctx context.Context, client *ent.Client, userId int) error {
	currentTime := time.Now()
	_, err := client.
		Debug().
		User.
		Update().
		SetDeletedAt(currentTime).
		Where(user.ID(userId)).
		Save(ctx) // saveXは実行できない場合にパニックを起こす
	if err != nil {
		log.Fatalf(err.Error())
	}

	return err
}

func getUser(ctx context.Context, client *ent.Client, userId int) ([]*ent.User, error) {
	user, err := client.
		Debug().
		User.
		Query().
		Select(user.Columns...).
		Where(user.ID(userId)).
		All(ctx)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return user, err
}
