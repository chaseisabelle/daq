PROTO_DIR := ${PWD}/api
PB_OUT := ${PWD}/pkg

genpb:
	rm -Rf ${PB_OUT} && \
	mkdir ${PB_OUT} && \
	docker pull jaegertracing/protobuf && \
	docker run --rm  \
		-v${PROTO_DIR}:${PROTO_DIR} \
		-v${PB_OUT}:${PB_OUT} \
		-w${PB_OUT} \
		jaegertracing/protobuf:latest \
			--proto_path=${PROTO_DIR} \
			--go_out=plugins=grpc:${PB_OUT} \
			-I/usr/include/github.com/gogo/protobuf \
			${PROTO_DIR}/daq.proto

test:
	go test ./...

vet:
	go vet ./...