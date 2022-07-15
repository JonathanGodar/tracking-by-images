build-server:
	go build -o bin/server-main server/*.go

run-migrations:
	sql-migrate up
	sqlboiler psql -c sqlboiler.toml --wipe

generate-code:
	oto -pkg myoto -template ./server/myoto/server.go.plush -out ./server/myoto/server.go ./definition
	oto -pkg api -template ./client/api/oto-client.go.plush -out ./client/api/oto-client.go ./definition
	gofmt -w ./server/myoto/server.go ./server/myoto/server.go
	gofmt -w ./client/api/oto-client.go ./client/api/oto-client.go

run-server:
	bin/server-main

brun-server: build-server run-server
mbrun-server: run-migrations brun-server

brun-client:
	go build -o bin/client-main client/*.go
	bin/client-main
	
