name=sample
MYSQL_URL="mysql://user:password@tcp(127.0.0.1:3306)/sample"

ent-list:
	go run -mod=mod entgo.io/ent/cmd/ent describe ./migrations/ent/schema

ent-ergen:
    go  run -mod=mod github.com/a8m/enter ./migrations/ent/schema

ent-schema:
	go run -mod=mod entgo.io/ent/cmd/ent init --target ./migrations/ent/schema --template ./migrations/ent/templ/schema.tmpl ${name}

ent-gen:
	go generate ./migrations/ent
	go run -mod=mod ./migrations/ent/migrate/main.go ${name}

sqllint:
	go run -mod=mod ariga.io/atlas/cmd/atlas@master migrate lint \
	--dev-url="mysql://user:password@localhost:3306/sample" \
	--dir="file://migrations/ent/migrate/sql" \
	--dir-format="golang-migrate" \
	--latest=1

migrate:
	migrate -path ./migrations/ent/migrate/sql -database ${MYSQL_URL} up

migrate-create:
	migrate create -ext sql -dir ./migrations/ent/migrate/sql -tz asia/tokyo ${name}
