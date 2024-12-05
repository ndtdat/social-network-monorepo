
build-proto:
	rm -rf pkg/api/go && cd api/ && buf mod update && buf format -w && buf build && cd .. && buf generate api
