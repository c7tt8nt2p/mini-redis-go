test-cov:
	go test -v -coverprofile coverage.out ./... \
	go tool cover -html=coverage.out

vul-check:
	govulncheck ./...