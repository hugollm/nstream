run:
	go run server.go

format:
	go fmt ./...

watch-tests:
	watch -n 1 go test -cover ./...

coverage:
	go test -coverprofile=/tmp/golang-coverage-report ./...
	go tool cover -html=/tmp/golang-coverage-report
