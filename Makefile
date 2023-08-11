test:
	go test -v ./...

test-integration:
	go test -v ./internal/integration_test/...

test-coverage:
	gocov test \
        ./internal/app/... \
        ./internal/config/... \
        ./internal/integration_test/... \
        ./internal/model/... \
        ./internal/service/... \
        ./internal/utils/... \
	| gocov report

test-coverage-html:
	go test -v -coverprofile coverage.out -coverpkg=./internal/... ./...
	go tool cover -html=coverage.out

vul-check:
	govulncheck ./...