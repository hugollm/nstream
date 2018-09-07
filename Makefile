pkg = ./...


run:
	go run server.go

format:
	go fmt $(pkg)

test:
	@go clean -testcache
	@go test -v $(pkg) | $(nicer_test_output)

watch-tests:
	@watch --color "go test -v $(pkg) | $(nicer_test_output)"

coverage:
	go test -coverprofile=/tmp/golang-coverage-report ./...
	go tool cover -html=/tmp/golang-coverage-report

database: pg-hba-notice
	sudo -u postgres psql -c "DROP DATABASE IF EXISTS nstream"
	sudo -u postgres psql -c "DROP ROLE IF EXISTS nstream"
	sudo -u postgres psql -c "CREATE ROLE nstream ENCRYPTED PASSWORD 'nstream' LOGIN"
	sudo -u postgres psql -c "CREATE DATABASE nstream OWNER nstream"
	PGPASSWORD=nstream psql -U nstream nstream -f data/schema.sql

schema: pg-hba-notice
	PGPASSWORD=nstream psql -U nstream nstream -f data/schema.sql

pg-hba-notice:
	@echo "\n  NOTE: requires entry on pg_hba.conf: local all all md5\n"


nicer_test_output = sed "/RUN/d" | sed "/PASS/s//$(green)/" | sed "/FAIL/s//$(red)/"
green = $(shell printf "\033[32mPASS\033[0m")
red = $(shell printf "\033[31mFAIL\033[0m")
