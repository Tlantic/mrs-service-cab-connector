.PHONY: proto bundle install

install:
	go mod vendor -v

bundle: install