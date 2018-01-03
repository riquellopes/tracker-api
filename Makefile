.SILENT:

front-end-test:
	docker-compose stop
	docker-compose run --rm frontend npm test
