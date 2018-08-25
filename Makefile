pkg = ./...

run:
	go run server.go

format:
	go fmt $(pkg)


watch-tests:
	watch -n 1 go test -v -cover $(pkg)

coverage:
	go test -coverprofile=/tmp/golang-coverage-report ./...
	go tool cover -html=/tmp/golang-coverage-report

database:
	sudo -u postgres psql -c "DROP DATABASE IF EXISTS nstream"
	sudo -u postgres psql -c "CREATE DATABASE nstream"
	sudo -u postgres psql nstream -f schema.sql
