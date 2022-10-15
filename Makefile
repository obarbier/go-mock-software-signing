# tool macros
GOPATH ?= /home/obarbier/sdk/go1.19.1/
GO ?= $(GOPATH)/bin/go
SWAGGER := $(GOPATH)/bin/swagger
# default rule
default: tools-install all

# phony rules
.PHONY: tools-install
tools-install:
	cd tools;\
    $(GO) install github.com/go-swagger/go-swagger/cmd/swagger;

.PHONY: swagger-validate
swagger-validate: core-swagger-validate

.PHONY: core-swagger-validate
core-swagger-validate:
	cd core;\
	$(SWAGGER) validate ./swagger.yml

#.PHONY: core-go-generate
#core-go-generate:
#	- export PATH=$(PATH):$(GOPATH);\
#		cd core;\
#		$(GO) generate ./...

.PHONY: generate-ssl-certificate
generate-ssl-certificate:
	- echo "generate ssl certs"

.PHONY: core-start-server
core-start-server:
	- cd core;\
	  export MYAPP_LOG_LEVEL=TRACE; \
	  #export MYAPP_LOG_FILE_OUTPUT=true; \
	  $(GO)  run ./pkg/cmd/core-server/main.go \
	  --port 8080

.PHONY: core-stress-test
core-stress-test:
	- ab  -n 2 -c 2 -H 'Authorization: Basic b2JhcmJpZXI6Y2hhbmdlaXQ=' http://localhost:8080/api/v1/user/3 >> "./docs/benchmark_get_user_$(shell date +'%Y_%m_%d_%H-%M-%S').txt"



.PHONY: core-format-project
core-format-project:
	- cd core;\
	  $(GO)  fmt ./



.PHONY: all
all: $(TARGET)
