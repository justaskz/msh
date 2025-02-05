test:
	@ go test ./...

build_app:
	@ go build -o bin/msh cmd/msh.go
	@ chmod +x bin/msh

# install:
# 	@ MASHINA_DEV=true bash install.sh

auto_test:
	@ fswatch -or . | xargs -n1 -I{} go test ./...

build:
	@ docker compose build

up:
	@ docker compose up -d

down:
	@ docker compose down

connect:
	@ docker compose exec app bash
