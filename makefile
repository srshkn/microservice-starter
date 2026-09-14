# ./

ifneq (,$(wildcard .env))
	include .env
	export
endif

.DEFUALT_GOAL := help

# -------------------------------------------------------------------------
# PHONY

# Base
.PHONY: help

# -------------------------------------------------------------------------
# BASE

help:
	@echo "microservice-starter"
	@echo ""
	@echo "Usage:"
	@echo "  make <target>"
	@echo ""
