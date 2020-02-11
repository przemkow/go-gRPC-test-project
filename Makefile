build_client: 
		go build -o ./dist/ ./grpc-client

build_server: 
		go build -o ./dist/ ./grpc-server

build: build_client build_server
	
run_client:
	./dist/grpc-client
		
run_server:
	./dist/grpc-server