# Generate Client and Server code using proto file
generate-all-pb:
	protoc -I proto \
		proto/piglet-bills/*.proto \
		--go_out=./gen/ \
		--go_opt=paths=source_relative \
		--go-grpc_out=./gen/ \
		--go-grpc_opt=paths=source_relative

	protoc -I proto \
    		proto/piglet-transactions/*.proto \
    		--go_out=./gen/ \
    		--go_opt=paths=source_relative \
    		--go-grpc_out=./gen/ \
    		--go-grpc_opt=paths=source_relative