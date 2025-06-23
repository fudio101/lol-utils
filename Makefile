# Makefile for Wails project

.PHONY: help dev build

help:
	@echo "Available targets:"
	@echo "  dev    - Run Wails in development mode (hot reload)"
	@echo "  build  - Build the Wails app for production"

# Run Wails in development mode
dev:
	cd cmd/lol-utils && wails dev

# Build the Wails app for production
build:
	cd cmd/lol-utils && wails build 