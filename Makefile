postgres: 
	docker run --name postgres-db -e POSTGRES_PASSWORD=docker -p 5432:5432 -d postgres
createdb:
	docker exec -it postgres-db psql -U postgres -c "CREATE DATABASE golang;"
dropdb:
	docker exec -it postgres-db psql -U postgres -c "DROP TABLE golang;"
.PHONY: postgres createdb dropdb

