default: all

IMAGE_NAME := datamancloud-verfiy
CONTAINER_NAME := datamancloud-verfiy
VERFIY_FILE := "./verifycode.go"

all: run


run:
	docker build -t "$(IMAGE_NAME)" .
	docker run -d --name="$(CONTAINER_NAME)" $(IMAGE_NAME)
	docker cp $(CONTAINER_NAME):/verify ./verify
	./verify $(VERFIY_FILE)
	docker rm -f $(CONTAINER_NAME)
