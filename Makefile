buf-generate:
	buf generate

run-server:
	cd ./grpc-server && go run main.go

run-echo:
	cd ./echo-server && go run main.go

run-client:
	cd ./grpc-client && go run main.go

sql-generate:
	sqlc generate

echo-swagger-generate:
	swag init -g echo-server/main.go --output echo-server/docs