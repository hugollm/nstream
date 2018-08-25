pkg = ./...


run:
	go run server.go

format:
	go fmt $(pkg)

test:
	@go test -v $(pkg) | $(nicer_test_output)

watch-tests:
	@watch --color "go test -v $(pkg) | $(nicer_test_output)"

coverage:
	go test -coverprofile=/tmp/golang-coverage-report ./...
	go tool cover -html=/tmp/golang-coverage-report

database:
	sudo -u postgres psql -c "DROP DATABASE IF EXISTS nstream"
	sudo -u postgres psql -c "CREATE DATABASE nstream"
	sudo -u postgres psql nstream -f schema.sql


nicer_test_output = sed "/RUN/d" | sed "/PASS/s//$(green)/" | sed "/FAIL/s//$(red)/"
green = $(shell printf "\033[32mPASS\033[0m")
red = $(shell printf "\033[31mFAIL\033[0m")
