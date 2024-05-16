all:
	@echo "Help:"
	@echo "make <target>"
	@echo "\nWhere <target> is one of:"
	@echo "\tbuild_frotned - Installs packages and builds for development environment" 
	@echo "\trun_frontend_dev - Runs frontend in development environment"
	@echo "\trun_frotned_preview - Runs frontend in a production like environment"
	@echo "\tbuild_backend - Builds the server"
	@echo "\trun_backend - Runs the server"
build_frontend:
	
	@npm install --workspace quiz_frontend/
	@npm run build --workspace quiz_frontend/

run_frontend_dev:
	@npm run dev --workspace quiz_frontend/

run_frontend_preview:
	@npm run preview --workspace quiz_frontend/

build_backend:
	@go build -C ./quiz_server/ .

run_backend:
	@go run -C ./quiz_server/ .

run_dev: run_frontend_dev run_backend
