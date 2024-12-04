swagger:
	echo " -- Starting swagger generating"
	swag init

MIGRATION_PATH=migrations
DB_USER=root
DB_PASSWORD=root
DB_HOST=localhost
DB_PORT=3306
DB_NAME=golang_service

migrate-up:
	echo "Running migration up..."
	migrate -path $(MIGRATION_PATH) -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" up

migrate-down:
	echo "Running migration down..."
	migrate -path $(MIGRATION_PATH) -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" down
