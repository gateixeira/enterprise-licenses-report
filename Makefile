build:
	go build -o bin/enterprise-licenses-report main.go
compile:
	GOOS=linux GOARCH=arm go build -o bin/enterprise-licenses-report-linux-arm main.go
	GOOS=linux GOARCH=amd64 go build -o bin/enterprise-licenses-report-linux-amd64 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/enterprise-licenses-report-windows-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/enterprise-licenses-report-darwin-arm64 main.go