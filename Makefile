now = $$(date +%s)

init:
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	go get -u github.com/Fs02/kamimai/cmd/kamimai
	dep ensure

build:
	go build -o bin/beasiswakita app/main.go

migrate:
	export $$(cat .env | grep -v ^\# | xargs) && \
	kamimai --driver=mysql --dsn="mysql://$$DEVELOPMENT_DATABASE_URL" --directory=./db/migrations sync

rollback:
	export $$(cat .env | grep -v ^\# | xargs) && \
	kamimai --driver=mysql --dsn="mysql://$$DEVELOPMENT_DATABASE_URL" --directory=./db/migrations down

migration:
	@touch ./db/migrations/$(now)_$(filter-out $@,$(MAKECMDGOALS))_up.sql
	@touch ./db/migrations/$(now)_$(filter-out $@,$(MAKECMDGOALS))_down.sql

%:
	@:
