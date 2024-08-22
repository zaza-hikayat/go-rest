migrate:
	sh .generate-migration-config.sh && soda migrate up
run:
	go run .