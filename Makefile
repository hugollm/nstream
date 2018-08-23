run:
	go run server.go

format:
	go fmt ./...

watch-tests:
	watch -n 1 go test ./...
