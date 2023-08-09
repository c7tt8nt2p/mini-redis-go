test-cov:
	go test -v -coverprofile coverage.out -coverpkg=./... ./...
	go tool cover -html=coverage.out

vul-check:
	govulncheck ./...