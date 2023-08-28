MIGRATE_CMD = migrate -database mysql://root:@/go-cardio-hexa-go -path internal/database/migrations

migrate-up:
	${MIGRATE_CMD} up
migrate-down:
	${MIGRATE_CMD} down