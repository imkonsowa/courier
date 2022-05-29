grpc-gen:
	protoc courierpb/courier.proto --go_out=plugins=grpc:.

up:
	docker-compose build
	docker-compose up -d
down:
	docker-compose down