start:
	docker-compose -f docker/docker-compose.yml up
stop:
	docker-compose -f docker/docker-compose.yml down
restart:
	docker-compose -f docker/docker-compose.yml down --remove-orphans && \
	docker-compose -f docker/docker-compose.yml up --build
remove:
	docker-compose -f docker/docker-compose.yml down --rmi all
proto:
	docker run -it --rm -v $(PWD)/app/protos:/protos docker_scraper_service sh \
	-c "cd /protos && protoc /protos/test.proto --go_out=plugins=grpc:. --proto_path=/protos"
