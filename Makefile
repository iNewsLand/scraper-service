start:
	docker-compose -f docker/docker-compose.yml up -d
stop:
	docker-compose -f docker/docker-compose.yml down
restart:
	docker-compose -f docker/docker-compose.yml down --remove-orphans && \
	docker-compose -f docker/docker-compose.yml up --build
remove:
	docker-compose -f docker/docker-compose.yml down --rmi all
