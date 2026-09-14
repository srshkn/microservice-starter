# ./services/gateway

ifneq (,$(wildcard .env))
	include .env
	export
endif

.DEFUALT_GOAL = help

# -------------------------------------------------------------------------
# CONSTANTS

# Environment
ENV_FILE=.env
ENV_EXAMPLE=.env.example

# API
V1API=api/v1/api.yml

# -------------------------------------------------------------------------
# PHONY

# Base
.PHONY: help

# Generation
.PHONY: env-gen v1api-gen

# -------------------------------------------------------------------------
# BASE

help:
	@echo "microservice-starter/services/gateway"
	@echo ""
	@echo "Usage:"
	@echo "  make <target>"
	@echo ""
	@echo "Generation:"
	@echo "  env-gen              Create .env from .env.example"
	@echo "  v1api-gen            Generate Go code from OpenAPI specification"
	@echo ""

# -------------------------------------------------------------------------
# GENERATION

env-gen:
	@if [ ! -f $(ENV_FILE) ]; then \
		echo "Creating .env from .env.example"; \
		cp $(ENV_EXAMPLE) $(ENV_FILE); \
	fi

v1api-gen:
	go tool oapi-codegen -config ./api/v1/config/server.yml ./$(V1API)
	go tool oapi-codegen -config ./api/v1/config/models.yml ./$(V1API)
	go tool oapi-codegen -config ./api/v1/config/spec.yml ./$(V1API)
