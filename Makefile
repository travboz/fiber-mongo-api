.PHONY: up down seed-db clean run build restart list-con

include .env # read from .env file

OUTPUT_BINARY=fiber-mongo
OUTPUT_DIR=./bin
ENTRY_DIR = ./cmd/api

build:
	@echo "Building..."
	@mkdir -p $(OUTPUT_DIR)
	@go build -o $(OUTPUT_DIR)/$(OUTPUT_BINARY) $(ENTRY_DIR)

run: build
	@$(OUTPUT_DIR)/$(OUTPUT_BINARY)

clean:
	@rm -rf $(OUTPUT_DIR)


# docker commands
up:	
	@echo "Starting containers..."
	@docker compose up -d

down:
	@echo "Stopping containers and deleting volumes..."
	@docker compose down -v

# Seeding with users
SCRIPTS_DIR=./

seed-db:
	@echo "Seeding database with users..."
	@bash $(SCRIPTS_DIR)/seed.sh


restart: down up run
	@echo "Restarting..."
