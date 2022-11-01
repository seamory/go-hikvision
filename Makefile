default:

windows:
	go build -o build.exe ./cmd/main.go

clean:
	go clean -r -cache -testcache
