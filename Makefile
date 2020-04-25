.PHONY: db-set docker-stop docker

db-set:
	docker-compose exec db bash -c 'mysql -u user -p'

stop:
	docker stop muslce-db

docker:
	docker-compose up -d