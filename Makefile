migrate-dev:
	@echo "Running migrate database..."
	go run main.go migrate-up --source postgresql://root:root@localhost:5432/simaku-elearning?sslmode=disable
	@echo "Migrate complete"

srv: 
	@echo "Running HTTP Gateway Server..."
	go run main.go http-gw-srv
.PHONY: migrate
