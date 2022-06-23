test:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out
dev:
	go install github.com/cosmtrek/air@latest
	air
mock:
	go install github.com/vektra/mockery/v2@latest
	mockery --all --inpackage --case snake