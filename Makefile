swagger:
	echo " -- Starting swagger generating"
	swag init

migrate:
	echo " -- Migrate database ..."
	go run migrations/setting/migrate.go