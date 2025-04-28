docker-up:
	docker compose up -d --build

app-start:
	docker compose exec app sh

docker-down:
	docker compose down -v

docker-stop:
	docker compose stop

docker-start:
	docker compose start