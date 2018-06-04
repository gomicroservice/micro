


debug-web:
	docker-compose run --rm -v ${PWD}/frontend/web:/app	--service-ports web npm start
