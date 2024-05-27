proto-user:
	@protoc --go_out=. --go_opt=paths=source_relative \
       		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
       		internal/proto/user/user.proto

proto-credit-card:
	@protoc --go_out=. --go_opt=paths=source_relative \
       		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
       		internal/proto/credit_card/credit_card.proto

server-keys:
	cd internal/tlsconfig/cert/server/; sh gen.sh;

client-keys:
	cd internal/tlsconfig/cert/client/; sh gen.sh;