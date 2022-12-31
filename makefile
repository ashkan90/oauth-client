generate-mocks:
	mockgen -source=./pkg/oauth/processor.go -destination=./pkg/oauth/mocks/processor_mock.go -package=mocks
