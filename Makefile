build:
	go build -o bin/github-licenses-report main.go
compile:
	GOOS=linux GOARCH=arm go build -o bin/github-licenses-report-linux-arm main.go
	GOOS=linux GOARCH=amd64 go build -o bin/github-licenses-report-linux-amd64 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/github-licenses-report-windows-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/github-licenses-report-darwin-arm64 main.go