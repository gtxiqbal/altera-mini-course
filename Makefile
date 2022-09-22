IMAGE_TAG_NAME=gtxiqbal4559/altera-mini-course:latest
build:
	go build -o ./Day-6

down:
	docker compose down --rmi local --remove-orphans -v

run:
	clear && go run main.go

show-test-result:
	go tool cover -html=cover.ou

stop:
	docker compose stop

test-all:
	clear && go test ./... -coverprofile=cover.out -p 1 && go tool cover -func=cover.out

up:
	docker compose --project-name gtxiqbal4559/altera-mini-course:latest up -d && docker compose start

build-image:
	docker build . -t ${IMAGE_TAG_NAME} -f Dockerfile

push-image:
	docker push ${IMAGE_TAG_NAME}