# Server
server:
	@echo "Starting server..."
	@go run ./main.go
	
# Create a new postgres migration. make create-migration-postgres name="migration name"
create-migration-postgres:
        @echo "Creating migration"
        @goose -dir src/infra/postgres/migrations create $$name go

# Up the migrations
migration-up-postgres:
        @echo "Up the migrations"
        @go run src/infra/goose/migrations.go up