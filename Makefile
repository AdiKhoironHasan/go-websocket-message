#!/bin/bash

run-server:
	@cd server && go run main.go

run-dms:
	@clear
	@cd dms && go run main.go

db-migrate-up:
	@migrate -path db/migration -database "mysql://root@tcp(localhost:3306)/ws_notifications" up

db-migrate-down:
	@migrate -path db/migration -database "mysql://root@tcp(localhost:3306)/ws_notifications" down