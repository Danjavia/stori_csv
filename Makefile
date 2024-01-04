build:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/main cmd/main.go

deploy: build
	sls deploy --verbose --stage prod --aws-profile personal

run_sls_local: @serverless dev --verbose --stage prod --aws-profile personal
	
run:
	@go run cmd/main.go

test:
	@go test -v ./..
