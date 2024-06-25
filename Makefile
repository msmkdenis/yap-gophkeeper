proto-user:
	@protoc --go_out=. --go_opt=paths=source_relative \
       		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
       		internal/proto/user/user.proto

proto-credit-card:
	@protoc --go_out=. --go_opt=paths=source_relative \
       		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
       		internal/proto/credit_card/credit_card.proto

proto-text-data:
	@protoc --go_out=. --go_opt=paths=source_relative \
       		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
       		internal/proto/text_data/text_data.proto

proto-credentials:
	@protoc --go_out=. --go_opt=paths=source_relative \
       		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
       		internal/proto/credentials/credentials.proto

server-keys:
	cd internal/tlsconfig/cert/server/; sh gen.sh;

client-keys:
	cd internal/tlsconfig/cert/client/; sh gen.sh;

mock-credit-card-service:
	@mockgen --build_flags=--mod=mod \
			 -destination=internal/mocks/credit_card/mock_credit_card_service.go \
			 -package=mocks github.com/msmkdenis/yap-gophkeeper/internal/credit_card/api/v1/grpchandlers CreditCardService

mock-text-service:
	@mockgen --build_flags=--mod=mod \
			 -destination=internal/mocks/text_data/mock_text_data_service.go \
			 -package=mocks github.com/msmkdenis/yap-gophkeeper/internal/text_data/api/v1/grpchandlers TextDataService