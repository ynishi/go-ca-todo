#!/bin/sh

set -eux

# build cmd excutable
CLI_DIR=./cmd/todo
CLI_BIN=$CLI_DIR/todo

go build -o $CLI_BIN $CLI_DIR
chmod +x $CLI_BIN

# buid web server
WEB_DIR=./web/todo
WEB_BIN=$WEB_DIR/todo

go build -o $WEB_BIN $WEB_DIR
chmod +x $WEB_BIN
