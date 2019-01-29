# -*- mode: Makefile-gmake -*-

SHELL := bash

TOP_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))

BUILD_DIR := build

DIAGRAMS := doc/diagram/examples.png

doc/diagram/%.png: doc/%.dot
	@mkdir -p $(DIAGRAM_DIR)
	dot -Tpng $< > $@

$(BUILD_DIR):
	@mkdir -p $@

clean:
	rm -rf $(BUILD_DIR)

fmt:
	go fmt ./...

vet:
	go vet ./...

doc: $(DIAGRAMS)

examples: $(BUILD_DIR)
	go build -o $(TOP_DIR)/$(BUILD_DIR)/rest github.com/shisa-platform/examples/rest
	go build -o $(TOP_DIR)/$(BUILD_DIR)/rpc github.com/shisa-platform/examples/rpc
	go build -o $(TOP_DIR)/$(BUILD_DIR)/idp github.com/shisa-platform/examples/idp
	go build -o $(TOP_DIR)/$(BUILD_DIR)/gw github.com/shisa-platform/examples/gw

docker:
	docker build --tag shisa/examples/gw -f gw/Dockerfile .
	docker build --tag shisa/examples/idp -f idp/Dockerfile .
	docker build --tag shisa/examples/rest -f rest/Dockerfile .
	docker build --tag shisa/examples/rpc -f rpc/Dockerfile .

.PHONY: examples docker
