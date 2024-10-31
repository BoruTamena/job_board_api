sqlc:
	- sqlc generate -f ./config/sqlc.yaml

swagger:
	- swag init -g initiator/initiator.go

migrate-create:
	- migrate create -ext sql -dir internal/domain/constant/schemas -tz "UTC" $(args)