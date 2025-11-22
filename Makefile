HOSTNAME=registry.terraform.io
NAMESPACE=rrhawk
NAME=vertexairag
BINARY=terraform-provider-${NAME}
VERSION=0.1.1
OS_ARCH=linux_amd64

default: install

build:
	cd terraform-provider-vertexairag && go build -o ${BINARY}

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv terraform-provider-vertexairag/${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

clean:
	rm -f terraform-provider-vertexairag/${BINARY}
