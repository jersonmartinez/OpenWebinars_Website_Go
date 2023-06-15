#!/bin/bash

mkdir -p \
    cmd/main \
    internal/config \
    internal/handlers \
    internal/models \
    internal/routes \
    web/templates \
    web/static/css \
    web/static/js \

touch \
    cmd/main/main.go \
    internal/config/config.go \
    internal/handlers/handlers.go \
    internal/models/models.go \
    internal/routes/routes.go \
    web/main.go \
    web/templates/home.html \
    web/templates/error.html \
    web/static/css/custom.css \
    web/static/js/custom.js \
