default: all

IMAGE_NAME := omega-repo
CONTAINER_NAME := omega-repo

all: build


build:
	docker build -t "$(IMAGE_NAME)" .

run:
	docker run -d --net=host --name="$(CONTAINER_NAME)" $(IMAGE_NAME)
