build-proto:
	rm -rf pkg/api/go && cd api/ && buf mod update && buf build && cd .. && buf generate api