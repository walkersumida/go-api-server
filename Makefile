ogen-gen:
	go generate ./ogen

ent-new-%:
	go run -mod=mod entgo.io/ent/cmd/ent new ${@:ent-new-%=%}

ent-gen:
	mkdir -p ./tmp/ent
	mv ./ent/schema ./tmp/ent/
	rm -rf ./ent
	mv ./tmp/ent ./ent
	rm -rf ./tmp
	go run ./cmd/ent generate

db-migrate-up:
	go run ./cmd/db migrate up

db-migrate-down:
	go run ./cmd/db migrate down

db-migration-create-%: ent-gen
	go run ./cmd/db migration create ${@:db-migration-create-%=%}

db-create:
	go run ./cmd/db create

db-drop:
	go run ./cmd/db drop

start:
	go run ./cmd/server
