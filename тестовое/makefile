// ТУТ СВОИ ДАННЫЕ
migration_up: migrate -path internal/database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose up

migration_down: migrate -path internal/database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose down

migration_fix: migrate -path internal/database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" force VERSION