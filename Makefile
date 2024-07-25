.PHONY: build run clean create

build:
	@echo "Сборка сервера..."
	docker build -t my-golang-shortlink-server .

run:
	docker-compose -f ./docker-compose.yaml up -d 

clean:
	@echo "Очистка..."
	docker stop my-server || true
	docker-compose -f docker-compose.yaml down --remove-orphans
	docker rmi my-golang-shortlink-server || true

createPSQL:
	docker exec -it sl-postgres-1 bash -c "psql -U Mamre32 -c '\c postgres' -c '\i ./mnt/createTable.sql' -c '\i ./mnt/limitations.sql'"

deletePSQL:
	docker exec -it sl-postgres-1 bash -c "psql -U Mamre32 -c '\c postgres' -c '\i ./mnt/dropTable.sql'"
