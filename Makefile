run:
	go run server.go

format:
	go fmt ./...

watch-tests:
	watch -n 1 go test -cover ./...

coverage:
	go test -coverprofile=/tmp/golang-coverage-report ./...
	go tool cover -html=/tmp/golang-coverage-report

database:
	sudo -u postgres psql -c "DROP DATABASE IF EXISTS nstream"
	sudo -u postgres psql -c "CREATE DATABASE nstream"
	sudo -u postgres psql nstream -f schema.sql
