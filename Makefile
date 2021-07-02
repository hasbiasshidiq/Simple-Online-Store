build-appImage:
	docker build -t store-app .

run:
	docker-compose up

stop:
	docker-compose down