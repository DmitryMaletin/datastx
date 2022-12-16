test:
	go test ./internal/...

repl:
	go run cmd/main.go

purge-branches:
	git branch | grep -v "main" | xargs git branch -D 