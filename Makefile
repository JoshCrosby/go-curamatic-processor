# Define variables for docker compose and application
DOCKER_COMPOSE = docker-compose
APP_NAME = go-curamatic-processor

# Default action when running `make`
all: build

# Build the Docker images
build:
	@echo "Building the Docker images..."
	$(DOCKER_COMPOSE) build

# Start the services in the background
up:
	@echo "Starting services..."
	$(DOCKER_COMPOSE) up -d

# Stop the services
down:
	@echo "Stopping services..."
	$(DOCKER_COMPOSE) down

# View logs for the services
logs:
	@echo "Fetching logs..."
	$(DOCKER_COMPOSE) logs -f

# Rebuild the services
rebuild: down build up

# Access the application's Docker container
exec-app:
	@echo "Accessing the application container..."
	$(DOCKER_COMPOSE) exec $(APP_NAME) sh

# Run a bash shell in the app container (alternative to sh for Unix-like systems)
exec-bash:
	@echo "Accessing the application container with Bash..."
	$(DOCKER_COMPOSE) exec $(APP_NAME) bash

# Remove all Docker containers, networks, and volumes
clean:
	@echo "Removing containers, networks, and volumes..."
	$(DOCKER_COMPOSE) down -v

.PHONY: all build up down logs rebuild exec-app exec-bash clean
