build:
	docker build -t xsis:v1.0.0 .
run:
	docker-compose up -d
down:
	docker-compose down