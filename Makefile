generate:
	openapi-generator generate -g go -i specs/mgmt.yaml -o client/mgmt/go
	openapi-generator generate -g go-server -i specs/mgmt.yaml -o svc/mgmt --additional-properties hideGenerationTimestamp=true
	openapi-generator generate -g typescript-angular -i specs/mgmt.yaml -o ui/src/app/api/mgmt
