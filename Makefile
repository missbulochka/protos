# Generate Client, Server and gateway code using proto file
# flag --openapiv2_out generates documentation
generate-all-pb:
	protoc -I ./proto -I ./ \
		proto/piglet-bills/*.proto \
		--go_out ./gen --go_opt paths=source_relative \
		--go-grpc_out ./gen --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./gen --grpc-gateway_opt paths=source_relative \
		--openapiv2_out ./openapiv2

	protoc -I ./proto -I ./ \
		proto/piglet-transactions/*.proto \
    	--go_out ./gen --go_opt paths=source_relative \
    	--go-grpc_out ./gen --go-grpc_opt paths=source_relative \
    	--grpc-gateway_out ./gen --grpc-gateway_opt paths=source_relative \
    	--openapiv2_out ./openapiv2