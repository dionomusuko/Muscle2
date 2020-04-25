.PHONY: db-set docker-stop docker

db-set:
	docker-compose exec db bash -c 'mysql -u user -p'

docker-stop:
	docker stop dio-mysql

docker:
	docker-compose up -d