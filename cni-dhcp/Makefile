REGISTRY_NAME = zdnscloud
IMAGE_Name = cni-dhcp-init
IMAGE_VERSION = v0.1

.PHONY: all container

all: container

container: 
	docker build -t $(REGISTRY_NAME)/$(IMAGE_Name):$(IMAGE_VERSION) ./ --no-cache
