# Generate Client, Server and gateway code using proto file
generate-all-pb:
	protoc -I ./proto -I ./ \
		proto/piglet-bills/*.proto \
		--go_out ./gen --go_opt paths=source_relative \
		--go-grpc_out ./gen --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./gateway --grpc-gateway_opt paths=source_relative

	protoc -I ./proto -I ./ \
		proto/piglet-transactions/*.proto \
    	--go_out ./gen --go_opt paths=source_relative \
    	--go-grpc_out ./gen --go-grpc_opt paths=source_relative \
    	--grpc-gateway_out ./gateway --grpc-gateway_opt paths=source_relative