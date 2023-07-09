APP_VERSION=v$(shell cat version)
APP_CONFIG_PATH=$(shell pwd)/config.ini

ldflags = -X 'github.com/Liang-jonas/jnote/Cmd.version=${APP_VERSION}'
ldflags += -X 'github.com/Liang-jonas/jnote/Cmd.defaultCfgPath=${APP_CONFIG_PATH}'

.PHONY: build

build:
	go build -ldflags "${ldflags}" -o jnote-server
