now = $$(date +%s)

dep:
	go get -u github.com/julienschmidt/httprouter
	go get -u github.com/Fs02/kamimai/cmd/kamimai

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
