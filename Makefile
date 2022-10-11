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
	

.PHONY: all
all: $(TARGET)
