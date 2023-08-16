test:
	go test -v ./...

test-integration:
	go test -v ./e2e/...

test-coverage:
	gocov test \
	  	./e2e/... \
        ./internal/app/... \
        ./internal/config/... \
        ./internal/model/... \
        ./internal/service/... \
        ./internal/utils/... \
	| gocov report

test-coverage-html:
	go test -v -coverprofile coverage.out -coverpkg=./internal/... ./...
	go tool cover -html=coverage.out

vul-check:
	govulncheck ./...