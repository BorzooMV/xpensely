INIT_DB_SCRIPT_SRC=cmd/scripts/init-db/init-db.go
SEED_EXAMPLE_SCRIPT_SRC=cmd/scripts/seed-db-sample/seed-db-sample.go
CLEAN_DB_SCRIPT_SRC=cmd/scripts/clean-db/clean-db.go

init-db:
	go run ${INIT_DB_SCRIPT_SRC}

clean-db:
	go run ${CLEAN_DB_SCRIPT_SRC}

seed-db-example: clean-db
	go run ${SEED_EXAMPLE_SCRIPT_SRC}
